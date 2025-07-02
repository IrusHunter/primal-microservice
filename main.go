package main

import (
	"flag"
)

func main() {
	// client := client.New("http://localhost:3000")

	// result, err := client.Calculate(context.Background(), 1000, 0, "/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%f%s%f=%f", result.A, result.Operation, result.B, result.Result)
	// return
	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running")
	flag.Parse()

	svc := NewLoggingService(DummyCalculator{})

	// res, err := svc.Calculate(context.Background(), 10, 5, "+")
	// if err != nil {
	// 	log.Fatalf("Error: %v", err)
	// }
	//
	// fmt.Println("Result:", res)

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
