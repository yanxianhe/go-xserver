package nodemgr

import (
	"github.com/fananchong/go-xserver/common"
	"github.com/fananchong/gotcp"
)

// Node : 管理节点
type Node struct {
	gotcp.Server
}

// NewNode : 管理节点实现类的构造函数
func NewNode() *Node {
	return &Node{}
}

// Init : 初始化节点
func (node *Node) Init() bool {
	return true
}

// GetID : 获取本节点信息，节点ID
func (node *Node) GetID() common.NodeID {
	return common.NodeID{}
}

// GetIP : 获取本节点信息，IP
func (node *Node) GetIP(i common.IPType) string {
	return ""
}

// GetPort : 获取本节点信息，端口
func (node *Node) GetPort(i int) int {
	return 0
}

// GetOverload : 获取本节点信息，负载
func (node *Node) GetOverload(i int) uint {
	return 0
}

// GetVersion : 获取本节点信息，版本号
func (node *Node) GetVersion() string {
	return common.XCONFIG.Common.Version
}

// SelectOne : 根据节点类型，随机选择 1 节点
func (node *Node) SelectOne(nodeType common.NodeType) common.INode {
	return node
}

// SendOne : 根据节点类型，随机选择 1 节点，发送数据
func (node *Node) SendOne(nodeType common.NodeType, data []byte) {

}

// SendByType : 对某类型节点，广播数据
func (node *Node) SendByType(nodeType common.NodeType, data []byte, exclude []common.NodeID) {

}

// SendByID : 往指定节点，发送数据
func (node *Node) SendByID(nodeID common.NodeID, data []byte) {

}

// Send : 往该节点，发送数据
func (node *Node) Send(data []byte) {

}

// SendAll : 对服务器组，广播数据
func (node *Node) SendAll(data []byte, exclude []common.NodeID) {

}

func (node *Node) RegisterOnConnect()    {}
func (node *Node) RegisterOnRecv()       {}
func (node *Node) RegisterOnDisconnect() {}
