package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var inputFile = flag.String("in", "input.csv", "The file path of csv input")
var outFile = flag.String("out", "output.json", "The file path of json output")
var delimeter = flag.String("del", ";", "The delimeter character between columns")
var named = flag.Bool("named", true, "Whether the first line contains column names")

func main() {
	flag.Parse()
	fmt.Printf("Converting %s to %s with the delimeter %s\n", *inputFile, *outFile, *delimeter)
	dat, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Cannot read csv file %s\n Error: %s", *inputFile, err.Error())
	}
	records := strings.Split(string(dat), "\n")
	line := strings.Split(records[0], *delimeter)
	colNames := make([]string, len(line))
	startIndex := 0
	if *named { // if the first line contains the names of columns
		for i, val := range line {
			val = strings.TrimSpace(val)
			if val == "" {
				colNames[i] = fmt.Sprintf("Column%d", i)
			} else {
				colNames[i] = val
			}
		}
		startIndex = 1 // skip the first row for values afterwards
	} else {
		for i := range records[0] {
			colNames[i] = fmt.Sprintf("Column%d", i)
		}
	}
	jsonString := fmt.Sprintf("[")
	for i := startIndex; i < len(records); i++ {
		//fmt.Printf("%v\n", records[i])
		itemString := ""
		line = strings.Split(records[i], *delimeter)
		for j, val := range line {
			itemString += fmt.Sprintf("\"%s\": \"%s\", ", colNames[j], strings.TrimSpace(val))
		}
		//fmt.Printf("%s\n", itemString)
		jsonString += fmt.Sprintf("\n{%s},", strings.TrimRight(itemString, ", "))
	}
	jsonString = fmt.Sprintf("%s\n]", strings.TrimRight(jsonString, ","))
	fmt.Printf("%s", jsonString)
	err = ioutil.WriteFile(*outFile, []byte(jsonString), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
