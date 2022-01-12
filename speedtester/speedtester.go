package speedtester

type Service int

const (
	SpeedTest Service = iota
	NetflixFastCom
)

// returns download speed
func Test(s Service) float32 {
	return 0.0
}
