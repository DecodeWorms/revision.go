package handleers

import (
	"fmt"
	"testing"

	"revision.go/models"
)

//These are an example of table driven test on struct and string

func TestValidateData(t *testing.T) {

	type errorTestCase struct {
		input models.User
	}

	c := []errorTestCase{{
		input: models.User{
			Id:     1,
			Name:   "funke",
			Gender: "male",
		},
	}, {
		input: models.User{
			Id:     2,
			Name:   "yinka",
			Gender: "female",
		},
	}, {
		input: models.User{
			Id:     33,
			Name:   "fatai",
			Gender: "",
		},
	}}

	for _, scenario := range c {

		err := ValidateData(scenario.input)
		if err != nil {
			t.Error(err)

		}

	}

}

func TestValidateParameter(t *testing.T) {
	type name struct {
		n string
	}

	type errorTestCase struct {
		input name
	}

	c := []errorTestCase{{
		input: name{
			n: "name",
		},
	}, {
		input: name{
			n: "",
		},
	}, {
		input: name{
			n: "joe",
		},
	}}

	for _, scenario := range c {

		err := ValidateParameter(scenario.input.n)
		if err != nil {
			t.Error(err)
		}
	}

}

//A unit test on a func
func TestHello(t *testing.T) {

	got := Hello("Bola")
	expected := "Bola"
	if expected == got {
		fmt.Println("They are equal")
		return
	}
	t.Errorf("got %s expected %s", got, expected)

}

func TestCalc(t *testing.T) {
	got := Calc(5)
	expected := 2

	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func TestFruitsPrice(t *testing.T) {
	fruits := [4]int{100, 300, 50, 1000}

	got := FruitsPrice(fruits)
	expFruits := 1450

	fmt.Println(got, expFruits)

	t.Run("Result of the fruit test", func(t *testing.T) {
		if got != expFruits {
			t.Errorf("expected %v got %v", expFruits, got)
		}

	})

}
