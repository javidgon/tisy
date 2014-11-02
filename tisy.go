package main

import (
	"fmt"
	"github.com/dchest/uniuri"
	"io/ioutil"
	"net/http"
)

const URLS_FOLDER = "data"

type Url struct {
	Uri       string
	ShortLink string
}

func (u *Url) save() error {
	return ioutil.WriteFile(
		URLS_FOLDER+"/"+u.ShortLink, []byte(u.Uri), 0600)
}

func handler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	if id != "/" {
		body, _ := ioutil.ReadFile(URLS_FOLDER + "/" + id[1:len(id)-1])
		str := string(body)
		http.Redirect(w, r, str, 301)
	} else {
		fmt.Fprintf(w, "<h1>Create shortlink</h1>"+
			"<form action=\"/create/\" method=\"POST\">"+
			"<input type=\"text\" name=\"url\"/><br>"+
			"<input type=\"submit\" value=\"Save\">"+
			"</form>")
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	rawUrl := r.FormValue("url")
	id := uniuri.NewLen(5)
	url := &Url{Uri: rawUrl, ShortLink: id}
	url.save()
	fmt.Fprintf(w, "<h1>Generated shortlink</h1>"+
		"<p>Shortlink: <a href='/%s/'>%s<a/>, Url: %s", url.ShortLink, url.ShortLink, url.Uri)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/create/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
