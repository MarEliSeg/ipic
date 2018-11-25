package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/aeden/traceroute"
)

type Link struct {
	Dest    string
	Options traceroute.TracerouteOptions
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var linkList []Link
	f, err := os.Open("./links.csv")
	check(err)
	r := csv.NewReader(f)
	r.Comma = ';'
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)

		firstHop, err := strconv.Atoi(record[1])
		check(err)
		maxHop, err := strconv.Atoi(record[2])
		check(err)

		link := Link{}
		link.Dest = record[0]
		link.Options = traceroute.TracerouteOptions{}
		link.Options.SetRetries(1)
		link.Options.SetFirstHop(firstHop)
		link.Options.SetMaxHops(maxHop)

		linkList = append(linkList, link)
	}
	fmt.Println(linkList)

}
