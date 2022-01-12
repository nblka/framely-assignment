package speedtester

import "fmt"

type Service int

const (
	SpeedTest Service = iota
	NetflixFastCom
)

type Measure struct {
	Download float32
	Upload   float32
}

func (m Measure) String() string {
	return fmt.Sprintf("Download %v Mbps, upload %v Mbps", m.Download, m.Upload)
}

type iService interface {
	Run() Measure
}

func createService(s Service) (iService, error) {
	// TODO
	return nil, fmt.Errorf("unknown service")
}

// returns download and upload speed
func Test(s Service) (Measure, error) {
	service, err := createService(s)
	if err != nil {
		return Measure{0.0, 0.0}, err
	}
	return service.Run(), nil
}
