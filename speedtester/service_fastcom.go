package speedtester

import (
	"fmt"

	"github.com/ddo/go-fast"
	"github.com/montanaflynn/stats"
)

type FastComService struct{}

// iService interface implementation
func (s FastComService) Run() (Measure, error) {
	fmt.Println("fast.com service run")

	fastCom := fast.New()

	err := fastCom.Init()
	if err != nil {
		return Measure{}, err
	}

	urls, err := fastCom.GetUrls()
	if err != nil {
		return Measure{}, err
	}
	fmt.Println("Got urls:", urls)

	KbpsChan := make(chan float64)
	var measurements []float64

	fmt.Println("Starting speed test...")
	go func() {
		for Kbps := range KbpsChan {
			fmt.Printf("Got %.2f Mbps\n", Kbps/1000)
			measurements = append(measurements, Kbps/1000)
		}

		fmt.Println("done")
	}()

	err = fastCom.Measure(urls, KbpsChan)
	if err != nil {
		return Measure{}, err
	}

	// find a median of measurements
	median, _ := stats.Median(measurements)
	if err != nil {
		return Measure{}, err
	}

	// fast.com don't supply upload speed, so set it to zero
	return Measure{Download: median, Upload: 0.0}, nil
}
