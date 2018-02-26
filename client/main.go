// Listen for cries, and utter harkenings.
package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/echlebek/crier"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:6666", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client, err := crier.NewCrierClient(conn).ListenForCries(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		r := bufio.NewReader(os.Stdin)
		for {
			msg, err := r.ReadString('\n')
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Println(err)
				continue
			}
			hark := &crier.Hark{
				// truncate newline
				Harkening: msg[:len(msg)-1],
			}
			if err := client.Send(hark); err != nil {
				log.Println(err)
			}
		}
	}()
	for {
		cry, err := client.Recv()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("It's", time.Unix(cry.Time, 0), "and", cry.Status)
	}
}
