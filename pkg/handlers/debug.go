package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"k8s.io/klog/v2"
)

func Debug(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlHead)
	fmt.Fprintf(w, "Hello, you've requested: %s<p>", r.URL.Path)
	fmt.Fprintf(w, "Your lucky number is: %d", rand.Intn(20000))
	fmt.Fprint(w, "<form action=\"/ssh\"><label for=\"hostname\">SSH to:</label><input type=\"text\" id=\"hostname\" name=\"hostname\"><br><br>")
	fmt.Fprintf(w, "<input type=\"submit\" value=\"Submit\"></form>")
	if strings.Contains(r.URL.Path, "error") {
		klog.Error("This in an error")
	} else {
		klog.V(3).Infof("Path %s has been requested", r.URL.Path)
	}

	fmt.Fprintf(w, "<a href=\"/backend\">Backend</a>")
	fmt.Fprintf(w, "<p>")
	fmt.Fprintf(w, "<a href=\"/headers\">Headers</a>")

	fmt.Fprint(w, htmlFooter)
}
