package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	src := rand.New(rand.NewSource(time.Now().UnixNano()))

	fileName := fmt.Sprintf("turnttable_bulk_csv%d.csv", time.Now().Unix())

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Error creating file: " + err.Error())
	}

	writer := csv.NewWriter(file)
	headers := []string{"title", "artist", "label", "cn", "genre", "released", "purchased", "medium"}

	if err := writer.Write(headers); err != nil {
		log.Fatalln("Error writing line: " + err.Error())
	}

	mediums := []string{"CD", "Vinyl", "Cassette"}

	for i := 0; i < 100; i++ {
		row := []string{
			fmt.Sprintf("This is an album title%d", i),
			fmt.Sprintf("Your mother%d", i),
			fmt.Sprintf("holedaemon Records + %d", i),
			fmt.Sprintf("HDR%d", i),
			"skapunk",
			"2023-02-05",
			"2023-02-05",
			mediums[src.Intn(len(mediums))],
		}

		if err := writer.Write(row); err != nil {
			log.Fatalln("Error writing line: " + err.Error())
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatalln("Error encountered: " + err.Error())
	}
}
