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

// Addetcmember add current machineconfig  to configured cluster in Machine struct
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

// func (machine) printMachineStatus(stdOut []byte, stdErr error) {

// 	if strings.Contains(string(stdOut), "Added member named "+machine.name) {
// 		printOutput(
// 			stdOut,
// 		)

// 	} else {

// 		printError(stdErr)

// 	}
// }
