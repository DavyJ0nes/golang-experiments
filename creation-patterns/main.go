package main

import (
	"fmt"
	"github.com/davyj0nes/golang-experiments/creation-patterns/abstractfactory"
	"github.com/davyj0nes/golang-experiments/creation-patterns/builder"
	"github.com/davyj0nes/golang-experiments/creation-patterns/factory"
)

func main() {
	fmt.Println("constructing with Factory")
	constructWithFactory()
	fmt.Println("")

	fmt.Println("constructing with Builder")
	constructWithBuilder()
	fmt.Println("")

	fmt.Println("constructing with Abstract Factory")
	constructWithAbstractFactory()
}

func constructWithFactory() {
	res := factory.NewReservation("hotel", "20180101")
	fmt.Printf("Reservation: %+v\n", res)
	fmt.Println("Reservation is refundable:", res.IsRefundable())
}

func constructWithBuilder() {
	resBuilder := builder.NewReservationBuilder()
	resBuilder.Vertical("flight")
	resBuilder.ReservationDate("20190111")
	res := resBuilder.Build()

	fmt.Printf("Reservation: %+v\n", res)
	fmt.Println("Reservation is refundable:", res.IsRefundable())
}

func constructWithAbstractFactory() {
	hotelFactory := abstractfactory.GetFactory("hotel")

	res := hotelFactory.CreateReservation()
	res.ReservationDate("20190111")
	invoice := hotelFactory.CreateInvoice()

	fmt.Printf("Reservation: %+v\n", res)
	fmt.Printf("Invoice: %+v\n", invoice)
}
