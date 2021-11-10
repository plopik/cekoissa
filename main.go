package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type question struct {
	words    string
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

func error_words_template(c *gin.Context, subject string, qnumber int, q question) {
	c.HTML(http.StatusOK, "words_error.html", gin.H{
		"Question": q.words,
		"Response": q.response,
		"Next":     subject + "?q=" + strconv.Itoa(qnumber+1),
	})
}

func words_template(c *gin.Context, subject string, qnumber int, q question) {
	header := "who is that ?"
	if qnumber == 0 {
		header = "Vrai"
	}
	var a [][]string
	a = append(a, []string{q.response, subject + "?a=true&q=" + strconv.Itoa(qnumber+1)})
	for len(a) < 5 {
		i := rand.Intn(len(WITanswers))
		fa := WITanswers[i]
		if !contains2(a, fa) {
			a = append(a, []string{fa, subject + "?a=false&q=" + strconv.Itoa(qnumber)})
		}
	}
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	c.HTML(http.StatusOK, "words_question.html", gin.H{
		"Question": q.words,
		"Header":   header,
		"Answers":  a,
	})

}

func error_image_template(c *gin.Context, subject string, qnumber int, q question) {
	c.HTML(http.StatusOK, "image_error.html", gin.H{
		"Image":    q.image,
		"Response": q.response,
		"Next":     subject + "?q=" + strconv.Itoa(qnumber+1),
	})
}

func image_template(c *gin.Context, subject string, qnumber int, q question) {
	header := "c'est quoi ca ?"
	if qnumber == 0 {
		header = "Vrai"
	}
	var a [][]string
	a = append(a, []string{q.response, subject + "?a=true&q=" + strconv.Itoa(qnumber+1)})
	for len(a) < 5 {
		i := rand.Intn(len(CKSanswers))
		fa := CKSanswers[i]
		if !contains2(a, fa) {
			a = append(a, []string{fa, subject + "?a=false&q=" + strconv.Itoa(qnumber)})
		}
	}
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	c.HTML(http.StatusOK, "image_question.html", gin.H{
		"Image":   q.image,
		"Header":  header,
		"Answers": a,
	})
}

func question_template(c *gin.Context, subject string, qs []string, as []string) {
	answer := c.Query("a")
	qnumber, _ := strconv.Atoi(c.DefaultQuery("q", "0"))
	if qnumber >= len(qs) {
		qnumber = 0
	}
	q := questionsMap[qs[qnumber]]
	if answer == "false" && q.image != "" {
		error_image_template(c, subject, qnumber, q)
	} else if answer == "false" && q.image == "" {
		error_words_template(c, subject, qnumber, q)
	} else if q.image != "" {
		image_template(c, subject, qnumber, q)
	} else {
		words_template(c, subject, qnumber, q)
	}
}

func cekoissa_template(c *gin.Context) {
	question_template(c, "cekoissa", CKSquestions, CKSanswers)
}

func whoisthat_template(c *gin.Context) {
	question_template(c, "whoisthat", WITquestions, WITanswers)
}

func main() {

	CKSquestions, CKSanswers = import_image("radio_cerveau", CKSquestions, CKSanswers)
	WITquestions, WITanswers = import_image("neuro_anat", WITquestions, WITanswers)
	WITquestions, WITanswers = import_csv("data/neuro_anat/nerfs_craniaux.csv", WITquestions, WITanswers)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.StaticFS("/data", http.Dir("data"))
	r.StaticFile("styles.css", "./templates/styles.css")

	r.GET("/cekoissa", cekoissa_template)
	r.GET("/", home_template)
	r.GET("/whoisthat", whoisthat_template)
	r.Run(":4277")
}
