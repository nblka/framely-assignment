package speedtester

import (
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
)

type SpeedTestService struct{}

// iService interface implementation
func (s SpeedTestService) Run() (Measure, error) {
	// TODO: pass debug info to callback function
	fmt.Println("speedtest service run")

	user, err := speedtest.FetchUserInfo()
	if err != nil {
		return Measure{}, err
	}
	fmt.Println("User", user)

	serverList, err := speedtest.FetchServerList(user)
	if err != nil {
		return Measure{}, err
	}
	fmt.Println("Servers", serverList)

	targets, err := serverList.FindServer([]int{})
	if err != nil {
		return Measure{}, err
	}
	fmt.Println("Targets", targets)

	fmt.Print("Starting download speed test...")
	err = targets[0].DownloadTest(false)
	if err != nil {
		fmt.Print(" failed\n")
		return Measure{}, err
	}
	fmt.Print(" done\n")

	fmt.Print("Starting upload speed test...")
	err = targets[0].UploadTest(false)
	if err != nil {
		fmt.Print(" failed\n")
		return Measure{}, err
	}
	fmt.Print(" done\n")

	return Measure{Download: targets[0].DLSpeed, Upload: targets[0].ULSpeed}, nil
}
