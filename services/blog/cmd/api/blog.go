package main

import (
	"log"
	"net/http"
	"wang/common/logger"
	"wang/services/blog/cmd/api/handler"
)

func main() {
	logger.Info("aaa")
	logger.Infof("bb: %s", "aaa")
	logger.Errorf("ccc: %s", "aaa")
	logger.Error("ddd")
	mux := http.NewServeMux()
	mux.HandleFunc("/demo", handler.NewDemoHandler().Demo)
	var serve = http.Server{
		Addr: ":8001",
		Handler: mux,
	}
	err := serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

