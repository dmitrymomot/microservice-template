package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/dmitrymomot/microservice-template/pb/examplesrv"
	"github.com/dmitrymomot/microservice-template/service"
	"github.com/go-chi/chi"
	"google.golang.org/grpc"
)

var (
	wg sync.WaitGroup

	grpcSrv  = flag.Bool("grpc", false, "run with grpc handler")
	twirpSrv = flag.Bool("twirp", false, "run with twirp handler")

	grpcPort = ":8888"
	httpPort = ":8080"
)

func init() {
	flag.Parse()
}

func main() {
	srv := service.NewService()

	if *grpcSrv {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println("start grpc server")
			lis, err := net.Listen("tcp", grpcPort)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			s := grpc.NewServer()
			examplesrv.RegisterServiceServer(s, srv)
			if err := s.Serve(lis); err != nil {
				log.Fatalf("grpc server: %+v", err)
			}
		}()
	}

	if *twirpSrv {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println("start twirp server")
			handler := examplesrv.NewServiceServer(srv, nil)

			r := chi.NewRouter()
			r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("server is working " + handler.ProtocGenTwirpVersion() + " - " + handler.PathPrefix()))
			})
			r.Mount(handler.PathPrefix(), handler)

			log.Fatalf("http server: %+v", http.ListenAndServe(httpPort, r))
		}()
	}

	wg.Wait()

	log.Println("service has been stopped")
}
