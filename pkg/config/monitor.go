package config

import (
	"fmt"
	"log"
)

//configuration for ordering
//allows limits and definitions depending on workload
type MonitorConfig struct {
	InfoLog         *log.Logger
	Date            string
	UseCache        bool
	SetOrderLimit   int
	FirstClassLimit bool
	NextDayLimit    bool
	ThreeDayLimit   bool
}

func main() {
	fmt.Println("monitor configuration")
}
