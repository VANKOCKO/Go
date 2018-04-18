package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/**
         Model
	**/

type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

type Todos []Todo

// propriete du server

type Server struct {
	ServerName string
	ServerIp   string
}

// liste des serveurs
type Servers []Server

// list des personnes

type Person struct {
	UserName string
	Emails   []string
}

func main() {

	/**
	  Routage
	  **/
	router := mux.NewRouter().SkipClean(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/listserver", ListServer)
	log.Fatal(http.ListenAndServe(":8080", router))
}

/**
       Controller
	   **/

func Index(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, "Hello World!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

	todos := Todos{
		Todo{Name: "Ecrire la presentation ", Completed: true},
		Todo{Name: "Hotel"},
	}

	json.NewEncoder(w).Encode(todos)

}

func ListServer(w http.ResponseWriter, r *http.Request) {

	servers := Servers{
		Server{ServerName: "sv1", ServerIp: "127.0.01"},
		Server{ServerName: "sv2", ServerIp: "127.0.02"},
		Server{ServerName: "sv2", ServerIp: "127.0.03"},
		Server{ServerName: "sv3", ServerIp: "127.0.04"},
	}
	json.NewEncoder(w).Encode(servers)

}
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "todo show", todoId)
}
