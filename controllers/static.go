package controllers

import (
	"html/template"
	"net/http"

	"web-app.com/simple/views"
)

// Static handles static page routes
type Static struct {
	HomeTemplate    views.Template
	ContactTemplate views.Template
	FaqTemplate     views.Template
}

func (s *Static) Home(w http.ResponseWriter, r *http.Request) {
	s.HomeTemplate.Execute(w, r, nil)
}

func (s *Static) Contact(w http.ResponseWriter, r *http.Request) {
	s.ContactTemplate.Execute(w, r, nil)
}

func (s *Static) Faq(w http.ResponseWriter, r *http.Request) {
	s.FaqTemplate.Execute(w, r, faqData)
}

// FAQ data - defined once at package level
var faqData = []struct {
	Question string
	Answer   template.HTML
}{
	{Question: "Is there a free version?", Answer: "Yes! We offer a free trial for 30 days."},
	{Question: "What are your support hours?", Answer: "We have support 24/7."},
	{Question: "How do I contact support?", Answer: `Email <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>`},
	{Question: "Where is your office?", Answer: "Our team is fully remote!"},
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
