package main

import (
	"net/http"
	"log"
	"fmt"
	// "html/template"
	// "github.com/gorilla/mux"
	"github.com/tommarler/Beard_Of_knowledge/routes/page"
)

// func hdl(t *template.Template) http.Handler {
// 	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if err := t.Execute(w, r.URL.Query()); err !=nil{
// 			http.Error(w, fmt.Sprintf("error executing template (%s)", err), http.StatusInternalServerError)
// 		}
// 	})
// }



func main() {
	fmt.Println("Hello from main")
	http.HandleFunc("/home", page.HomePage)
	// tpl := template.Must(template.New("site.html").ParseGlob("templates/*.html"))
	// log.Fatal(http.ListenAndServe(":8080", hdl(tpl)))
	log.Fatal(http.ListenAndServe(":8080", nil))

	
}








