package main

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/dchest/captcha"
)

//go:embed tpl.html
var formTemplateSrc string

var formTemplate = template.Must(template.New("example").Parse(formTemplateSrc))

type Captcha struct {
	CaptchaID string
}

func showFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	d := Captcha{captcha.New()}
	if err := formTemplate.Execute(w, &d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func processFormHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if !captcha.VerifyString(r.FormValue("captchaId"), r.FormValue("captchaSolution")) {
		_, _ = io.WriteString(w, "Wrong captcha solution! No robots allowed!\n")
	} else {
		_, _ = io.WriteString(w, "Great job, human! You solved the captcha.\n")
	}
	_, _ = io.WriteString(w, "<br><a href='/'>Try another one</a>")
}

func main() {
	http.HandleFunc("/", showFormHandler)
	http.HandleFunc("/process", processFormHandler)
	http.Handle("/captcha/", captcha.Server(captcha.StdWidth, captcha.StdHeight))
	fmt.Println("Server is at localhost:8666")
	if err := http.ListenAndServe(":8666", nil); err != nil {
		log.Fatal(err)
	}
}
