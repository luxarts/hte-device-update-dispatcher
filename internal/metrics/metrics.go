package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// StartServer creates a http server on default port 8081 with the endpoint /metrics which calls the prometheus handler
func StartServer(port ...string) {
	if len(port) == 0 {
		port = []string{"8081"}
	}
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Serving metrics...")
		if err := http.ListenAndServe(":"+port[0], nil); err != nil {
			log.Fatalf("err creating metrics server: %+v\n", err)
		}
	}()
}
