package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type question struct {
	question string
	image    string
	response string
}

var questionsMap = map[string]question{}

var CKSquestions = []string{}
var CKSanswers = []string{}

var WITquestions = []string{}
var WITanswers = []string{}

func home_template(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}
func whoisthat_template(c *gin.Context) {

	answer := c.DefaultQuery("a", "who is that")
	qnumber, _ := strconv.Atoi(c.DefaultQuery("q", "0"))
	if answer == "false" {
		q := c.DefaultQuery("question", "bug")
		c.HTML(http.StatusOK, "whoisthatError.html", gin.H{
			"Question": q,
			"Response": questionsMap[q].response,
			"Next":     "q=" + strconv.Itoa(qnumber+1),
		})
	} else {
		header := "who is that ?"
		if answer == "true" {
			header = "Vrai"
		}
		q := questionsMap[WITquestions[qnumber]]
		var a [][]string
		a = append(a, []string{q.response, "a=true&q=" + strconv.Itoa(qnumber+1)})
		for len(a) < 5 {
			i := rand.Intn(len(WITanswers))
			fa := WITanswers[i]
			if !contains2(a, fa) {
				a = append(a, []string{fa, "a=false&question=" + q.image + "&q=" + strconv.Itoa(qnumber)})
			}
		}
		c.HTML(http.StatusOK, "whoisthat.html", gin.H{
			"Question": q.question,
			"Header":   header,
			"Answers":  a,
		})
	}
}

func cekoissa_template(c *gin.Context) {
	answer := c.DefaultQuery("a", "c'est quoi ca ?")
	qnumber, _ := strconv.Atoi(c.DefaultQuery("q", "0"))
	if answer == "false" {
		q := c.DefaultQuery("question", "bug")
		c.HTML(http.StatusOK, "cekoissaError.html", gin.H{
			"Image":    q,
			"Response": questionsMap[q].response,
			"Next":     "q=" + strconv.Itoa(qnumber+1),
		})
	} else {
		header := "c'est quoi ca ?"
		if answer == "true" {
			header = "Vrai"
		}
		q := questionsMap[CKSquestions[qnumber]]
		var a [][]string
		a = append(a, []string{q.response, "a=true&q=" + strconv.Itoa(qnumber+1)})
		for len(a) < 5 {
			i := rand.Intn(len(CKSanswers))
			fa := CKSanswers[i]
			if !contains2(a, fa) {
				a = append(a, []string{fa, "a=false&question=" + q.image + "&q=" + strconv.Itoa(qnumber)})
			}
		}
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

func contains2(s [][]string, e string) bool {
	for _, a := range s {
		if a[0] == e {
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
			CKSquestions = append(CKSquestions, "radio_cerveau/"+name)
		}
	}

	for _, q := range CKSquestions {
		q2 := strings.Replace(q, ".png", "", -1)
		q2 = strings.Trim(q2, "2")
		q2 = strings.Trim(q2, "3")
		la := strings.Split(q2, "/")
		a := strings.Replace(la[len(la)-1], "_", " ", -1)
		questionsMap[q] = question{"", q, a}
		if !contains(CKSanswers, a) {
			CKSanswers = append(CKSanswers, a)
		}
	}

	WITquestions = append(WITquestions, "test")
	questionsMap["test"] = question{"quel est le muscle ?", "", "c'est le con"}
	WITanswers = CKSanswers

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.StaticFS("/data", http.Dir("data"))
	r.StaticFile("styles.css", "./templates/styles.css")

	r.GET("/cekoissa", cekoissa_template)
	r.GET("/", home_template)
	r.GET("/whoisthat", whoisthat_template)
	r.Run(":4277")
}
