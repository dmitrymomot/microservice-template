package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dmitrymomot/microservice-template/pb/examplesrv"
	"github.com/dmitrymomot/microservice-template/service"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

const (
	envLocal = "local"
	envProd  = "production"
)

var (
	grpcSrv  = flag.Bool("grpc", false, "run with grpc handler")
	twirpSrv = flag.Bool("twirp", false, "run with twirp handler")

	env = flag.String("env", envLocal, "set environment")

	grpcPort = ":8888"
	httpPort = ":8080"
)

func init() {
	flag.Parse()
}

func main() {
	var err error

	eg, _ := errgroup.WithContext(newCtx())

	// TLS config for do services
	var tlsConfig *tls.Config
	if *env == envProd {
		tlsConfig, err = loadTLSRootCert(os.Getenv("DO_CERT_PATH"))
		if err != nil {
			panic(err)
		}
	}

	// Set up datapase connection
	db := pg.Connect(&pg.Options{
		Addr:            fmt.Sprintf("%s:%s", os.Getenv("PG_HOST"), os.Getenv("PG_PORT")),
		User:            os.Getenv("PG_USER"),
		Password:        os.Getenv("PG_PASSWORD"),
		Database:        os.Getenv("PG_DB"),
		ApplicationName: os.Getenv("APP_NAME"),
		PoolSize:        22,
		TLSConfig:       tlsConfig,
	})
	defer db.Close()

	if *env == envLocal {
		// Log each database query
		db.AddQueryHook(newDBLogger())
	}

	// Init service
	srv := service.NewService()

	if *grpcSrv {
		eg.Go(runGRPCServer(srv))
	}

	if *twirpSrv {
		eg.Go(runHTTPServer(srv))
	}

	if err := eg.Wait(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to wait goroutine group: %v.", err.Error())
	}

	log.Println("service has been stopped")
}

func runGRPCServer(srv examplesrv.Service) func() error {
	return func() error {
		log.Println("start grpc server")

		lis, err := net.Listen("tcp", grpcPort)
		if err != nil {
			return errors.Wrap(err, "failed to listen")
		}

		s := grpc.NewServer()

		// registering of grpc services
		examplesrv.RegisterServiceServer(s, srv)

		if err := s.Serve(lis); err != nil {
			return errors.Wrap(err, "grpc server")
		}

		return nil
	}
}

func runHTTPServer(srv examplesrv.Service) func() error {
	return func() error {
		log.Println("start twirp server")
		handler := examplesrv.NewServiceServer(srv, nil)

		r := chi.NewRouter()
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("server is working " + handler.ProtocGenTwirpVersion() + " - " + handler.PathPrefix()))
		})

		// Registering of http services
		r.Mount(handler.PathPrefix(), handler)

		if err := http.ListenAndServe(httpPort, r); err != nil {
			return errors.Wrap(err, "http server")
		}
		return nil
	}
}

func loadTLSRootCert(path string) (*tls.Config, error) {
	caPool := x509.NewCertPool()
	severCert, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	caPool.AppendCertsFromPEM(severCert)
	return &tls.Config{
		RootCAs:            caPool,
		InsecureSkipVerify: true,
	}, nil
}

func newCtx() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sCh := make(chan os.Signal, 1)
		signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM)
		<-sCh
		cancel()
	}()
	return ctx
}
