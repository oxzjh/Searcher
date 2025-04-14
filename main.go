package main

import (
	"api/searcher"
	"flag"
	"fmt"
	"golib/server"
	"golib/server/http"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var addr string
	flag.StringVar(&addr, "a", "0.0.0.0:3000", "Addr")
	flag.Parse()

	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	searcher.Initialize(filepath.Join(configDir, "Searcher"))

	fmt.Println("Serve HTTP on", addr)
	log.Fatal(server.ServeHTTP(addr, http.NewServer(), 5*time.Second))
}
