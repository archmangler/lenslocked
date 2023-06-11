package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site~!</h1>")
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Sorry, that page doesn't exist</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:traiano@gmail.com\">traiano@gmail.com</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	  <ul>
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
		  Email us - <a href="mailto:support@engeneon.com">support@engeneon.com</a>
		</li>
	  </ul>
	  `)
}

type Router struct {
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "You're Sh** out of Luck", http.StatusNotFound)

		//Alternatively:
		//w.WriteHeader(http.StatusNotFound)
		//fmt.Fprint(w, "You're Sh** out of Luck.")

	}

}

func pathHandler(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "You're Sh** out of Luck", http.StatusNotFound)
	}
}

func main() {
	var router Router
	//var router http.HandlerFunc = pathHandler
	//http.HandleFunc("/", homeHandler)
	//http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on: 3001 ...")
	//http.ListenAndServe(":3001", router)
	http.ListenAndServe(":3001", router)

}
