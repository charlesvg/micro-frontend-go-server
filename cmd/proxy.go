package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"micro-frontend-go-server/internal"
	"net/http"
	"net/http/httputil"
	"os"
)

func initlog(logLevel string) {
	parsedLevel, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Panicln(err)
	}
	log.SetOutput(os.Stdout)
	log.SetLevel(parsedLevel)
	log.SetFormatter(&log.TextFormatter{ ForceColors: true })
}

func main() {

	var config = internal.ReadProxyConfig()

	initlog(config.Log.Level)

	//port := flag.Int("port", 8080, "port to listen on")
	//targetURL := flag.String("target-url", "", "downstream service url to proxy to")
	//flag.Parse()
	//
	//u, err := url.Parse(*targetURL)
	//if err != nil {
	//	log.Fatalf("Could not parse downstream url: %s", *targetURL)
	//}

	proxy := httputil.NewSingleHostReverseProxy(config.Server.DownStreamURL.URL)

	director := proxy.Director
	proxy.Director = func(req *http.Request) {
		director(req)
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = req.URL.Host
	}

	http.HandleFunc("/", proxy.ServeHTTP)
	log.Println("Listening on port ", config.Server.Port )
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port ), nil))
}
