package main

import (
	"hatsumi/sqltemplate/people"
	"log"
)

func main() {
	fileName := "gomi.csv"
	data := people.ReadCsvFile(fileName)
	err := people.WriteTemplateFile("gomiout.sql", data)
	if err != nil {
		panic(err)
	}
	log.Println("end.")
}
