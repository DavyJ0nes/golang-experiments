package builder

import "github.com/davyj0nes/golang-experiments/creation-patterns/reservation"

func NewReservationBuilder() ReservationBuilder {
	return &reservationBuilder{}
}

type ReservationBuilder interface {
	Vertical(string) ReservationBuilder
	ReservationDate(string) ReservationBuilder
	Build() reservation.Reservation
}

type reservationBuilder struct {
	vertical string
	rdate    string
}

func (r *reservationBuilder) Vertical(v string) ReservationBuilder {
	r.vertical = v
	return r
}

func (r *reservationBuilder) ReservationDate(date string) ReservationBuilder {
	r.rdate = date
	return r
}

func (r *reservationBuilder) Build() reservation.Reservation {
	var builtReservation reservation.Reservation

	switch r.vertical {
	case "flight":
		builtReservation = &reservation.Flight{
			Date: r.rdate,
		}
	case "hotel":
		builtReservation = &reservation.Hotel{
			Date: r.rdate,
		}
	}

	return builtReservation
}
