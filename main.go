package main

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"encoding/csv"
	"log"
	"flag"
	"os"
	"io"
	"io/ioutil"
//	"fmt"
)

func main() {
	flag.Parse()

	file, err := os.Open(flag.Arg(0))
	failOnError(err)
	defer file.Close()

	reader := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else {
			failOnError(err)
		}

		log.Printf("%#v", record)
		content := []byte(record[0] + "\n" + record[1] + "\n")
		ioutil.WriteFile("out.txt", content, os.ModePerm)
	}	
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}
