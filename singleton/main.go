package main

import (
	"fmt"

	"github.com/davyj0nes/golang-experiments/singleton/cnfg"
)

func main() {
	cnfg.MustLoad("config.yml")

	cpu := cnfg.Get("cpuMultiplier")

	fmt.Println("cpuMultiplier:", cpu)
}
