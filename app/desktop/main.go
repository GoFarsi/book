//go:build aix || darwin || dragonfly || freebsd || linux || nacl || netbsd || openbsd || solaris || windows

package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

//go:embed content
var Book embed.FS

func main() {
	port := flag.Int64("port", 8080, "port for listen and serve example 8080")
	flag.Parse()

	if port == nil || *port <= 0 || *port >= 65535 {
		*port = 8080
	}

	address := fmt.Sprintf(":%s", strconv.Itoa(int(*port)))

	subFS, err := fs.Sub(Book, "content")
	if err != nil {
		log.Fatalln(err)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	fileFS := http.FileServer(http.FS(subFS))

	http.Handle("/", fileFS)

	log.Printf("start serve on address %s", address)

	go func() {
		<-time.After(100 * time.Millisecond)
		if err := openBrowser(fmt.Sprintf("http://localhost:%s", strconv.Itoa(int(*port)))); err != nil {
			log.Fatalln(err)
		}
	}()

	log.Fatalln(http.ListenAndServe(address, logRequest(http.DefaultServeMux)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	log.Printf("start address %s in browser", url)
	return exec.Command(cmd, args...).Start()
}
