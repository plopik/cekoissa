package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func error_words_template(c *gin.Context, s *serie, qnumber int, q question) {
	answer := c.Query("a")
	header := "Faux"
	headerColor := "#ff0000"
	if answer == "nil" {
		header = "Réponse"
		headerColor = "#0000ff"
	}
	c.HTML(http.StatusOK, "words_error.html", gin.H{
		"Header":      header,
		"Headercolor": headerColor,
		"Question":    q.words,
		"Response":    q.response,
		"Next":        s.subject + "?q=" + strconv.Itoa(qnumber+1),
	})
}

func words_template(c *gin.Context, s *serie, qnumber int, q question) {
	header := "A ton avis ?"

	var a [][]string
	a = append(a, []string{q.response, s.subject + "?a=true&q=" + strconv.Itoa(qnumber+1), "#ffffff"})
	for len(a) < 4 {
		i := rand.Intn(len(s.as))
		fa := s.as[i]
		if !contains2(a, fa) {
			a = append(a, []string{fa, s.subject + "?a=false&q=" + strconv.Itoa(qnumber), "#ffffff"})
		}
	}
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	a = append(a, []string{"sais pas", s.subject + "?a=nil&q=" + strconv.Itoa(qnumber), "#0000ff"})
	c.HTML(http.StatusOK, "words_question.html", gin.H{
		"Question": q.words,
		"Header":   header,
		"Answers":  a,
	})

}

func error_image_template(c *gin.Context, s *serie, qnumber int, q question) {
	answer := c.Query("a")
	header := "Faux"
	headerColor := "#ff0000"
	if answer == "nil" {
		header = "Réponse"
		headerColor = "#0000ff"
	}
	c.HTML(http.StatusOK, "image_error.html", gin.H{
		"Header":      header,
		"Headercolor": headerColor,
		"Imagecolor":  q.color,
		"Image":       q.image,
		"Response":    q.response,
		"Next":        s.subject + "?q=" + strconv.Itoa(qnumber+1),
	})
}

func image_template(c *gin.Context, s *serie, qnumber int, q question) {
	header := "c'est quoi ca ?"

	var a [][]string
	a = append(a, []string{q.response, s.subject + "?a=true&q=" + strconv.Itoa(qnumber+1), "#ffffff"})
	for len(a) < 4 {
		i := rand.Intn(len(s.as))
		fa := s.as[i]
		if !contains2(a, fa) {
			a = append(a, []string{fa, s.subject + "?a=false&q=" + strconv.Itoa(qnumber), "#ffffff"})
		}
	}
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	a = append(a, []string{"sais pas", s.subject + "?a=nil&q=" + strconv.Itoa(qnumber), "#0000ff"})
	c.HTML(http.StatusOK, "image_question.html", gin.H{
		"Imagecolor": q.color,
		"Image":      q.image,
		"Header":     header,
		"Answers":    a,
	})
}

func question_template(c *gin.Context, s *serie) {
	answer := c.Query("a")
	qnumber, _ := strconv.Atoi(c.DefaultQuery("q", "0"))
	if qnumber >= len(s.qs) {
		qnumber = 0
	}
	q := questionsMap[s.qs[qnumber]]
	if (answer == "false" || answer == "nil") && q.image != "" {
		error_image_template(c, s, qnumber, q)
	} else if (answer == "false" || answer == "nil") && q.image == "" {
		error_words_template(c, s, qnumber, q)
	} else if q.image != "" {
		image_template(c, s, qnumber, q)
	} else {
		words_template(c, s, qnumber, q)
	}
}
