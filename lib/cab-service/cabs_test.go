package cs

import (
	"testing"
)

func TestGetCabs(t *testing.T) {
	newCabs := []Cab{
		Cab{
			VehNo: "123",
			Lat:   12.11,
			Lon:   66.12,
		},
		Cab{
			VehNo: "124",
			Lat:   12.11,
			Lon:   66.12,
		},
		Cab{
			VehNo: "125",
			Lat:   12.11,
			Lon:   66.12,
		},
	}

	db := NewCabsMockDb(newCabs)

	cs := NewCabsService(db)

	cabs, _ := cs.ListAll(12.11, 66.12, 10)

	if len(cabs) != 3 {
		t.FailNow()
	}
}
