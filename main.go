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
	qs      []string
	as      []string
}

var questionsMap = map[string]question{}
var serieIRM = serie{subject: "irmcrane"}
var serieCrane = serie{subject: "basecrane"}
var serieNerfs = serie{subject: "nerfs"}
var serieParasite = serie{subject: "parasite"}

func home_template(c *gin.Context) {
	series := [][]string{
		{"irmcrane", "IRM craniales"},
		{"basecrane", "Schema de la base du crane"},
		{"nerfs", "Nerfs craniaux"},
		{"parasite", "Parasites"},
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Series": series,
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

	r.GET("/", home_template)
	r.GET("/nerfs", func(c *gin.Context) { question_template(c, serieNerfs) })
	r.GET("/irmcrane", func(c *gin.Context) { question_template(c, serieIRM) })
	r.GET("/basecrane", func(c *gin.Context) { question_template(c, serieCrane) })
	r.GET("/parasite", func(c *gin.Context) { question_template(c, serieParasite) })
	r.Run(":4277")
}
