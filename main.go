package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type PageModel struct {
	Title   string
	Content template.HTML
	LastMod time.Time
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getDir).Methods("GET")
	router.HandleFunc("/{path}", getPage).Methods("GET")
	router.HandleFunc("/ws/{path}", serveWs)

	log.Println("md-live-server web server running")
	log.Print("port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getDir(w http.ResponseWriter, r *http.Request) {
	elements := readDir("./")
	var content string
	content = `<ul>`
	for _, elem := range elements {
		content += `
		<li><a href="` + elem + `">` + elem + `</a></li>
		`
	}
	content += `</ul>`
	var page PageModel
	page.Title = "dir"
	page.Content = template.HTML(content)

	tmplPage := template.Must(template.New("t").Parse(dirTemplate))
	tmplPage.Execute(w, page)
}

func getPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := vars["path"]
	path = strings.Replace(path, "%", "/", -1)
	log.Println(path)

	if strings.Split(path, ".")[1] != "md" {
		http.ServeFile(w, r, path)
	}

	content, err := fileToHTML(path)
	check(err)

	var page PageModel
	page.Title = path
	page.Content = template.HTML(content)

	tmplPage := template.Must(template.New("t").Parse(htmlTemplate))
	tmplPage.Execute(w, page)
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := vars["path"]
	path = strings.Replace(path, "%", "/", -1)
	log.Println("websocket", path)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	// watch file
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					writer(ws, path)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func writer(ws *websocket.Conn, path string) {
	content, err := fileToHTML(path)
	check(err)
	if err := ws.WriteMessage(websocket.TextMessage, []byte(content)); err != nil {
		return
	}
}
