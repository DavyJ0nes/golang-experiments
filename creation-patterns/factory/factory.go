package factory

import "github.com/davyj0nes/golang-experiments/creation-patterns/reservation"

type Reservation interface {
	IsRefundable() bool
}

func NewReservation(vertical, reservationDate string) Reservation {
	switch vertical {
	case "flight":
		return &reservation.Flight{
			Date: reservationDate,
		}
	case "hotel":
		return &reservation.Hotel{
			Date: reservationDate,
		}
	default:
		return nil
	}
}
