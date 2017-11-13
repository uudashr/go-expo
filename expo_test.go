package expo_test

import (
	"fmt"
	"reflect"
	"testing"

	expo "github.com/uudashr/go-expo"
)

func ExampleExpandOptions() {
	opts := expo.ExpandOptions{"author", "publisher.contact"}

	expand := opts.Expand("author")
	fmt.Println("Expand author:", expand)

	expand = opts.Expand("category")
	fmt.Println("Expand category:", expand)

	expand = opts.Expand("publisher")
	fmt.Println("Expand publisher:", expand)

	pubOpts := opts.Sub("publisher")
	expand = pubOpts.Expand("contact")
	fmt.Println("Expand publisher.contact:", expand)

	// Output:
	// Expand author: true
	// Expand category: false
	// Expand publisher: true
	// Expand publisher.contact: true
}

func TestNoExpand(t *testing.T) {
	opts := expo.ExpandOptions{"author", "publisher.contact"}
	authorOpts := opts.Sub("publisher")
	if got, want := authorOpts, expo.NoExpand; !reflect.DeepEqual(authorOpts, expo.NoExpand) {
		t.Error("got:", got, "want:", want)
	}
}
