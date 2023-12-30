package restapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/koksmat-com/koksmat/model"
)

func executePowerShell(w http.ResponseWriter, r *http.Request) {

	host := r.URL.Query().Get("host")
	if (host != "exchange") && (host != "pnp") {
		w.WriteHeader(500)
		fmt.Fprint(w, "Invalid host")
		return
	}

	b, err := io.ReadAll(r.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		fmt.Fprint(w, "Error reading body")
		return
	}

	script := string(b)
	if !strings.Contains(strings.ToLower(script), "$result") {
		w.WriteHeader(500)
		log.Println(err)
		fmt.Fprint(w, `
		
		Missing a reference to a return value - Please set $result
		
		e.g.
		
		$result = Get-Date

		`)
		return
	}
	script = fmt.Sprintf(`%s
	ConvertTo-Json -InputObject $result
	| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM `, script)

	result, err := model.ExecutePowerShellScript("koksmat", host, script, "")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
	}
	w.WriteHeader(200)
	fmt.Fprint(w, result)

}
