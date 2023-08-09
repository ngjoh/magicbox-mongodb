package restapi

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/koksmat-com/koksmat/model"
)

func executePowerShell(w http.ResponseWriter, r *http.Request) {

	host := r.URL.Query().Get("host")

	b, err := io.ReadAll(r.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		fmt.Fprint(w, "Error reading body")
		return
	}

	script := string(b)

	result, err := model.ExecutePowerShellScript("koksmat", host, script, "")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
	}
	w.WriteHeader(200)
	fmt.Fprint(w, result)

}
