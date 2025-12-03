package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// CRDTOperation representa la operación que vamos a replicar
type CRDTOperation struct {
	Op     string `json:"op"`
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
	Value  string `json:"value,omitempty"`
	Ts     int64  `json:"ts"`
	NodeID string `json:"node_id"`
}

// WatchKV inicia un watch sobre un bucket KV y publica cambios en rep.kv.ops
func WatchKV(sourceKV, repKV nats.KeyValue, nodeID string) {
	watch, err := sourceKV.WatchAll()
	if err != nil {
		log.Fatalf("Error iniciando watch: %v", err)
	}

	for update := range watch.Updates() {
		var opType string
		switch update.Operation() {
		case nats.KeyValuePut:
			opType = "put"
		case nats.KeyValueDelete:
			opType = "delete"
		default:
			continue
		}

		op := CRDTOperation{
			Op:     opType,
			Bucket: sourceKV.Bucket(),
			Key:    update.Key(),
			Value:  string(update.Value()),
			Ts:     time.Now().Unix(),
			NodeID: nodeID,
		}

		data, _ := json.Marshal(op)
		fmt.Printf("CRDT generado: %s\n", data)

		// Publicar en rep.kv.ops con un ID único
		repKV.Put(fmt.Sprintf("%d", time.Now().UnixNano()), data)
	}
}
