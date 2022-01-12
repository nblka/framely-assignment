package speedtester

import "fmt"

type Measure struct {
	Download float64
	Upload   float64
}

func (m Measure) String() string {
	return fmt.Sprintf("Downloading speed %v Mbps, uploading speed %v Mbps", m.Download, m.Upload)
}
