package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func end_template(c *gin.Context, s *serie) {
	c.HTML(http.StatusOK, "end_serie.html", gin.H{
		"Next": s.subject,
	})
}

func error_words_template(c *gin.Context, s *serie, qnumber int, q question) {
	answer := c.Query("a")
	header := "Faux"
	headerColor := "#ff0000"
	if answer == "nil" {
		header = "Réponse"
		headerColor = "#0000ff"
	}
	qhtml := [][]template.HTML{}
	for _, words := range q.words {
		wordshtml := []template.HTML{}
		for _, word := range words {
			wordshtml = append(wordshtml, template.HTML(strings.ReplaceAll(word, "\n", "<br>")))
		}
		qhtml = append(qhtml, wordshtml)
	}
	c.HTML(http.StatusOK, "words_error.html", gin.H{
		"Header":      header,
		"Headercolor": headerColor,
		"Questions":   qhtml,
		"Response":    q.response,
		"Next":        s.subject + "?q=" + strconv.Itoa(qnumber+1),
	})
}

func words_template(c *gin.Context, s *serie, qnumber int, q question) {
	header := "Qui correspond le mieux ?"

	var a [][]string
	a = append(a, []string{q.response, s.subject + "?a=true&q=" + strconv.Itoa(qnumber+1), "button"})
	for len(a) < 4 {
		i := rand.Intn(len(s.as))
		fa := s.as[i]
		if !contains2(a, fa) {
			a = append(a, []string{fa, s.subject + "?a=false&q=" + strconv.Itoa(qnumber), "button"})
		}
	}
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	a = append(a, []string{"sais pas", s.subject + "?a=nil&q=" + strconv.Itoa(qnumber), "bluebutton"})
	qhtml := [][]template.HTML{}
	for _, words := range q.words {
		wordshtml := []template.HTML{}
		for _, word := range words {
			wordshtml = append(wordshtml, template.HTML(strings.ReplaceAll(word, "\n", "<br>")))
		}
		qhtml = append(qhtml, wordshtml)
	}
	c.HTML(http.StatusOK, "words_question.html", gin.H{
		"Counter":   fmt.Sprintf("%v/%v", qnumber+1, len(s.qs)),
		"Questions": qhtml,
		"Header":    header,
		"Answers":   a,
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
	header := "C'est quoi ca ?"

	var a [][]string
	a = append(a, []string{q.response, s.subject + "?a=true&q=" + strconv.Itoa(qnumber+1), "button"})
	for len(a) < 4 {
		i := rand.Intn(len(s.as))
		fa := s.as[i]
		if !contains2(a, fa) {
			a = append(a, []string{fa, s.subject + "?a=false&q=" + strconv.Itoa(qnumber), "button"})
		}
	}
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	a = append(a, []string{"Sais pas", s.subject + "?a=nil&q=" + strconv.Itoa(qnumber), "bluebutton"})
	c.HTML(http.StatusOK, "image_question.html", gin.H{
		"Counter":    fmt.Sprintf("%v/%v", qnumber+1, len(s.qs)),
		"Imagecolor": q.color,
		"Image":      q.image,
		"Header":     header,
		"Answers":    a,
	})
}

func question_template(c *gin.Context, s *serie) {
	answer := c.Query("a")
	qnumber, _ := strconv.Atoi(c.DefaultQuery("q", "0"))
	qlast := qnumber - 1
	end := false
	if qnumber >= len(s.qs) {
		qnumber = 0
		end = true
	}

	q := questionsMap[s.qs[qnumber]]
	qq := s.qs[qnumber]
	ip := c.ClientIP()
	if answer == "true" {
		qq = s.qs[qlast]
	}
	if answer != "" {
		fmt.Printf("LOG ip=%v s=%v q=%v a=%v\n", ip, s.subject, qq, answer)
	}
	if end {
		end_template(c, s)
	} else if (answer == "false" || answer == "nil") && q.image != "" {
		error_image_template(c, s, qnumber, q)
	} else if (answer == "false" || answer == "nil") && q.image == "" {
		error_words_template(c, s, qnumber, q)
	} else if q.image != "" {
		image_template(c, s, qnumber, q)
	} else {
		words_template(c, s, qnumber, q)
	}
}
