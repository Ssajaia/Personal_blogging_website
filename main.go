package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

var articles []Article
const dataFile = "articles.json"
const adminUser = "admin"
const adminPass = "password123"

func main() {
	loadArticles()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/article/", articleHandler)
	http.HandleFunc("/admin/", adminAuthMiddleware(adminDashboardHandler))
	http.HandleFunc("/admin/login", adminLoginHandler)
	http.HandleFunc("/admin/add", adminAuthMiddleware(addArticleHandler))
	http.HandleFunc("/admin/edit/", adminAuthMiddleware(editArticleHandler))
	http.HandleFunc("/admin/delete/", adminAuthMiddleware(deleteArticleHandler))
	http.HandleFunc("/admin/logout", logoutHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadArticles() {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			articles = []Article{}
			return
		}
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &articles); err != nil {
		log.Fatal(err)
	}
}

func saveArticles() {
	data, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(dataFile, data, 0644); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("home").Parse(homeTemplate))
	tmpl.Execute(w, articles)
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/article/")
	for _, article := range articles {
		if article.ID == id {
			tmpl := template.Must(template.New("article").Parse(articleTemplate))
			tmpl.Execute(w, article)
			return
		}
	}
	http.NotFound(w, r)
}

func adminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("dashboard").Parse(dashboardTemplate))
	tmpl.Execute(w, articles)
}

func addArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")
		date := time.Now().Format("2006-01-02")
		id := filepath.Base(title + "-" + date)

		articles = append(articles, Article{
			ID:      id,
			Title:   title,
			Content: content,
			Date:    date,
		})
		saveArticles()
		http.Redirect(w, r, "/admin/", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.New("add").Parse(addTemplate))
	tmpl.Execute(w, nil)
}

func editArticleHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/admin/edit/")
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")

		for i, article := range articles {
			if article.ID == id {
				articles[i].Title = title
				articles[i].Content = content
				saveArticles()
				http.Redirect(w, r, "/admin/", http.StatusSeeOther)
				return
			}
		}
		http.NotFound(w, r)
		return
	}

	for _, article := range articles {
		if article.ID == id {
			tmpl := template.Must(template.New("edit").Parse(editTemplate))
			tmpl.Execute(w, article)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/admin/delete/")
	for i, article := range articles {
		if article.ID == id {
			articles = append(articles[:i], articles[i+1:]...)
			saveArticles()
			http.Redirect(w, r, "/admin/", http.StatusSeeOther)
			return
		}
	}
	http.NotFound(w, r)
}

func adminLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == adminUser && password == adminPass {
			http.SetCookie(w, &http.Cookie{
				Name:  "auth",
				Value: "true",
				Path:  "/",
			})
			http.Redirect(w, r, "/admin/", http.StatusSeeOther)
			return
		}
	}

	tmpl := template.Must(template.New("login").Parse(loginTemplate))
	tmpl.Execute(w, nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func adminAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth")
		if err != nil || cookie.Value != "true" {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}