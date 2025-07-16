package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	upd "github.com/Christian1984/go-update-checker"
	"github.com/Songmu/prompter"
)

//go:embed content
var Book embed.FS

var VERSION = "dev"

func main() {
	port := flag.Int("port", 8080, "port for listen and serve example 8080")
	version := flag.Bool("version", false, "show version and exit")
	flag.Parse()

	if *version {
		fmt.Printf("GoFarsi Book Version: %s\n", VERSION)
		os.Exit(0)
	}

	if *port <= 0 || *port >= 65535 {
		*port = 8080
	}

	address := fmt.Sprintf(":%d", *port)

	subFS, err := fs.Sub(Book, "content")
	if err != nil {
		log.Fatal(err)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	fileFS := http.FileServer(http.FS(subFS))
	http.Handle("/", fileFS)

	if updateAvailable, downloadLink := isAvailableNewVersion(); updateAvailable {
		if prompter.YN("Are you want update book to new version?", false) {
			openBrowser(downloadLink)
			os.Exit(0)
		}
	}

	ready := make(chan struct{})
	go func() {
		<-ready
		time.Sleep(100 * time.Millisecond)
		if err := openBrowser(fmt.Sprintf("http://localhost:%d", *port)); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("start serve on address %s", address)
	close(ready)
	log.Fatal(http.ListenAndServe(address, logRequest(http.DefaultServeMux)))
}

func isAvailableNewVersion() (bool, string) {
	uc := upd.New("GoFarsi", "book", "Go Programming Language Persian",
		"https://github.com/GoFarsi/book/releases", 0, false)
	uc.CheckForUpdate(VERSION)
	if uc.UpdateAvailable {
		uc.PrintMessage()
	}
	return uc.UpdateAvailable, uc.DownloadLink
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
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
