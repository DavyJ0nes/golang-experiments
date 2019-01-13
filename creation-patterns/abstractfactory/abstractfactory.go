package abstractfactory

import "github.com/davyj0nes/golang-experiments/creation-patterns/reservation"

// We have Reservation and Invoice as two generic products
type Invoice interface{}

type AbstractFactory interface {
	CreateReservation() reservation.Reservation
	CreateInvoice() Invoice
}

type HotelFactory struct{}

func (f HotelFactory) CreateReservation() reservation.Reservation {
	return &reservation.Hotel{}
}

func (f HotelFactory) CreateInvoice() Invoice {
	return new(HotelInvoice)
}

type FlightFactory struct{}

func (f FlightFactory) CreateReservation() reservation.Reservation {
	return &reservation.Flight{}
}

func (f FlightFactory) CreateInvoice() Invoice {
	return new(FlightInvoice)
}

type HotelInvoice struct{}
type FlightInvoice struct{}

func GetFactory(vertical string) AbstractFactory {
	var factory AbstractFactory
	switch vertical {
	case "flight":
		factory = FlightFactory{}
	case "hotel":
		factory = HotelFactory{}
	}

	return factory
}
