package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <path_to_JSON_config>\n", os.Args[0])
		return
	}
	config, err := ReadConfigFile(os.Args[1])
	if err != nil {
		str := fmt.Sprintf("Reading config error: %s\n", err.Error())
		panic(str)
	}

	server := NewServer(config)
	mux := http.NewServeMux()
	mux.HandleFunc("/", server.getRootHandler())
	mux.HandleFunc("/bc", server.getBCHandler())

	port := fmt.Sprintf(":%d", config.PORT)
	log.Println("Server is on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, mux))
}
