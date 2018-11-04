package page

import(
	"net/http"
	"html/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	test := "Home Page"
	t, _ := template.ParseFiles("template/home.html")
	t.Execute(w, test)
	
}