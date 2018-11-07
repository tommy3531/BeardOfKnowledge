package page

import(
	"net/http"
	"html/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// t, _ := template.ParseFiles("template/home.html")
	t, _ := template.ParseFiles("system/app/template/site.html", "system/app/template/header.html", "system/app/template/footer.html")
	// t.Execute(w, title)
	t.ExecuteTemplate(w, "site", "")
	// w.Write([]byte("Hello from HomePage Function"))
	
}

func ContactPage(w http.ResponseWriter, r *http.Request) {
	title := "Contact Page"
	t, _ := template.ParseFiles("system/app/template/contact.html")
	t.Execute(w, title)
}