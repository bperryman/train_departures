package train_departures

import "testing"

func TestStatusValues(tst *testing.T) {
	if OnTime != 1 {
		tst.Errorf("On time status should be 1")
	}
	if Delayed != 2 {
		tst.Errorf("Delayed status should be 2")
	}
	if Cancelled != 3 {
		tst.Errorf("Cancelled status should be 3")
	}
	if Late != 4 {
		tst.Errorf("Late status should be 4")
	}
	if Unknown != 5 {
		tst.Errorf("Unknown status should be 5")
	}
}
