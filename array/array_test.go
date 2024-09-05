package array

import (
	"cmp"
	"reflect"
	"strings"
	"testing"

	A "github.com/IBM/fp-go/array"
	F "github.com/IBM/fp-go/function"
	O "github.com/IBM/fp-go/ord"
)

func TestPipeSlice(t *testing.T) {
	got := F.Pipe3(
		A.From("Daniel", "Kate", "John", "Tom", "Donald", "Nancy", "Patrick"),
		A.Filter(func(s string) bool {
			return len(s) > 3
		}),
		A.Map(func(s string) string {
			return strings.ToUpper(s)
		}),
		A.Sort(O.FromCompare(func(a, b string) int {
			return cmp.Or(
				cmp.Compare(len(a), len(b)),
				strings.Compare(a, b),
			)
		})),
	)
	want := []string{"JOHN", "KATE", "NANCY", "DANIEL", "DONALD", "PATRICK"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}
