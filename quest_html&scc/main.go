package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

)

type Note struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
}

var tpl = template.Must(template.ParseFiles("/Users/anastasiasakuta/Desktop/web quest/GoPet/quest_html&scc/mainpage.html"))


func createDB() {
	pgxpool.Connect()
}
func showNote(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("<h1> Hello World</h1>"))
	tpl.Execute(w, nil)
}

func connectDB() {

}
func createNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод запрещен", 405)
		return
	}
	var bytes []byte
	var note Note
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&note)
	if err != nil {
		log.Fatal("Can't decode: ", err)
	}
	fmt.Println(bytes)
	
	w.Write([]byte("Создание новой заметки..."))
}
func main() {
	fmt.Println("hello!")

	//fs := http.FileServer(http.Dir("./styles"))

	mux := http.NewServeMux()

	//mux.Handle("/", fs)
	mux.HandleFunc("/showNote", showNote)
	mux.HandleFunc("/createNote", createNote)
	http.ListenAndServe(":80", mux)
}
