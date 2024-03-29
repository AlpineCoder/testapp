package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/AlpineCoder/testapp/pkg/handlers"
	"github.com/heptiolabs/healthcheck"
	"github.com/sfreiberg/simplessh"
	"k8s.io/klog/v2"
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
	port := ":" + os.Getenv("HTTP_PORT")
	klog.V(1).Info("Staring app...")
	health := healthcheck.NewHandler()
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))
	go http.ListenAndServe("0.0.0.0:8086", health)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, htmlHead)
	// 	fmt.Fprintf(w, "Hello, you've requested: %s<p>", r.URL.Path)
	// 	fmt.Fprintf(w, "Your lucky number is: %d", rand.Intn(20000))
	// 	fmt.Fprint(w, "<form action=\"/ssh\"><label for=\"hostname\">SSH to:</label><input type=\"text\" id=\"hostname\" name=\"hostname\"><br><br>")
	// 	fmt.Fprintf(w, "<input type=\"submit\" value=\"Submit\"></form>")
	// 	if strings.Contains(r.URL.Path, "error") {
	// 		klog.Error("This in an error")
	// 	} else {
	// 		klog.V(3).Infof("Path %s has been requested", r.URL.Path)
	// 	}

	// 	fmt.Fprintf(w, "<a href=\"/backend\">Backend</a>")
	// 	fmt.Fprintf(w, "<p>")
	// 	fmt.Fprintf(w, "<a href=\"/headers\">Headers</a>")

	// 	fmt.Fprint(w, htmlFooter)
	// })

	http.HandleFunc("/", handlers.Base)

	http.HandleFunc("/foobar", handlers.FooBar)
	http.HandleFunc("/ssh", handlers.SHHandler)
	http.HandleFunc("/debug", handlers.Debug)
	// http.HandleFunc("/backendCurl", getSomethingFromBackendCurl)
	// http.HandleFunc("/backend", getSomethingFromBackend)
	http.HandleFunc("/headers", handlers.DumpHeaders)
	http.HandleFunc("/image", handlers.ServeImage)
	// go doDNSLookup()
	http.ListenAndServe(port, nil)
}

// func getSomethingFromBackendCurl(w http.ResponseWriter, r *http.Request) {
// 	var sb string
// 	easy := curl.EasyInit()
// 	defer easy.Cleanup()

// 	easy.Setopt(curl.OPT_URL, "http://testapp-backend:10000/foobar")

// 	fooTest := func(buf []byte, userdata interface{}) bool {
// 		// println("DEBUG: size=>", len(buf))
// 		// println("DEBUG: content=>", string(buf))
// 		sb = string(buf)
// 		return true
// 	}

// 	easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)

// 	if err := easy.Perform(); err != nil {
// 		fmt.Printf("ERROR: %v\n", err)
// 		sb = string(err.Error())
// 	}

// 	klog.Infof("From Backend: %s", sb)
// 	fmt.Fprintf(w, "From Curl Backend: %s", sb)

// }

// func getSomethingFromBackend(w http.ResponseWriter, r *http.Request) {
// 	resp, err := http.Get("http://testapp-backend:10000/foobar")
// 	if err != nil {
// 		klog.Error(err)
// 	}
// 	//We Read the response body on the line below.
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		klog.Error(err)
// 	}
// 	//Convert the body to type string
// 	sb := string(body)
// 	fmt.Fprintf(w, "From Backend: %s", sb)
// }

func SshAndRunCommand(hostname, username, command string) ([]byte, error) {
	var client *simplessh.Client
	var err error

	// Option A: Using a specific private key path:
	// if client, err = simplessh.ConnectWithKeyFile(hostname, username, identityFile); err != nil {
	if client, err = simplessh.ConnectWithPasswordTimeout(hostname, username, "", 5*time.Second); err != nil {

		// Option B: Using your default private key at $HOME/.ssh/id_rsa:
		//if client, err = simplessh.ConnectWithKeyFile("hostname_to_ssh_to", "username"); err != nil {

		// Option C: Use the current user to ssh and the default private key file:
		//if client, err = simplessh.ConnectWithKeyFile("hostname_to_ssh_to"); err != nil {
		return make([]byte, 0), err
	}
	defer client.Close()

	// Now run the commands on the remote machine:
	if result, err := client.Exec(command); err != nil {
		klog.Error(err)
		return result, err
	} else {
		return result, err

	}
}

func doDNSLookup() {
	for {
		ips, err := net.LookupIP("testapp")
		if err != nil {
			klog.Errorf("DNS resolution failed %s", err)
		}
		for _, ip := range ips {
			klog.V(1).Infof("Service testapp has ip %s", ip)
		}
		time.Sleep(time.Second)
	}
}
