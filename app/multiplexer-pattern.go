package app

import (
	"fmt"
	"time"

	"github.com/urfave/cli"
)

func SetupMultiplexerPatternExampleCommand(app *cli.App) {
	type NumberType int

	const (
		Even NumberType = iota
		Odd
	)

	numberTypeToString := func(n NumberType) string {
		return [...]string{"Even", "Odd"}[n]
	}

	generateNumbers := func(amount int, numberType NumberType) <-chan string {
		channel := make(chan string)

		go func() {
			value := int(numberType)
			for range amount {
				time.Sleep(time.Millisecond * 300)
				channel <- fmt.Sprintf("%s: %d", numberTypeToString(numberType), value)
				value += 2
			}
			close(channel)
		}()

		return channel
	}

	multiplexer := func(channel1, channel2 <-chan string) <-chan string {
		channel := make(chan string)
		isChannel1Open := true
		isChannel2Open := true

		go func() {
		ForLoop:
			for {
				select {
				case msg, open := <-channel1:
					if open {
						channel <- msg
						continue
					}
					isChannel1Open = false
					if !isChannel2Open {
						break ForLoop
					}
				case msg, open := <-channel2:
					if open {
						channel <- msg
						continue
					}
					isChannel2Open = false
					if !isChannel1Open {
						break ForLoop
					}
				}
			}
			close(channel)
		}()

		return channel
	}

	app.Commands = append(app.Commands, cli.Command{
		Name:  "multiplexer-pattern",
		Usage: "Implements the Multiplexer Pattern",
		Action: func(c *cli.Context) {
			fmt.Println("Multiplexer Pattern example!")
			channel := multiplexer(generateNumbers(10, Even), generateNumbers(10, Odd))
			for msg := range channel {
				fmt.Println(msg)
			}
		},
	},
	)
}
