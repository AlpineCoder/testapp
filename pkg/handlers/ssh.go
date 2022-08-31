package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sfreiberg/simplessh"
	"k8s.io/klog/v2"
)

func SHHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["hostname"]
	fmt.Fprint(w, htmlHead)

	if !ok || len(keys[0]) < 1 {
		klog.Error("Url Param 'hostname' is missing")
		fmt.Fprint(w, "SSH without host! D'oh!")
		return
	}

	hostname := keys[0]

	fmt.Fprintf(w, "SSH to %s requested<p>", hostname)
	if strings.Contains(r.URL.Path, "error") {
		klog.Error("This in an error")
	} else {
		klog.V(3).Infof("Path %s has been requested", r.URL.Path)
	}

	_, err := sshAndRunCommand(hostname, "dummy", "pass")
	if err != nil {
		fmt.Fprintf(w, "SSH failed with %s", err)
	}
	fmt.Fprint(w, "<p><a href=\"../\">Back</a>")
	fmt.Fprint(w, htmlFooter)

}

func sshAndRunCommand(hostname, username, command string) ([]byte, error) {
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
