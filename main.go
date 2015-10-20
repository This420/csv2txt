package main

import (
	"encoding/csv"
	"flag"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"os"
	"bufio"
)

func main() {
	var writer *bufio.Writer

	write_file, _ := os.OpenFile("out.txt", os.O_WRONLY|os.O_CREATE, 0600)
	writer = bufio.NewWriter(write_file)

	flag.Parse()

	file, err := os.Open(flag.Arg(0))
	failOnError(err)
	defer file.Close()

	reader := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))

	var cnt = 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else {
			failOnError(err)
		}

		cnt++
		if cnt == 1 { continue }

		log.Printf("%#v", record)

		var contact string

		if record[1] == "" {
			contact = "一般"
		} else {
			contact = record[1]
		}

		content := []byte("----------------------------------------------------------------------------------------------------------\n" +
			"●問い合わせ元\n" +
			contact + "\n" +
			"●質問\n" +
			record[7] + "\n" +
			"★回答\n" +
			record[11] + "\n")
		writer.Write(content)
		writer.Flush()
	}
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}
