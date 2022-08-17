package framework

import (
	"errors"
	"strings"
)

// Tree 空的根节点
type Tree struct {
	root *node
}

type node struct {
	isLast   bool                //是否可以形成
	segment  string              //存储uri中的段
	handlers []ControllerHandler //控制器
	childs   []*node             //子节点
	parent   *node
}

func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
	}
}

func NewTree() *Tree {
	return &Tree{
		root: newNode(),
	}
}

func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}
	if isWildSegment(segment) {
		return n.childs
	}

	//未传0，有报错，第一个是
	allNodes := make([]*node, 0, len(n.childs))
	for _, v := range n.childs {
		if isWildSegment(v.segment) {
			allNodes = append(allNodes, v)
		} else if v.segment == segment {
			allNodes = append(allNodes, v)
		}
	}
	return allNodes
}

func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)
	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	cnodes := n.filterChildNodes(segment)
	if cnodes == nil || len(cnodes) == 0 {
		return nil
	}
	if len(segments) == 1 {
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}
		return nil
	}
	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil
}

// 将 uri 解析为 params
func (n *node) parseParamsFromEndNode(uri string) map[string]string {
	ret := map[string]string{}
	segments := strings.Split(uri, "/")
	cnt := len(segments)
	cur := n
	for i := cnt - 1; i >= 0; i-- {
		if cur.segment == "" {
			break
		}
		// 如果是通配符节点
		if isWildSegment(cur.segment) {
			// 设置 params
			ret[cur.segment[1:]] = segments[i]
		}
		cur = cur.parent
	}
	return ret
}
func (t *Tree) addRouter(uri string, handlers []ControllerHandler) error {
	n := t.root
	if n.matchNode(uri) != nil {
		return errors.New("route exist:" + uri)
	}
	segments := strings.Split(uri, "/")
	for index, segment := range segments {
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1
		var objNode *node
		childNodes := n.filterChildNodes(segment)
		if len(childNodes) > 0 {
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}
		if objNode == nil {
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handlers = handlers
			}
			cnode.parent = n
			n.childs = append(n.childs, cnode)
			objNode = cnode

		}
		n = objNode
	}

	return nil
}
func (t *Tree) FindHandler(uri string) []ControllerHandler {
	matchNode := t.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handlers
}
