package main

import (
	"fmt"
	"time"
)

func main() {
	var startTime = time.Now()
	fmt.Printf("\n\n---- BrainGoFuck START %s ----\n", startTime.Format(time.TimeOnly))
	defer func() {
		fmt.Printf("---- BrainGoFuck END %s (%f) ----\n", startTime.Format(time.TimeOnly), time.Since(startTime).Seconds())
	}()
}
