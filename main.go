package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/dmitrymomot/microservice-template/proto/microservice_template"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

var (
	wg sync.WaitGroup

	grpc  = flag.Bool("grpc", false, "run with grpc handler")
	twirp = flag.Bool("twirp", false, "run with twirp handler")
)

func init() {
	flag.Parse()
}

func main() {
	srv := NewService()

	if *grpc {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println("start grpc server")
			lis, err := net.Listen("tcp", grpcPort)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			s := grpc.NewServer()
			microservice_template.RegisterServiceServer(s, &server{id: uuid.New()})
			if err := s.Serve(lis); err != nil {
				log.Fatalf("grpc server: %+v", err)
			}
		}()
	}

	if *twirp {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println("start twirp server")
			handler := microservice_template.NewServiceServer(&server{}, nil)

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
