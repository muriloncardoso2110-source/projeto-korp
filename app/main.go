package main

import (
"encoding/json"
"fmt"
"log"
"net/http"
"runtime"
"time"
)

type Response struct {
Nome    string `json:"nome"`
Horario string `json:"horario"`
}

// Handler manual para simular o formato de texto plano que o Prometheus ama ler
func metricsHandler(w http.ResponseWriter, r *http.Request) {
var m runtime.MemStats
runtime.ReadMemStats(&m)

w.Header().Set("Content-Type", "text/plain; version=0.0.4")
w.WriteHeader(http.StatusOK)

// Cospe a métrica exata de memória alocada que o Grafana vai ler
fmt.Fprintf(w, "# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.\n")
fmt.Fprintf(w, "# TYPE go_memstats_heap_alloc_bytes gauge\n")
fmt.Fprintf(w, "go_memstats_heap_alloc_bytes %d\n", m.HeapAlloc)
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodGet {
w.WriteHeader(http.StatusMethodNotAllowed)
return
}
res := Response{
Nome:    "Projeto Korp",
Horario: time.Now().UTC().Format(time.RFC3339),
}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(res)
}

func main() {
http.HandleFunc("/metrics", metricsHandler)
http.HandleFunc("/projeto-korp", projetoKorpHandler)

log.Println("Servidor iniciado na porta 8080...")
if err := http.ListenAndServe(":8080", nil); err != nil {
log.Fatalf("Erro: %s", err)
}
}
