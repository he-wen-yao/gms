package ds

import "strings"

type Node struct {
	Pattern  string  // 待匹配路由，例如 /p/:lang
	Part     string  // 路由中的一部分，例如 :lang
	Children []*Node // 子节点，例如 [doc, tutorial, intro]
	IsWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (n *Node) matchChild(part string) *Node {
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *Node) matchChildren(part string) []*Node {
	nodes := make([]*Node, 0)
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *Node) Insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.Pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &Node{Part: part, IsWild: part[0] == ':' || part[0] == '*'}
		n.Children = append(n.Children, child)
	}
	child.Insert(pattern, parts, height+1)
}

func (n *Node) Search(parts []string, height int) *Node {
	if len(parts) == height || strings.HasPrefix(n.Part, "*") {
		if n.Pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.Search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
