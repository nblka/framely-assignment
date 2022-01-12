package speedtester

import "fmt"

type Service int

const (
	SpeedTest Service = iota
	NetflixFastCom
)

type iService interface {
	Run() (Measure, error)
}

func createService(s Service) (iService, error) {
	switch s {
	case SpeedTest:
		return SpeedTestService{}, nil
		// TODO
	}
	return nil, fmt.Errorf("unknown service")
}

// returns download and upload speed
func Test(s Service) (Measure, error) {
	service, err := createService(s)
	if err != nil {
		return Measure{0.0, 0.0}, err
	}
	return service.Run()
}
