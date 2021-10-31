package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type question struct {
	image    string
	response string
}

var questions = []string{}
var questionsMap = map[string]question{}
var answers = []string{}

func home_template(c *gin.Context) {
	answer := c.DefaultQuery("answer", "c'est quoi ca ?")
	if answer == "false" {
		q := c.DefaultQuery("question", "bug")
		c.HTML(http.StatusOK, "cekoissaError.html", gin.H{
			"Image":    q,
			"Response": questionsMap[q].response,
		})
	} else {
		header := "c'est quoi ca ?"
		if answer == "true" {
			header = "Vrai"
		}
		index := rand.Intn(len(questions))
		q := questionsMap[questions[index]]
		a := map[string]string{q.response: "answer=true"}
		for len(a) < 5 {
			i := rand.Intn(len(answers))
			fa := answers[i]
			if a[fa] == "" {
				a[answers[i]] = "answer=false&question=" + q.image
			}
		}
		fmt.Println(q.image)
		c.HTML(http.StatusOK, "cekoissa.html", gin.H{
			"Image":   q.image,
			"Header":  header,
			"Answers": a,
		})
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	files, err := ioutil.ReadDir("data/radio_cerveau")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".png") {
			questions = append(questions, "radio_cerveau/"+name)
		}

	}

	for _, q := range questions {
		q2 := strings.Replace(q, ".png", "", -1)
		q2 = strings.Trim(q2, "2")
		q2 = strings.Trim(q2, "3")
		la := strings.Split(q2, "/")
		a := strings.Replace(la[len(la)-1], "_", " ", -1)
		questionsMap[q] = question{q, a}
		if !contains(answers, a) {
			answers = append(answers, a)
		}
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.StaticFS("/data", http.Dir("data"))
	r.StaticFile("styles.css", "./templates/styles.css")

	r.GET("/cekoissa", home_template)
	r.Run()
}
