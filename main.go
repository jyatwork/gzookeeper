package main

import (
	"fmt"
	"github_atwork/gzookeeper"
)

func main() {
	err := gzk.InitConn()
	if err != nil {
		fmt.Println("InitConn error:", err.Error())
	}
	children, _ := gzk.ChildList("/go_servers") //has [3] children: [server3 server1 server2]
	fmt.Println(children)                       //[server3 server1 server2]
	for _, v := range children {
		str, _ := gzk.GetNodeData("/go_servers" + "/" + v)
		fmt.Println(str) //111 222 333
	}
	gzk.CreatNode("/2018", "2018data", 0)
	children2, _ := gzk.ChildList("/") // has [6] children: [ElectMasterDemo mirror 2018 go_servers test zookeeper]
	fmt.Println(children2)             //[ElectMasterDemo mirror 2018 go_servers test zookeeper]
}
