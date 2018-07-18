// zk_operator create by xjy on 2017.10
package gzk

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var (
	TESTIP = []string{"192.168.137.89:2181"} //虚拟机ip and clientport
	conn   *zk.Conn
)

//初始化连接
func InitConn() error {
	var err error
	conn, _, err = zk.Connect(TESTIP, 1*time.Second)
	if err != nil {
		fmt.Println("InitConn:", err.Error())
		return err
	}
	return nil
}

//列出父节点的所有子节点
func ChildList(parentPath string) ([]string, error) {
	exist, err := IsExist(parentPath)
	if err != nil {
		fmt.Println("ChildList IsExist:", err.Error())
		return nil, err
	}
	if exist == false {
		fmt.Println(parentPath + " has not exist.")
		return nil, nil
	}
	children, _, err := conn.Children(parentPath)
	if err != nil {
		fmt.Println("ChildList:", err.Error())
		return nil, err
	}
	fmt.Printf("[%s] has [%d] children: %+v\n", parentPath, len(children), children)
	return children, nil
}

//得到某个节点的数据
func GetNodeData(nodePath string) (string, error) {
	exist, err := IsExist(nodePath)
	if err != nil {
		fmt.Println("ChildList IsExist:", err.Error())
		return "", err
	}
	if exist == false {
		fmt.Println(nodePath + " has not exist.")
		return "", nil
	}
	data, _, err := conn.Get(nodePath)
	if err != nil {
		fmt.Println("GetNodeData:", err.Error())
		return "", err
	}
	fmt.Printf("Data of [%s] exist: %+v\n", nodePath, string(data))
	return string(data), nil
}

//查看某个节点是否存在
func IsExist(nodePath string) (bool, error) {
	flag, _, err := conn.Exists(nodePath)
	if err != nil {
		fmt.Println("IsExist:", err.Error())
		return false, err
	}
	fmt.Printf("[%s] exist: %+v\n", nodePath, flag)
	return flag, nil
}

//创建一个节点
func CreatNode(nodePath, data string, ephemeral int32) error { //ephemeral 只有两个取值 0、1，==flags表示是否为临时节点 1是 0否
	_, err := conn.Create(nodePath, []byte(data), ephemeral, zk.WorldACL(zk.PermAll)) //permission权限
	if err != nil {
		if err.Error() == "zk: node already exists" {
			fmt.Println("zk: node already exists.")
			return nil
		}
		fmt.Println("CreatNode:", err.Error())
		return err
	}
	fmt.Printf("[%s] created, data is [%s]\n", nodePath, data)
	return nil
}

//设置节点的数据值
func SetNode(nodePath, data string) error {
	_, err := conn.Set(nodePath, []byte(data), -1)
	if err != nil {
		fmt.Println("SetNode:", err.Error())
		return err
	}
	fmt.Printf("[%s] set, data is [%s]\n", nodePath, data)
	return nil
}

//删除节点
func DeleteNode(nodePath string) error {
	err := conn.Delete(nodePath, -1)
	if err != nil {
		fmt.Println("DeleteNode:", err.Error())
		return err
	}
	fmt.Printf("[%s] del!\n", nodePath)
	return nil
}
