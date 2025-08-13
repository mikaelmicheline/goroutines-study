package app

import (
	"fmt"
	"time"

	"github.com/urfave/cli"
)

const numberOfFactorials = 20

func SetupWorkerPoolsPatternExampleCommand(app *cli.App) {
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

			go worker("Worker 1", tasks, results)
			go worker("Worker 2", tasks, results)
			go worker("Worker 3", tasks, results)
			go worker("Worker 4", tasks, results)

			for range uint(numberOfFactorials) {
				fmt.Println(<-results)
			}

			close(results)
		},
	},
	)
}

func worker(name string, tasks <-chan uint, results chan<- string) {
	for value := range tasks {
		time.Sleep(time.Millisecond * 300)
		result := factorial(value)
		results <- fmt.Sprintf("%s: %d! = %d", name, value, result)
	}
}

func factorial(value uint) uint {
	if value == 0 {
		return 1
	}
	return value * factorial(value-1)
}
