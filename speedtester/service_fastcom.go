package speedtester

import (
	"fmt"
	"sync"

	"github.com/ddo/go-fast"
	"github.com/montanaflynn/stats"
)

type FastComService struct{}

var wg sync.WaitGroup

func (s FastComService) Run() (Measure, error) {
	fmt.Println("fast.com service run")

	fastCom := fast.New()

	// init
	err := fastCom.Init()
	if err != nil {
		return Measure{}, err
	}

	// get urls
	urls, err := fastCom.GetUrls()
	if err != nil {
		return Measure{}, err
	}
	fmt.Println("Got urls:", urls)

	// measure
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

	err = fastCom.Measure(urls[:1], KbpsChan)
	if err != nil {
		return Measure{}, err
	}

	// find a median of measurements
	median, _ := stats.Median(measurements)

	// fast.com don't supply upload speed, so set it to zero
	return Measure{Download: median, Upload: 0.0}, nil
}
