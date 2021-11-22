package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

func (s *serie) import_image(folder string, color string) {
	files, err := ioutil.ReadDir("data/" + folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".png") || strings.HasSuffix(name, ".jpg") || strings.HasSuffix(name, ".GIF") || strings.HasSuffix(name, ".gif") {
			s.qs = append(s.qs, folder+"/"+name)
		}

	}
	rand.Shuffle(len(s.qs), func(i, j int) { s.qs[i], s.qs[j] = s.qs[j], s.qs[i] })

	for _, q := range s.qs {
		q2 := strings.Replace(q, ".png", "", -1)
		q2 = strings.Replace(q2, ".jpg", "", -1)
		q2 = strings.Replace(q2, ".GIF", "", -1)
		q2 = strings.Replace(q2, ".gif", "", -1)
		q2 = strings.Trim(q2, "2")
		q2 = strings.Trim(q2, "3")
		q2 = strings.ToLower(q2)
		q2 = strings.Title(q2)
		la := strings.Split(q2, "/")
		a := strings.Replace(la[len(la)-1], "_", " ", -1)
		questionsMap[q] = question{"", q, a, color}
		if !contains(s.as, a) {
			s.as = append(s.as, a)
		}
	}
}

func (s *serie) import_csv(file string) {
	records := readCsvFile(file)
	sentence := records[0][1]
	for i, line := range records {
		if i == 0 {
			continue
		}
		q := line[0]
		a := line[1]

		questionsMap[q] = question{sentence + " " + q + " ?", "", a, ""}
		s.qs = append(s.qs, q)
		if !contains(s.as, a) {
			s.as = append(s.as, a)
		}
	}
	rand.Shuffle(len(s.qs), func(i, j int) { s.qs[i], s.qs[j] = s.qs[j], s.qs[i] })
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

func readCsvFile(filePath string) [][]string {
	f, _ := os.Open(filePath)
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	records, _ := csvReader.ReadAll()

	return records
}
