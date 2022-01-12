package speedtester

import (
	"fmt"
	"testing"
)

func TestCreateSpeedTestService(t *testing.T) {
	f := serviceFactory{}
	s, ok := f.createService(SpeedTest)
	if ok != nil {
		t.Error()
		return
	}

	_, ok2 := s.(SpeedTestService)
	if ok2 == false {
		t.Error()
	}
}

func TestCreateFastcomService(t *testing.T) {
	f := serviceFactory{}
	s, ok := f.createService(NetflixFastCom)
	if ok != nil {
		t.Error()
		return
	}

	_, ok2 := s.(FastComService)
	if ok2 == false {
		t.Error()
	}
}

type mockServiceFactory struct {
	fail        bool
	serviceFail bool
}

type mockService struct {
	fail bool
}

func (f mockServiceFactory) createService(s Service) (iService, error) {
	if f.fail {
		return nil, fmt.Errorf("fail")
	}
	return mockService{fail: f.serviceFail}, nil

}

func (s mockService) Run() (Measure, error) {
	if s.fail {
		return Measure{}, fmt.Errorf("fail")
	}
	return Measure{Download: 3.14, Upload: 2.71}, nil
}

func TestTestWithFailingFactory(t *testing.T) {
	f := mockServiceFactory{fail: true}
	if _, err := test(SpeedTest, f); err == nil {
		t.Error()
	}
}

func TestTestWithFailingServiceRun(t *testing.T) {
	f := mockServiceFactory{fail: false, serviceFail: true}
	m, err := test(SpeedTest, f)
	if err == nil {
		t.Error()
		return
	}

	if m.Download != 0.0 || m.Upload != 0.0 {
		t.Error()
	}
}

func TestTestWithOkServiceRun(t *testing.T) {
	f := mockServiceFactory{fail: false, serviceFail: false}
	m, err := test(SpeedTest, f)
	if err != nil {
		t.Error()
		return
	}

	if m.Download != 3.14 || m.Upload != 2.71 {
		t.Error()
	}
}

func BenchmarkSpeedTestServiceUsage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test(SpeedTest)
	}
}

func BenchmarkNetflixFastComServiceUsage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test(NetflixFastCom)
	}
}
