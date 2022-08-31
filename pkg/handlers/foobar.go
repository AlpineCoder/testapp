package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"k8s.io/klog/v2"
)

func FooBar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from testapp-backend %d", rand.Intn(65000))
	if strings.Contains(r.URL.Path, "error") {
		klog.Error("This in an error")
	} else {
		klog.Infof("Path %s has been requested", r.URL.Path)
	}
}
