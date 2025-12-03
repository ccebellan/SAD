package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	natsURL := flag.String("nats-url", "nats://localhost:4222", "NATS server URL")
	bucket := flag.String("bucket", "", "KV bucket")
	nodeID := flag.String("node-id", "", "Node ID")

	flag.Parse()

	fmt.Println("URL:", *natsURL)
	fmt.Println("Bucket:", *bucket)
	fmt.Println("Node ID:", *nodeID)

	// Conexión a NATS
	nc, err := nats.Connect(*natsURL)
	if err != nil {
		log.Fatalf("Error conectando a NATS: %v", err)
	}
	defer nc.Drain()

	// Abrir bucket KV
	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("Error obteniendo JetStream: %v", err)
	}

	kv, err := js.KeyValue(*bucket)
	if err != nil {
		log.Fatalf("Error abriendo bucket KV %s: %v", *bucket, err)
	}

	repKV, err := js.KeyValue("rep.kv.ops")
	if err != nil {
		log.Fatalf("Error abriendo bucket KV %s: %v", "rep.kv.ops", err)
	}

	// Iniciar watch en otro goroutine
	go WatchKV(kv, repKV, *nodeID)

	// Mantener la aplicación corriendo
	select {}
}
