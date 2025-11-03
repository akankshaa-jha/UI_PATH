package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type Emp struct {
	ID   int
	Name string
	City string
}

var emps = []Emp{
	{1, "Akanksha", "Hyderabad"},
	{2, "Riya", "Pune"},
}

var tpl = template.Must(template.New("index").Parse(`
<!DOCTYPE html><html><body>
<h2>Employee List</h2>
<table border=1><tr><th>ID</th><th>Name</th><th>City</th><th>Action</th></tr>
{{range .}}
<tr>
<td>{{.ID}}</td><td>{{.Name}}</td><td>{{.City}}</td>
<td><a href="/edit?id={{.ID}}">Edit</a> | <a href="/del?id={{.ID}}">Delete</a></td>
</tr>{{end}}</table>

<h3>Add Employee</h3>
<form method="POST" action="/add">
Name:<input name="name"> City:<input name="city">
<input type="submit" value="Add">
</form>
</body></html>
`))

var editTpl = template.Must(template.New("edit").Parse(`
<!DOCTYPE html><html><body>
<h2>Edit Employee</h2>
<form method="POST" action="/update">
<input type="hidden" name="id" value="{{.ID}}">
Name:<input name="name" value="{{.Name}}">
City:<input name="city" value="{{.City}}">
<input type="submit" value="Update">
</form>
</body></html>
`))

func main() {
	http.HandleFunc("/", show)
	http.HandleFunc("/add", add)
	http.HandleFunc("/del", del)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/update", update)
	http.ListenAndServe(":8080", nil)
}

func show(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, emps)
}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := len(emps) + 1
		emps = append(emps, Emp{id, r.FormValue("name"), r.FormValue("city")})
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func del(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	for i, e := range emps {
		if e.ID == id {
			emps = append(emps[:i], emps[i+1:]...)
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func edit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	for _, e := range emps {
		if e.ID == id {
			editTpl.Execute(w, e)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("id"))
		for i := range emps {
			if emps[i].ID == id {
				emps[i].Name = r.FormValue("name")
				emps[i].City = r.FormValue("city")
				break
			}
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
