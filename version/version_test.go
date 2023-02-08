package version

import (
	"fmt"
	"testing"
)

func TestVersionCompare(t *testing.T) {
	v1 := NewVersion("1.4.3")
	v2 := NewVersion("1.4.9")
	v3 := NewVersion("1.14.13")

	fmt.Println("Version comparison:")
	fmt.Printf("%v %d %v\n", v1, v1.Compare(v2), v2)
	fmt.Printf("%v %d %v\n", v2, v2.Compare(v3), v3)
	fmt.Printf("%v %d %v\n", v2, v2.Compare(v1), v1)
}
