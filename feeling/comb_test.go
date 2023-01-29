package feeling

import (
	"fmt"
	"testing"
)

func TestComb(t *testing.T) {
	fmt.Println(GetDL("ab"))
	fmt.Println(GetDL10("ab"))
	fmt.Println(IsXKP("fr"))
	fmt.Println(IsDKP("ce"))
	fmt.Println(IsCS("xe"))
	fmt.Println(IsXZGR("as"))
}
