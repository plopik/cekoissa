package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

func import_image(folder string, qs []string, as []string) ([]string, []string) {
	files, err := ioutil.ReadDir("data/" + folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".png") {
			qs = append(qs, folder+"/"+name)
		}

	}
	rand.Shuffle(len(qs), func(i, j int) { qs[i], qs[j] = qs[j], qs[i] })

	for _, q := range qs {
		q2 := strings.Replace(q, ".png", "", -1)
		q2 = strings.Trim(q2, "2")
		q2 = strings.Trim(q2, "3")
		la := strings.Split(q2, "/")
		a := strings.Replace(la[len(la)-1], "_", " ", -1)
		questionsMap[q] = question{"", q, a}
		if !contains(as, a) {
			as = append(as, a)
		}
	}
	return qs, as
}

func import_csv(file string, qs, as []string) ([]string, []string) {
	records := readCsvFile(file)
	sentence := records[0][1]
	for i, line := range records {
		if i == 0 {
			continue
		}
		q := line[0]
		a := line[1]

		questionsMap[q] = question{sentence + " " + q + " ?", "", a}
		qs = append(qs, q)
		if !contains(as, a) {
			as = append(as, a)
		}
	}
	rand.Shuffle(len(qs), func(i, j int) { qs[i], qs[j] = qs[j], qs[i] })
	return qs, as
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
