package pretty_test

import (
	"fmt"
	"github.com/fritzkeyzer/go-utils/pretty"
	"testing"
)

type Example struct {
	Title    string
	Desc     string
	Integers []int
	Floats   []float64
}

func ExamplePrint() {
	type Object struct {
		Field        string
		privateField string
	}

	o := Object{
		Field:        "hello world",
		privateField: "hidden",
	}

	fmt.Println(pretty.Print(o))
	//Output:Object:
	//   Field = "hello world"
}

func TestPrint(t *testing.T) {
	o := Example{
		Title:    "title",
		Desc:     "desc",
		Integers: []int{0, 1, 2},
		Floats:   []float64{0, 1, 2},
	}

	want := `Example:
  Title = "title"
  Desc = "desc"
  Integers = [0, 1, 2]
  Floats = [0.0, 1.0, 2.0]
  `
	got := fmt.Sprint(pretty.Print(o))

	if got != want {
		t.Fatalf("got != want:\ngot: %v\nwant: %v", got, want)
	}

}
