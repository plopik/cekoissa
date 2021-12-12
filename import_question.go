package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

var extension = []string{"png", "jpg", "jpeg", "gif"}

func (s *serie) import_image(folder string, color string) {
	files, err := ioutil.ReadDir("data/" + folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		for _, ext := range extension {
			if strings.HasSuffix(name, "."+ext) || strings.HasSuffix(name, "."+strings.ToUpper(ext)) {
				s.qs = append(s.qs, folder+"/"+name)
			}
		}
	}
	rand.Shuffle(len(s.qs), func(i, j int) { s.qs[i], s.qs[j] = s.qs[j], s.qs[i] })

	for _, q := range s.qs {
		q2 := strings.ToLower(q)
		for _, ext := range extension {
			q2 = strings.Replace(q2, "."+ext, "", -1)
		}
		q2 = strings.Trim(q2, "2")
		q2 = strings.Trim(q2, "3")
		q2 = strings.Title(q2)
		la := strings.Split(q2, "/")
		a := strings.Replace(la[len(la)-1], "_", " ", -1)
		questionsMap[q] = question{nil, q, a, color}
		if !contains(s.as, a) {
			s.as = append(s.as, a)
		}
	}
}

func (s *serie) import_xlsx(file string) {

	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet 1")
	if err != nil {
		fmt.Println(err)
		return
	}
	sentences := rows[0]
	fmt.Println(sentences)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		a := row[0]
		q := [][]string{}
		for j, carac := range row {
			if j != 0 && carac != "" {
				carac = strings.ReplaceAll(carac, " | ", "\n")
				q = append(q, []string{sentences[j], carac})
			}
		}
		if len(q) > 1 {
			questionsMap[a] = question{words: q, response: a}
			s.qs = append(s.qs, a)
			if !contains(s.as, a) {
				s.as = append(s.as, a)
			}
		}
	}

	rand.Shuffle(len(s.qs), func(i, j int) { s.qs[i], s.qs[j] = s.qs[j], s.qs[i] })
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

		questionsMap[q] = question{words: [][]string{{sentence, q}}, response: a}
		s.qs = append(s.qs, q)
		if !contains(s.as, a) {
			s.as = append(s.as, a)
		}
	}
	rand.Shuffle(len(s.qs), func(i, j int) { s.qs[i], s.qs[j] = s.qs[j], s.qs[i] })
}

func (s *serie) import_csv2(file string) {
	records := readCsvFile(file)
	sentences := records[0]
	for i, line := range records {
		if i == 0 {
			continue
		}
		a := line[0]
		q := [][]string{}
		for j, carac := range line {
			if j != 0 && carac != "" {
				if len(carac) > 3 {
					carac = strings.Title(carac)
				}
				q = append(q, []string{sentences[j], carac})
			}
		}
		if len(q) > 1 {
			questionsMap[a] = question{words: q, response: a}
			s.qs = append(s.qs, a)
			if !contains(s.as, a) {
				s.as = append(s.as, a)
			}
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
