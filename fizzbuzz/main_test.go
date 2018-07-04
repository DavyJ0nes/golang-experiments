package main

import (
	"reflect"
	"testing"
)

var want = []string{
	"1",
	"2",
	"Fizz",
	"4",
	"Buzz",
	"Fizz",
	"7",
	"8",
	"Fizz",
	"Buzz",
	"11",
	"Fizz",
	"13",
	"14",
	"FizzBuzz",
}

func TestBasicFizzBuzz(t *testing.T) {
	if got := basicFizzBuzz(15); !reflect.DeepEqual(got, want) {
		t.Errorf("basicFizzBuzz() = %v, want %v", got, want)
	}
}

func TestImprovedFizzBuzz(t *testing.T) {
	if got := improvedFizzBuzz(15); !reflect.DeepEqual(got, want) {
		t.Errorf("improvedFizzBuzz() = %v, want %v", got, want)
	}
}

func TestChanFizzBuzz(t *testing.T) {
	var got []string

	for out := range chanFizzBuzz(15) {
		got = append(got, out)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("chanFizzBuzz() = %v, want %v", got, want)
	}
}
