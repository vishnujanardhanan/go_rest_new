package main

//each and every file in go must have a package name
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()
	c := cors.AllowAll()
	handler := c.Handler(router)
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	//	handler := http.NewServeMux()

	fmt.Print("here-----")
	///we create a new router to expose our api
	//to our users
	http.HandleFunc("/api/httpmain", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("here too----")
		//notice how this function takes two parameters
		//the first parameter is a ResponseWriter writer and
		//this is where we write the response we want to send back to the user
		//in this case Hello world
		//the second parameter is a pointer of type  http.Request this holds
		//all information of the request sent by the user
		//this may include query parameters,path parameters and many more
		fmt.Fprintf(w, `Hello world`)
		var resByte []byte
		resByte, _ = json.Marshal("hello world")
		w.Header().Set("Content_Type", "application/json")
		w.Write(resByte)
	})
	//Every time a  request is sent to the endpoint ("/api/hello")
	//the function SayHello will be invoked
	//	http.ListenAndServe("0.0.0.0", handler)
	//we tell our api to listen to all request to port 8080.
	if err := http.ListenAndServe(":"+port, trailingSlashHandler(handler)); err != nil {
		log.Fatalf("unable to start http server, %s", err)
	}
}

func trailingSlashHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Print("here too----")
	//notice how this function takes two parameters
	//the first parameter is a ResponseWriter writer and
	//this is where we write the response we want to send back to the user
	//in this case Hello world
	//the second parameter is a pointer of type  http.Request this holds
	//all information of the request sent by the user
	//this may include query parameters,path parameters and many more
	fmt.Fprintf(w, `Hello world`)
	var resByte []byte
	resByte, _ = json.Marshal("hello world")
	w.Header().Set("Content_Type", "application/json")
	w.Write(resByte)
}
