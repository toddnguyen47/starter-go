package main

import (
	"fmt"

	"github.com/toddnguyen47/starter-projects/go/starter-go/pkg/prettylog"
)

func main() {

	logger := prettylog.GetConsoleLogger()
	// Looks like we have to end with `Msg` for the log message to appear
	logger.Info().Str("hello", "world").Msg("hi")
	fmt.Println("hello world!")
}
