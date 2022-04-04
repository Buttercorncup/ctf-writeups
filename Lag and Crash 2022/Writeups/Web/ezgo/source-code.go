

package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"fmt"
)

type Page struct {
	Title string
	Body string
}


func (u Page) Getflag() string {
	flag := "LNC2022{The_flag_is_here}"
	return string(flag)
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, []byte(p.Body), 0600)
}

func loadPage(title string, w http.ResponseWriter) (p *Page, err error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: string(body)}, nil
}


func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title, w)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title, w)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: body}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}


func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {

	switch {
	case tmpl == "view":
		var tmp = fmt.Sprintf(`<h1>{{.Title}}</h1>
		<p>[<a href="/edit/{{.Title}}">edit</a>]</p>
		<div>%s</div>
		`, p.Body)

		t, err := template.New("page").Parse(tmp)

		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		t.Execute(w, p)
	case tmpl == "edit":
		var tmp = fmt.Sprintf(`
		<h1>Editing {{.Title}}</h1>
		<form action="/save/{{.Title}}" method="POST">
		<div><textarea name="body" rows="20" cols="80">%s</textarea></div>
		<div><input type="submit" value="Save"></div>
		</form>`, p.Body)
		
		t, err := template.New("page").Parse(tmp)

		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		t.Execute(w, p)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	
	fmt.Println("[================================]")
	fmt.Println("Please access http://localhost:8080/view/a_random_string") 
	fmt.Println("For example: http://localhost:8080/view/ANewPage")
	fmt.Println("if there's any error, just change the random_string")
	fmt.Println("[================================]")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}