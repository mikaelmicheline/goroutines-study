package app

import (
	"fmt"
	"strconv"
	"time"

	"github.com/urfave/cli"
)

func SetupWorkerPoolsPatternExampleCommand(app *cli.App) {
	const numberOfFactorials = 20
	const numberOfWorkers = 4

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

	app.Commands = append(app.Commands, cli.Command{
		Name:  "worker-pools-pattern",
		Usage: "Implements the Worker Pools Pattern",
		Action: func(c *cli.Context) {
			fmt.Println("Worker Pools Pattern example!")
			tasks, results := make(chan uint, numberOfFactorials), make(chan string, numberOfFactorials)

			for i := range uint(numberOfFactorials) {
				tasks <- i
			}
			close(tasks)

			for i := range numberOfWorkers {
				go worker("Worker "+strconv.Itoa(i+1), tasks, results)
			}

			for range uint(numberOfFactorials) {
				fmt.Println(<-results)
			}
			close(results)
		},
	},
	)
}
