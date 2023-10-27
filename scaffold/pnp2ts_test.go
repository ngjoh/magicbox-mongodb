package scaffold

import (
	"fmt"
	"testing"
)

func TestPnp2Ts(t *testing.T) {
	//sharepointMap :=
	Pnp2Ts("/Users/nielsgregersjohansen/code/koksmat/ui/apps/www/app/cava/[site]/sharepoint/lists", "template.xml", "cava")
	Pnp2Ts("/Users/nielsgregersjohansen/code/koksmat/ui/apps/www/app/devops/[site]/sharepoint/lists", "devops.xml", "devops")
	fmt.Println(1)
	// fmt.Println(sharepointMap)
	//clipboard.WriteAll(sharepointMap)

}
