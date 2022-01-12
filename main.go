package main

import (
	"fmt"
	"main/speedtester"
)

func main() {
	fmt.Println(speedtester.Test(speedtester.SpeedTest))
	fmt.Println(speedtester.Test(speedtester.NetflixFastCom))
}
