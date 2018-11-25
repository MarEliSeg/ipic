package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aeden/traceroute"
)

func main() {
	var linkList []traceroute.TracerouteOptions
	f, err := os.Open("./links.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	r.Comma = ';'
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		firstHop, err := strconf.Atoi(record[1])
		maxHop, err := strconf.Atoi(record[2])
		link := traceroute.TracerouteOptions{}
		link.SetRetries(1)
		link.SetFirstHop(firstHop)
		link.SetMaxHops(maxHop)
		linkList = append(linkList, link)
	}
	fmt.Println(linkList)

}
