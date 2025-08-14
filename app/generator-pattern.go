package app

import (
	"fmt"
	"strconv"
	"time"

	"github.com/urfave/cli"
)

func SetupGeneratorPatternExampleCommand(app *cli.App) {
	calculateFactorials := func(numberOfFactorials uint, numberOfWorkers uint) <-chan string {
		var factorial func(value uint) uint
		factorial = func(value uint) uint {
			if value <= 1 {
				return 1
			}
			return value * factorial(value-1)
		}

		worker := func(name string, tasks <-chan uint, results chan<- string) {
			for value := range tasks {
				time.Sleep(time.Millisecond * 300)
				result := factorial(value)
				results <- fmt.Sprintf("%s: %d! = %d", name, value, result)
			}
		}

		tasks, results := make(chan uint, numberOfFactorials), make(chan string)

		for i := range int(numberOfWorkers) {
			go worker("Worker "+strconv.Itoa(i+1), tasks, results)
		}

		for i := range numberOfFactorials {
			tasks <- i
		}

		close(tasks)
		return results
	}

	app.Commands = append(app.Commands, cli.Command{
		Name:  "generator-pattern",
		Usage: "Implements the Generator Pattern",
		Action: func(c *cli.Context) {
			fmt.Println("Generator Pattern example!")
			const numberOfFactorials uint = 20
			const numberOfWorkers uint = 4
			channel := calculateFactorials(numberOfFactorials, numberOfWorkers)

			for range numberOfFactorials {
				fmt.Println(<-channel)
			}
		},
	},
	)
}
