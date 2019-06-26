package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvfilename := flag.String("csv", "problems.csv", "the csv file in the form of ")
	timelimit := flag.Int("limit", 10, "the time limit is in seconds")
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
	timer := time.NewTimer((time.Duration(*timelimit)) * time.Second)

	for i, p := range problems {
		fmt.Printf("problem no: %d :%s=\n", i+1, p.q)
		answerch := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerch <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("you scored %d out of %d", correct, total)
			return
		case answer := <-answerch:
			if answer == (strings.TrimSpace(p.a)) {
				fmt.Println("correct answer :-)")
				correct++
			}
		}

	}

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
