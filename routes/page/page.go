package page

import(
	"net/http"
	"html/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// t, _ := template.ParseFiles("template/home.html")
	t, _ := template.ParseFiles("template/site.html", "template/header.html", "template/footer.html")
	// t.Execute(w, title)
	t.ExecuteTemplate(w, "site", "")
	
}

func ContactPage(w http.ResponseWriter, r *http.Request) {
	title := "Contact Page"
	t, _ := template.ParseFiles("template/contact.html")
	t.Execute(w, title)
}