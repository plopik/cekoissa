package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type question struct {
	words    [][]string
	image    string
	response string
	color    string
}

type serie struct {
	subject string
	label   string
	date    string
	qs      []string
	as      []string
}

var questionsMap = map[string]question{}
var serieIRM = serie{subject: "irmcrane", label: "IRM craniales", date: "01/11"}
var serieCrane = serie{subject: "basecrane", label: "Schéma de la base du crane", date: "01/11"}
var serieNerfs = serie{subject: "nerfs", label: "Nerfs craniaux", date: "01/11"}
var serieParasite = serie{subject: "parasite", label: "Parasitologie", date: "04/12"}
var serieBacterio = serie{subject: "bacterio", label: "Bactériologie", date: "28/11"}
var series = []*serie{&serieIRM, &serieCrane, &serieNerfs, &serieParasite, &serieBacterio}

func home_template(c *gin.Context) {
	ss := [][]string{}
	for _, s := range series {
		ss = append(ss, []string{s.subject, s.label + " - " + s.date})
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Series": ss,
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())

	serieIRM.import_image("radio_cerveau", "#000000")
	serieCrane.import_image("neuro_anat", "#ebebeb")
	serieNerfs.import_csv("data/neuro_anat/nerfs_craniaux.csv")
	serieParasite.import_image("parasite", "#000000")
	serieBacterio.import_csv2("data/bacterio.csv")

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.StaticFS("/data", http.Dir("data"))
	r.StaticFile("styles.css", "./templates/styles.css")
	r.StaticFile("home_icon.svg", "./templates/home_black_24dp.svg")
	r.StaticFile("end.gif", "./templates/end.gif")

	r.GET("/", home_template)

	for _, s := range series {
		ss := s
		r.GET("/"+s.subject, func(c *gin.Context) { question_template(c, ss) })
	}
	fmt.Println("Ready")
	r.Run(":4277")

}
