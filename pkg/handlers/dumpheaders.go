package handlers

import (
	"fmt"
	"net/http"
	"sort"
)

func DumpHeaders(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlHead)

	keys := make([]string, 0, len(r.Header))
	for k := range r.Header {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Fprintf(w, "%s : %s<br>", k, r.Header[k])
	}
	fmt.Fprint(w, htmlFooter)
}
