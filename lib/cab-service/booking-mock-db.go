package cs

type mockBookingsDb struct {
	brs []BookingRequest
}

func newBookingsMockDb() *mockBookingsDb {
	s := []BookingRequest{}
	return &mockBookingsDb{
		brs: s,
	}
}

func (repo *mockBookingsDb) SaveBooking(br BookingRequest) (id int, err error) {
	repo.brs = append(repo.brs, br)
	return len(repo.brs), nil
}

func (repo *mockBookingsDb) FetchBookings(userID int) ([]BookingRequest, error) {
	filtered := []BookingRequest{}
	for i, v := range repo.brs {
		br := v
		br.ID = i
		if br.UserID == userID {
			filtered = append(filtered, br)
		}
	}
	return filtered, nil
}

func (repo *mockBookingsDb) FetchUnconfirmedBookings(userID int) ([]BookingRequest, error) {
	filtered := []BookingRequest{}
	for i, v := range repo.brs {
		br := v
		br.ID = i
		if br.IsConfirmed == 0 && br.UserID == userID {
			filtered = append(filtered, br)
		}
	}
	return filtered, nil
}
