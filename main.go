package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"k8s.io/klog/v2"

	"inventx.ch/testapp/pkg/geo"
)

func init() {
	klog.InitFlags(nil)
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")
	if os.Getenv("DEBUGLEVEL") == "" {
		flag.Set("v", "2")
	} else {
		fmt.Printf("Parsing %s\n", os.Getenv("DEBUGLEVEL"))
		i, err := strconv.Atoi(os.Getenv("DEBUGLEVEL"))
		if err != nil {
			klog.Errorf("Failed to parse Debuglevel %s", os.Getenv("DEBUGLEVEL"))
			klog.V(1).Info("Starting with debug level 2")
			flag.Set("v", "2")
		} else {
			if i > 4 {
				klog.V(1).Info("Starting with debug level 4")
				flag.Set("v", "4")
			} else {
				klog.V(1).Infof("Starting with debug level %d", i)
				flag.Set("v", os.Getenv("DEBUGLEVEL"))
			}
		}

	}
	flag.Parse()

}

func main() {
	defer klog.Flush()
	klog.V(1).Info("Staring app...")
	klog.Error("Some shit")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		if strings.Contains(r.URL.Path, "error") {
			klog.Error("This in an error")
		} else {
			klog.V(3).Infof("Path %s has been requested", r.URL.Path)
		}

		g := &geo.Geo{}
		g.New(5, 8)
		g.Multiply()
	})

	http.ListenAndServe(":10000", nil)
}
