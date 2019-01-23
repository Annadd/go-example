package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	FL_IP             = "localhost"
	FL_PORT           = "8080"
	FL_DIR            = "http://" + FL_IP + ":" + FL_PORT + "/wm"
	PATH_ADD_GROUP    = "/staticgroup/add"
	PATH_MODIFY_GROUP = "/staticgroup/modify"
	PATH_DELETE_GROUP = "/staticgroup/delete"
	PATH_ADD_FLOW     = "/staticflowpusher/json/store"
	PATH_DEL_FLOW     = "/staticflowpusher/json/delete"
)

type ControllerGroup struct {
	GroupID    int    `json:"group_id"`
	Type       string `json:"type"`
	Dpid       string `json:"dpid"`
	Vlan       int    `json:"vlan"`
	BucketType string `json:"bucket_type"`
	Bucket     []int  `json:"bucket"`
}

type ControllerFlow struct {
	Dpid                    string `json:"switch"`
	Name                    string `json:"name"`
	Active                  bool   `json:"active"`
	Priority                int    `json:"priority"`
	Table                   int    `json:"table"`
	HardTimeout             int    `json:"hard_timeout"`
	InPort                  int    `json:"in_port"`
	SrcModPort              int    `json:"src_mod_port"`
	ActSetOutput            int    `json:"actset_output"`
	EthTyep                 int    `json:"eth_type"`
	EthSrc                  string `json:"eth_src"`
	EthDst                  string `json:"eth_dst"`
	EthVlanVid              int    `json:"eth_vlan_vid"`
	Proto                   int    `json:"ip_proto"`
	Ipv4Src                 string `json:"ipv4_src"`
	Ipv4Dst                 string `json:"ipv4_dst"`
	InstructionApplyActions string `json:"instruction_apply_actions"`
	InstructionGotoTable    int    `json:"instruction_goto_table"`
}

type CID struct {
	Device string `json:"device"`
	Port   string `json:"Port"`
}

type Snapshot struct {
	Cid             []CID  `json:"CID"`
	DstIp           string `json:"dst_ip"`
	SrcIp           string `json:"src_ip"`
	EthTyep         string `json:"ethtype"`
	Hardwaretimeout int32  `json:"hardware_timeout"`
	OutDevice       string `json:"out_device"`
	OutPort         int32  `json:"out_port"`
}

func makeMCGroup(id int, dpid string) ControllerGroup {
	bucket := []int{11, 12, 13, 14}

	group := ControllerGroup{
		GroupID: 4096 + id,
		Type:    "all",
		Dpid:    dpid,
		Bucket:  bucket,
	}
	return group
}

func makeRWGroup(id int, dpid string, memberType string, member []int) ControllerGroup {
	group := ControllerGroup{
		GroupID:    id,
		Type:       "indirect",
		Dpid:       dpid,
		BucketType: memberType,
		Bucket:     member,
	}
	return group
}

func makeFlow(inPort int, dpid string, actions string) ControllerFlow {
	flow := ControllerFlow{
		Name:                    "flow_name_" + time.Now().String(),
		Active:                  true,
		Priority:                2,
		Table:                   60,
		Dpid:                    dpid,
		InPort:                  inPort,
		InstructionApplyActions: actions,
	}
	return flow
}

func callGroup(group ControllerGroup, operation string) {
	content, err := json.Marshal(group)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
	var resp *http.Response

	switch {
	case operation == "add":
		resp, err = http.Post(FL_DIR+PATH_ADD_GROUP, string(content), nil)
	case operation == "modify":
		resp, err = http.Post(FL_DIR+PATH_MODIFY_GROUP, string(content), nil)
	case operation == "delete": //put
		payload := strings.NewReader(string(content))
		request, _ := http.NewRequest("PUT", FL_DIR+PATH_DELETE_GROUP, payload)
		resp, err = http.DefaultClient.Do(request)
	default:
		fmt.Println("operation is error")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func callFlow(flow ControllerFlow, operation string) {
	content, err := json.Marshal(flow)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
	var resp *http.Response
	switch {
	case operation == "add":
	case operation == "modify":
		resp, err = http.Post(FL_DIR+PATH_ADD_FLOW, string(content), nil)
	case operation == "delete":
		resp, err = http.Post(FL_DIR+PATH_DEL_FLOW, string(content), nil)
	default:
		fmt.Println("operation is error")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func main() {
	dpid := "00:00:00:10:a0:2f:c0:93"
	mcGroup := makeMCGroup(1, dpid)
	member := make([]int, 1)
	member = append(member, mcGroup.GroupID)
	rwGroup := makeRWGroup(1, dpid, "mc", member)
	callGroup(mcGroup, "add")
	callGroup(rwGroup, "add")

	host := "192.168.3.58:1993"
	url := "http://" + host + "/matrix/snapshot"
	buf, err := ioutil.ReadFile("C:\\Users\\Administrator\\Desktop\\snapshot.json")
	if err != nil {
		fmt.Println(err)
	}

	resp, err := http.Post(url, string(buf), nil)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
