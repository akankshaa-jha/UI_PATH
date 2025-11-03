package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Create a simple HTML template for the contact form
var tmpl = template.Must(template.New("form").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>Contact Form</title>
</head>
<body>
	<h2>Contact Us</h2>
	<form method="POST" action="/submit">
		Name: <input type="text" name="name"><br><br>
		Email: <input type="email" name="email"><br><br>
		Message:<br>
		<textarea name="message" rows="5" cols="30"></textarea><br><br>
		<input type="submit" value="Send">
	</form>
</body>
</html>
`))

func main() {
	// Route 1 — show form
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	// Route 2 — handle form submission
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		fmt.Fprintf(w, "<h3>Thank you for contacting us!</h3>")
		fmt.Fprintf(w, "<p><strong>Name:</strong> %s</p>", name)
		fmt.Fprintf(w, "<p><strong>Email:</strong> %s</p>", email)
		fmt.Fprintf(w, "<p><strong>Message:</strong> %s</p>", message)
	})

	// Start server
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
