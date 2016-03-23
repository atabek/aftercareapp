package main

import (
	"html/template"
	"log"
	"net/http"
	r "github.com/dancannon/gorethink"
	"encoding/json"
)


// Struct tags are used to map struct fields to fields in the database
type Person struct {
    Id    string `gorethink:"id,omitempty"`
    Name  string `gorethink:"name"`
    Place string `gorethink:"place"`
}

var tpl *template.Template


func init() {
    var err error
    session, err = r.Connect(r.ConnectOpts{
        Address:  "localhost:28015",
        Database: "test",
    })
    if err != nil {
        fmt.Println(err)
        return
    }
}

func main() {
	tpl = template.Must(template.ParseGlob("templates/html/*.html"))
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	tpl.ExecuteTemplate(res, "home.html", nil)
}
