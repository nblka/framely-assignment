package main

import (
	"fmt"

	"github.com/nblka/framely-assignment/speedtester"
)

func main() {
	fmt.Println(speedtester.Test(speedtester.SpeedTest))
	fmt.Println(speedtester.Test(speedtester.NetflixFastCom))
}
