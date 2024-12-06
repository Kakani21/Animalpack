package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h>Welcome to the world of Subbu</h>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Contact Page</h1><p>To get in touch, email me at <a 
	href="mailto:Kakani2789@gmail.com">Kakani2789@gmail.com</a>.</p>`)

}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page </h1>
	<u1>
	<li>
	<b>Is there a free version?</b>
	Yes! We offer a free trial for 30 days on any paid plans.
	</li>
	<li>
	<b>What are your support hours?</b>
	We have support staff answering emails 24/7, though response
    times may be a bit slower on weekends.
	</li>
	<li>
	<b>How do I contact support?</b>
	Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
	</li>
	</u1>`)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {

// 	/*if r.URL.Path == "/" {
// 		homeHandler(w, r)
// 	} else if r.URL.Path == "/contact" {
// 		contactHandler(w, r)
// 	}*/
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/FAQ":
// 		FAQHandler(w, r)
// 	default:
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprint(w, "Page not found")
// 	}
// }

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprint(w, "Page not found")
// 	}
// }

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on: 3000...")
	http.ListenAndServe(":3000", r)
}

// func main() {
// 	fmt.Println("Starting the server on: 3000...")
// 	http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
// }

// func main() {
// 	http.HandleFunc("/", pathHandler)
// 	//http.HandleFunc("/contact", contactHandler)
// 	//fmt.Fprintln(os.Stdout, "Hello World!")
// 	fmt.Println("Starting the server on: 3000...")
// 	http.ListenAndServe(":3000", nil)

// }
