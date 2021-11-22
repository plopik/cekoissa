package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type question struct {
	words    string
	image    string
	response string
	color    string
}

type serie struct {
	subject string
	label   string
	qs      []string
	as      []string
}

var questionsMap = map[string]question{}
var serieIRM = serie{subject: "irmcrane", label: "IRM craniales"}
var serieCrane = serie{subject: "basecrane", label: "Schéma de la base du crane"}
var serieNerfs = serie{subject: "nerfs", label: "Nerfs craniaux"}
var serieParasite = serie{subject: "parasite", label: "Parasites 1"}
var series = []*serie{&serieIRM, &serieCrane, &serieNerfs, &serieParasite}

func home_template(c *gin.Context) {
	ss := [][]string{}
	for _, s := range series {
		ss = append(ss, []string{s.subject, s.label})
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Series": ss,
	})
}

func main() {

	serieIRM.import_image("radio_cerveau", "#000000")
	serieCrane.import_image("neuro_anat", "#ebebeb")
	serieNerfs.import_csv("data/neuro_anat/nerfs_craniaux.csv")
	serieParasite.import_image("parasite", "#000000")

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.StaticFS("/data", http.Dir("data"))
	r.StaticFile("styles.css", "./templates/styles.css")
	r.StaticFile("home_icon.svg", "./templates/home_black_24dp.svg")

	r.GET("/", home_template)

	for _, s := range series {
		ss := s
		r.GET("/"+s.subject, func(c *gin.Context) { question_template(c, ss) })
	}
	r.Run(":4277")
}
