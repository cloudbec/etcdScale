package machine

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

type machine struct {
	name      string
	peerip    net.IP
	clusterip net.IP
}

func (currentmachine machine) addetcdmember() string {

	url := fmt.Sprint("http://", currentmachine.clusterip, "/notes")
	var jsonStr = []byte(fmt.Sprint(`{"peerURLs":["http://`, currentmachine.peerip, `]}`))
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
