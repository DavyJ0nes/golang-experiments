package reservation

type Reservation interface {
	IsRefundable() bool
	ReservationDate(string)
}

type Flight struct {
	Date  string
	offer bool
}

func (f *Flight) IsRefundable() bool {
	if f.offer {
		return false
	}

	return true
}

func (f *Flight) ReservationDate(rDate string) {
	f.Date = rDate
}

type Hotel struct {
	Date  string
	offer bool
}

func (h *Hotel) IsRefundable() bool {
	if h.offer {
		return false
	}

	return true
}

func (h *Hotel) ReservationDate(rDate string) {
	h.Date = rDate
}
