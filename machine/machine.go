package machine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
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

type errorAddEtcdMember struct {
	s string
}

func (e *errorAddEtcdMember) Error() string {
	return e.s
}

type Member struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	PeerURLs   []string `json:"peerURLs"`
	ClientURLs []string `json:"clientURLs"`
}

// type Cluster struct {
// 	Id         int      `json:"id"`
// 	PeerURLs   []string `json:"peerURLs"`
// 	ClientURLs []string `json:"clientURLs"`
// }

type Message struct {
	Txt string `json:"message" mapstructure:"message"`
}

// Addetcmember add current machineconfig  to configured cluster in Machine struct
func (currentmachine Machine) Addetcdmember() (*Member, error) {

	// cluster := &Cluster{}

	//	url := fmt.Sprint("http://", currentmachine.Clusterip, ":4002/v2/members")
	url := fmt.Sprint("http://", currentmachine.Clusterip, ":", currentmachine.ClusterPort, "/v2/members")
	var jsonStr = []byte(fmt.Sprint(`{"peerURLs":["http://`, currentmachine.Peerip.String(), ":", currentmachine.Port, `"]}`))

	//	println(fmt.Sprint(`{"peerURLs":["http://`, currentmachine.Peerip.String(), ":", currentmachine.Port, `]}`))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	byt := []byte(body)

	// We need to provide a variable where the JSON
	// package can put the decoded data. This
	// `map[string]interface{}` will hold a map of strings
	// to arbitrary data types.
	//	var dat map[]interface{}
	// var cluster{}
	// Here's the actual decoding, and a check for
	// associated errors.

	cluster := &Member{}
	err = json.Unmarshal(byt, &cluster)

	message := &Message{}
	err = json.Unmarshal(byt, &message)
	if !reflect.DeepEqual(new(Member), cluster) {
		fmt.Println("equal")
		return cluster, nil
	} else if message != nil {
		return new(Member), &errorAddEtcdMember{message.Txt}

	}

	return new(Member), &errorAddEtcdMember{err.Error()}

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
