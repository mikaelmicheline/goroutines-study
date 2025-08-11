# github.com/mikaelmicheline/goroutines-study
Study of Goroutines

`go run . wait-group`
- Uses a WaitGroup to wait for Goroutines to finish

`go run . channel`
- Uses a Channel to send messages from an additional Goroutine to the main Goroutine

`go run . channel-with-buffer`
- Uses a Channel with Buffer to pause the sender when the Buffer's capacity is reached and to release the sender when the receiver frees space