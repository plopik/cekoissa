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
var CKS = serie{subject: "irmcrane"}
var WIT = serie{subject: "basecrane"}
var serieNerf = serie{subject: "nerfs"}

func home_template(c *gin.Context) {
	series := [][]string{{"irmcrane", "IRM craniales"}, {"basecrane", "Schema de la base du crane"}, {"nerfs", "Nerfs craniaux"}}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Series": series,
	})
}

func main() {

	CKS.import_image("radio_cerveau", "#000000")
	WIT.import_image("neuro_anat", "#ebebeb")
	serieNerf.import_csv("data/neuro_anat/nerfs_craniaux.csv")

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.StaticFS("/data", http.Dir("data"))
	r.StaticFile("styles.css", "./templates/styles.css")

	r.GET("/", home_template)
	r.GET("/nerfs", func(c *gin.Context) { question_template(c, serieNerf) })
	r.GET("/irmcrane", func(c *gin.Context) { question_template(c, CKS) })
	r.GET("/basecrane", func(c *gin.Context) { question_template(c, WIT) })
	r.Run(":4277")
}
