package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

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
	linkFile, err := os.Open("./links.csv")
	check(err)
	r := csv.NewReader(linkFile)
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
	err = linkFile.Close()
	check(err)
	
	t := time.NewTicker(time.Minute)
	for {
		for _, link := range linkList {
			res, err := traceroute.Traceroute(link.Dest, &link.Options)
			check(err)
			saveResults(link, res)
		}
		<- t.C
	}

}

func saveResults(link Link, res traceroute.TracerouteResult) {
	fName := fmt.Sprintf("res-%s.csv", link.Dest)
	f, err := os.OpenFile(fName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	
	var line [6]string
	line[0] = strconv.FormatInt(time.Now().Unix(), 10)
	line[1] = fmt.Sprintf("%d.%d.%d.%d", res.Hops[0].Address[0], res.Hops[0].Address[1], res.Hops[0].Address[2], res.Hops[0].Address[3])
	line[2] = fmt.Sprintf("%d.%d.%d.%d", res.Hops[1].Address[0], res.Hops[1].Address[1], res.Hops[1].Address[2], res.Hops[1].Address[3])
	if res.Hops[0].Success{
		line[3] = strconv.FormatFloat(float64(res.Hops[0].ElapsedTime.Nanoseconds())/1000000, 'f', 4, 64)
	} else {
		line[3] = "-1"
	}
	if res.Hops[1].Success{
		line[4] = strconv.FormatFloat(float64(res.Hops[1].ElapsedTime.Nanoseconds())/1000000, 'f', 4, 64)
	} else {
		line[4] = "-1"
	}
	if res.Hops[0].Success && res.Hops[1].Success {
		line[5] = strconv.FormatFloat(
			(float64(res.Hops[1].ElapsedTime.Nanoseconds())/1000000) - 
			(float64(res.Hops[0].ElapsedTime.Nanoseconds())/1000000),
			'f', 4, 64)
	} else {
		line[5] = "-1"
	}

	w := csv.NewWriter(f)
	w.Comma = ';'
	err = w.Write(line[:])
	check(err)
	w.Flush()
	check(w.Error())

	err = f.Close()
	check(err)
}
