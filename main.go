package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type question struct {
	words         [][]string
	image         string
	response      string
	falseResponse []string
	color         string
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
var seriePharmaco = serie{subject: "pharmaco", label: "Pharmacologie", date: "11/12"}
var serieOphtalmo = serie{subject: "ophtalmo", label: "Ophtalmologie", date: "12/12"}
var serieThalamus = serie{subject: "thalamus", label: "Thalamus", date: "26/12"}
var serieImagerie = serie{subject: "imagerie", label: "Imagerie", date: "30/12"}
var series = []*serie{&serieIRM, &serieCrane, &serieNerfs, &serieParasite, &serieBacterio, &seriePharmaco, &serieOphtalmo, &serieThalamus, &serieImagerie}

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
	serieOphtalmo.import_image("ophtalmo", "#ebebeb")
	serieParasite.import_image("parasite", "#000000")
	serieThalamus.import_image("thalamus", "#ebebeb")
	serieImagerie.import_image("imagerie", "#000000")

	serieNerfs.import_csv("data/neuro_anat/nerfs_craniaux.csv")
	serieBacterio.import_csv2("data/bacterio.csv")
	seriePharmaco.import_xlsx("data/pharmaco.xlsx")

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.StaticFS("/data", http.Dir("data"))
	r.StaticFile("styles.css", "./templates/styles.css")
	r.StaticFile("home_icon.svg", "./templates/home_black_24dp.svg")
	r.StaticFile("end.gif", "./templates/end.gif")
	r.StaticFile("end2.mp4", "./templates/end2.mp4")

	r.GET("/", home_template)

	for _, s := range series {
		ss := s
		r.GET("/"+s.subject, func(c *gin.Context) { question_template(c, ss) })
	}
	fmt.Println("Ready")
	r.Run(":4277")

}
