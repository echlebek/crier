// Utter cries, and listen for harkenings.
package main

import (
	"log"
	"net"
	"sync"
	"time"

	"github.com/echlebek/crier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	lastHark string
	mu       sync.Mutex
}

func (s *server) ListenForCries(srv crier.Crier_ListenForCriesServer) error {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			status := "all's well!"
			s.mu.Lock()
			if s.lastHark != "" {
				status = s.lastHark
			}
			s.mu.Unlock()
			cry := &crier.Cry{
				Time:   time.Now().Unix(),
				Status: status,
			}
			if err := srv.Send(cry); err != nil {
				log.Println(err)
				return
			}
		}
	}()
	for {
		hark, err := srv.Recv()
		if err != nil {
			return err
		}
		s.mu.Lock()
		log.Println("Hark!", hark.Harkening)
		s.lastHark = hark.Harkening
		s.mu.Unlock()
	}
	return nil
}

func main() {
	conn, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	crier.RegisterCrierServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
