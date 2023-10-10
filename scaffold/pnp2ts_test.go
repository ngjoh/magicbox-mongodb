package scaffold

import (
	"fmt"
	"testing"

	"github.com/atotto/clipboard"
)

func TestPnp2Ts(t *testing.T) {
	sharepointMap := Pnp2Ts("template.xml")
	fmt.Println(1)
	// fmt.Println(sharepointMap)
	clipboard.WriteAll(sharepointMap)

}

func TestPnp2Go(t *testing.T) {
	sharepointMap := Pnp2Go("template.xml")
	// fmt.Println(sharepointMap)
	clipboard.WriteAll(sharepointMap)

}
