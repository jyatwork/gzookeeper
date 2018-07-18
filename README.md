# gzookeeper
zookeeper operator include CreatNode, GetNode, SetNode, DeleteNode, ChildList and so on. use "github.com/samuel/go-zookeeper/zk"

##example
--------------

package main


import (
	"fmt"
	"github_atwork/gzookeeper"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	err := gzk.InitConn()
	if err != nil {
		fmt.Println("InitConn error:", err.Error())
	}
	//watcher
	change, errors := gzk.Watcher("/go_servers")
	go func() {
		for {
			select {
			case c := <-change:
				fmt.Printf("watcher recieved %+v\n", c)
			case err := <-errors:
				fmt.Printf("watcher recieved err %+v\n", err)
			}
		}
	}()

	//normal operator
	children, _ := gzk.ChildList("/go_servers") //has [3] children: [server3 server1 server2]
	fmt.Println(children)                       //[server3 server1 server2]
	for _, v := range children {
		str, _ := gzk.GetNodeData("/go_servers" + "/" + v)
		fmt.Println(str) //111 222 333
	}
	gzk.CreatNode("/2018", "2018data", 0)
	children2, _ := gzk.ChildList("/") // has [6] children: [ElectMasterDemo mirror 2018 go_servers test zookeeper]
	fmt.Println(children2)             //[ElectMasterDemo mirror 2018 go_servers test zookeeper]
	flags := int32(zk.FlagEphemeral)
	gzk.CreatNode("/go_servers/test", "test", flags)
	select {}
}
----------