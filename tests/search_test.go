package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	library "github.com/Rayato159/tdd-practice"
)

type testSearch struct {
	id     int
	expect *library.Book
}

func ToJsonStringtify(obj any) string {
	data, _ := json.MarshalIndent(obj, "", "\t")
	return string(data)
}

func TestSearch(t *testing.T) {
	testObjs := []testSearch{
		{
			id: 2,
			expect: &library.Book{
				Id:    2,
				Title: "Physics",
			},
		},
	}

	for _, obj := range testObjs {
		// Search
		result := library.NewShelf().Search(obj.id)

		// Compare result
		if ToJsonStringtify(&result) != ToJsonStringtify(&obj.expect) {
			t.Errorf(
				"\nexpect:\n%v\ngot:\n%v\n",
				ToJsonStringtify(&obj.expect),
				ToJsonStringtify(&result),
			)
		}

		// Output the pass result
		fmt.Printf(
			"\nexpect:\n%v\ngot:\n%v\n",
			ToJsonStringtify(&obj.expect),
			ToJsonStringtify(&result),
		)
	}
}
