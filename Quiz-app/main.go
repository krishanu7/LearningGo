package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func problemPuller(fileName string) ([]problem, error) {
	// Read all the problems from the quiz.csv
	// 1. Open the file
	if fObj, err := os.Open(fileName); err == nil {
		defer fObj.Close()
		// 2. We'll use the csv package to create a reader
		csvReader := csv.NewReader(fObj)
		// 3. We'll read the file
		if lines, err := csvReader.ReadAll(); err == nil {
			// 4. We'll parse the problems
			return parseProblem(lines), nil
		} else {
			return nil, fmt.Errorf("Could not read the file: %s", err.Error())
		}
	}
	return nil, fmt.Errorf("Could not open the file: %s", fileName)
}

func parseProblem(lines [][]string) []problem {
	res := make([]problem, len(lines))

	for i := 0; i < len(lines); i++ {
		res[i] = problem{
			question: lines[i][0],
			answer:   lines[i][1],
		}
	}
	return res
}

func main() {
	//1. Input the name of the file
	fname := flag.String("f", "quiz.csv", "path of csv file")
	//2. Set the duration of the timer
	timer := flag.Int("t", 30, "time limit for the quiz")
	flag.Parse()
	//3. Pull the problems from the file (calling our problem puller func)
	problems, err := problemPuller(*fname)
	//4. Handle the error
	fmt.Println(problems)
	if err != nil {
		exit(fmt.Sprintf("Something went wrong: %s\n", err.Error()))
	}
	//5. Create a variable to count the number of correct answers
	correctAns := 0
	// 6. Create a channel to signal the end of the timer
	ansCh := make(chan string)
	//7. Using the duration of the timer, we want to initialize the timer
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	//8. Loop through the problems, print the questions, we will accept the answers
problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansCh <- ans
		}()

		select {
		case <-tObj.C:
			fmt.Println()
			break problemLoop
		case ans := <-ansCh:
			if ans == p.answer {
				correctAns++
			}
			if i == len(problems)-1 {
				tObj.Stop()
				close(ansCh)
			}
		}
	}
	fmt.Printf("You scored %d out of %d\n", correctAns, len(problems))
	fmt.Printf("Press enter to exit")
	fmt.Scanln()
}

// Try without labelled loop

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
