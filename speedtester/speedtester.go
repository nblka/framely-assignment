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

// returns download and upload speed
func Test(s Service) Measure {
	return Measure{0.0, 0.0}
}