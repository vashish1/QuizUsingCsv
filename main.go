package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	csvfilename := flag.String("csv", "problems.csv", "the csv file in the form of ")
	flag.Parse()

	file, err := os.Open(*csvfilename)
	if err != nil {
		log.Fatal("cannot open the csv file")
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal("failed to parse the csv file")
	}
	total := len(lines)
	problems := parselines(lines)
	var correct int

	for i, p := range problems {
		fmt.Printf("problem no: %d :%s=\n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == (strings.TrimSpace(p.a)) {
			fmt.Println("correct answer :-)")
			correct++
		}

	}
	fmt.Printf("you scored %d out of %d", correct, total)

}

type problem struct {
	q string
	a string
}

func parselines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret

}
