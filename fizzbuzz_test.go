package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGivenNonFizzBuzzNumber(t *testing.T) {
	Convey("GivenNonFizzBuzzNumber", t, func() {
		Convey("Input = 1", func() {
			number := 1

			words := fizzBuzz(number)

			So(len(words), ShouldEqual, 1)
			So(words, ShouldContain, "1")
		})

		Convey("Input = 2", func() {
			number := 2

			words := fizzBuzz(number)

			So(len(words), ShouldEqual, 2)
			So(words, ShouldContain, "1")
			So(words, ShouldContain, "2")
		})
	})
}

func TestGivenFizzNumber(t *testing.T) {
	Convey("GivenFizzNumber", t, func() {
		Convey("Input = 3", func() {
			number := 3

			words := fizzBuzz(number)

			So(len(words), ShouldEqual, 3)
			So(words, ShouldContain, "1")
			So(words, ShouldContain, "2")
			So(words, ShouldContain, "Fizz")
		})
	})
}

func TestGivenFizzNumberAndBuzzNumber(t *testing.T) {
	Convey("GivenFizzNumberAndBuzzNumber", t, func() {
		Convey("Input = 5", func() {
			number := 5
			expecteds := []string{
				"1", "2", "Fizz", "4", "Buzz",
			}

			words := fizzBuzz(number)

			So(len(words), ShouldEqual, 5)
			for _, expected := range expecteds {
				So(words, ShouldContain, expected)
			}
		})
	})
}

func TestGivenFizzBuzz(t *testing.T) {
	Convey("GivenFizzBuzz", t, func() {
		Convey("Input = 15", func() {
			number := 15
			expecteds := []string{
				"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8",
				"Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz",
			}

			words := fizzBuzz(number)

			So(len(words), ShouldEqual, 15)
			for _, expected := range expecteds {
				So(words, ShouldContain, expected)
			}
		})
	})
}
