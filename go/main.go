package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("go working!")

	// loop log
	go func() {
		n, tick := 0, 0
		for {
			n++
			time.Sleep(5 * time.Second)
			if (n < 100) || (n%12 == 0) {
				tick++
				log.Printf("tick=%d n=%d", tick, n)
			}
		}
	}()

	// web server
	addr, dir := ":8080", "/www"
	// get config from env
	if v, ok := os.LookupEnv("WEB_HTTP_ADDR"); ok {
		addr = v
	}
	if v, ok := os.LookupEnv("WEB_HTTP_DIR"); ok {
		dir = v
	}
	// from cli
	if len(os.Args) > 1 {
		addr = os.Args[1]
		if len(os.Args) > 2 {
			dir = os.Args[2]
		}
	}

	log.Printf("web server started in dir=%s at addr=%s", dir, addr)
	log.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir(dir))))
}
