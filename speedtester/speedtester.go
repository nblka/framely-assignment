package speedtester

import "fmt"

type Service int

const (
	// SpeedTest.Net service
	SpeedTest Service = iota
	// Netflix fast.com service
	NetflixFastCom
)

type iService interface {
	// Runs speed testing and return measurement results. If data is not available, e.g. upload speed, it is set to 0.
	Run() (Measure, error)
}

// factory method for services
func createService(s Service) (iService, error) {
	switch s {
	case SpeedTest:
		return SpeedTestService{}, nil
	case NetflixFastCom:
		return FastComService{}, nil
	}
	return nil, fmt.Errorf("unknown service")
}

// returns download and upload speed along with error
func Test(s Service) (Measure, error) {
	service, err := createService(s)
	if err != nil {
		return Measure{0.0, 0.0}, err
	}
	return service.Run()
}
