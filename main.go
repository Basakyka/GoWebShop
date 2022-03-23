package main

import (
	"html/template"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("HTML/home_page.html")
	tmpl.Execute(w, "")
	expire := time.Now().Add(20 * time.Minute)
	cookie := http.Cookie{Name: "username", Value: "nonsecureuser", Path: "/", Expires: expire, MaxAge: 86400}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{Name: "secureusername", Value: "secureuser", Path: "/", Expires: expire, MaxAge: 86400, HttpOnly: true, Secure: true}
	http.SetCookie(w, &cookie)
	if err != nil {
		println(err.Error())
	}
}

func shopping_cart(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("HTML/shopping_cart.html")
	tmpl.Execute(w, "")
	if err != nil {
		println(err.Error())
	}
}

func credits(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("HTML/about_page.html")
	tmpl.Execute(w, "")
	if err != nil {
		println(err.Error())
	}
}

func catalog(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("HTML/catalog.html")
	tmpl.Execute(w, "")
	if err != nil {
		println(err.Error())
	}
}

func handleRequest() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", home_page)
	http.HandleFunc("/shopping_cart", shopping_cart)
	http.HandleFunc("/catalog", catalog)
	http.HandleFunc("/credits", credits)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println(err.Error())
	}
}

func main() {
	handleRequest()
}
