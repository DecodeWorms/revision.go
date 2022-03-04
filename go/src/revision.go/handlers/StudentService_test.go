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

	t.Run("Result of the fruit test", func(t *testing.T) {
		if got != expFruits {
			t.Errorf("expected %v got %v", expFruits, got)
		}

	})

}

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 5.0)
	expected := 30.0
	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {

	t.Run("testing for an area", func(t *testing.T) {
		rectange := Dimension{area: 22.0, wid: 55.0}
		got := Area(rectange)
		want := 154.0
		if got != want {
			t.Errorf("got %.2f and want %.2f", got, want)
		}
	})

}

func TestCircle(t *testing.T) {
	t.Run("testing circle ", func(t *testing.T) {
		d := Dimension{area: 25.0, wid: 100}
		got := d.Circle()
		want := 2500.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestGetResources(t *testing.T) {

	res := func(t *testing.T, serv Services, got []models.User) {

		want := []models.User{
			{Id: 1, Name: "Yinka", Gender: "female"},

			{Id: 2, Name: "Biola", Gender: "male"},

			{Id: 3, Name: "Kunle", Gender: "male"},
		}

		if len(got) != len(want) {
			t.Error("the length of got and want is not the same")
		}

		for index, value := range want {

			if value != got[index] {
				t.Errorf("got %v expected %v", got, want)
			}

		}

	}

	t.Run("testing resources", func(t *testing.T) {
		var j Junior

		got := j.GetResources()

		res(t, j, got)
	})
}

func TestCreateResources(t *testing.T) {

	res := func(t *testing.T, ser Services, got models.User) {

		want := models.User{Id: 1, Name: "Bola", Gender: "female"}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

	}

	t.Run("testing create resources", func(t *testing.T) {
		var j Junior

		user := models.User{Id: 2, Name: "Bolaa", Gender: "female"}

		got := j.CreateResource(user)

		res(t, j, got)

	})
}
