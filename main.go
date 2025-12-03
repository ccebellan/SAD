package main

import (
    "flag"
    "fmt"
)

func main() {
    natsURL := flag.String("nats-url", "nats://localhost:4222", "NATS server URL")
    bucket := flag.String("bucket", "", "KV bucket")
    nodeID := flag.String("node-id", "", "Node ID")
    repSubj := flag.String("rep-subj", "", "Replication subject")

    flag.Parse()

    fmt.Println("URL:", *natsURL)
    fmt.Println("Bucket:", *bucket)
    fmt.Println("Node ID:", *nodeID)
    fmt.Println("Rep Subj:", *repSubj)
}
