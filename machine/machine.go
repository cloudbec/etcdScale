package machine

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

// Machine represent etcd machine
type Machine struct {
	Name        string
	Peerip      net.IP
	Clusterip   net.IP
	Port        int
	ClusterPort int
	ClusterURL  string
}

// An IPNet represents an IP network.
type IPNet struct {
	IP   net.IP     // network number
	Mask net.IPMask // network mask
}

// Reverse my comment
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func (currentmachine Machine) Addetcdmember() string {

	//	url := fmt.Sprint("http://", currentmachine.Clusterip, ":4002/v2/members")
	url := fmt.Sprint("http://", currentmachine.Clusterip, ":", currentmachine.ClusterPort, "/v2/members")
	var jsonStr = []byte(fmt.Sprint(`{"peerURLs":["http://`, currentmachine.Peerip.String(), ":", currentmachine.Port, `"]}`))

	println(fmt.Sprint(`{"peerURLs":["http://`, currentmachine.Peerip.String(), ":", currentmachine.Port, `]}`))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

// Print print hello
func Print() {
	fmt.Print("hello")
}

// func (machine) printMachineStatus(stdOut []byte, stdErr error) {

// 	if strings.Contains(string(stdOut), "Added member named "+machine.name) {
// 		printOutput(
// 			stdOut,
// 		)

// 	} else {

// 		printError(stdErr)

// 	}
// }
