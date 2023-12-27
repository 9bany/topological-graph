package topologicalgraph

import (
	"fmt"
	"os"
	"testing"
)

func TestTo(t *testing.T) {
	g := NewGraph[string]()

	g.Add("a", "b")
	g.Add("a", "c")
	g.Add("b", "c")
	g.Add("c", "e")
	g.Add("b", "e")

	resutl := g.Sort("a")

	if resutl[0] != "e" || resutl[1] != "c" || resutl[2] != "b" || resutl[3] != "s" {
		fmt.Fprintf(os.Stderr, "unexpected")
	}

	resutl2 := g.Sort("b")

	if resutl2[0] != "e" || resutl2[1] != "c" || resutl2[2] != "b" {
		fmt.Fprintf(os.Stderr, "unexpected")
	}
}
