package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Driver struct {
	Name  string
	Score int
}

func newDriver(fields []string) (d Driver, err error) {
	if len(fields) < 2 {
		return d, fmt.Errorf("not enough data for Driver")
	}
	d.Name = fields[0]
	d.Score, err = strconv.Atoi(fields[1])
	if err != nil {
		return d, fmt.Errorf("failed to convert score into int")
	}
	return d, nil
}

type byScore []Driver

func (a byScore) Len() int           { return len(a) }
func (a byScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byScore) Less(i, j int) bool { return a[j].Score < a[i].Score }

func main() {
	drivers, err := readData("input.csv")

	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(byScore(drivers))

	output, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	for _, driver := range drivers {
		line := ""
		charsLen := utf8.RuneCountInString(driver.Name) + countLengthOfInt(driver.Score)
		if charsLen < 21 {
			line = fmt.Sprintln(driver.Name, strings.Repeat(" ", 21-charsLen), driver.Score)
		} else {
			line = fmt.Sprintln(driver.Name, driver.Score)
		}
		output.WriteString(line)
	}
}

func countLengthOfInt(number int) int {
	str := strconv.Itoa(number)
	return utf8.RuneCountInString(str)
}

func readData(fileName string) ([]Driver, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)

	defer f.Close()

	var drivers []Driver

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		driver, err := newDriver(record)
		if err != nil {
			continue
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}
