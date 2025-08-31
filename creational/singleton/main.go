package main

import (
	"fmt"
	"log/slog"

	"github.com/Wmuga/go-patterns/models/logger"
)

func main() {
	lg1 := logger.New()
	lg2 := logger.New()

	fmt.Println(lg1 == lg2)
	lg1.Info("Hello world", slog.Bool("same", lg1 == lg2))
}
