//go:build aix || darwin || dragonfly || freebsd || linux || nacl || netbsd || openbsd || solaris || windows

package main

import (
	"embed"
	"flag"
	"fmt"
	upd "github.com/Christian1984/go-update-checker"
	"github.com/Songmu/prompter"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

//go:embed content
var Book embed.FS

const VERSION = "5.2.1"

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

	if updateAvailable, downloadLink := isAvailableNewVersion(); updateAvailable {
		if prompter.YN("Are you want update book to new version?", false) {
			openBrowser(downloadLink)
			os.Exit(0)
		}
	}

	go func() {
		<-time.After(100 * time.Millisecond)
		if err := openBrowser(fmt.Sprintf("http://localhost:%s", strconv.Itoa(int(*port)))); err != nil {
			log.Fatalln(err)
		}
	}()

	log.Printf("start serve on address %s", address)
	log.Fatalln(http.ListenAndServe(address, logRequest(http.DefaultServeMux)))
}

func isAvailableNewVersion() (bool, string) {
	uc := upd.New("GoFarsi", "book", "Go Programming Language Persian", "https://github.com/GoFarsi/book/releases", 0, false)
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
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
