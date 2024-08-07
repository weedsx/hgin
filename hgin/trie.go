package hgin

import "strings"

type node struct {
	pattern  string  // 完整待匹配路由，只有末尾节点保存完整路由，例如 /p/:lang/doc，只有当前 part 为 doc 时，pattern 才为 /p/:lang/doc
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [p, :lang, doc]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// insert 将一个路由插入到路由前缀树
//
// pattern: 待匹配路由，例如 /p/:lang，
// parts: 拆分后的路由，例如 [p, :lang]，
// height: 当前递归到哪一层，从 0 开始
func (n *node) insert(pattern string, parts []string, height int) {
	if height == len(parts) {
		// 递归结束，在末尾节点设置完整路由
		n.pattern = pattern
		return
	}

	nextPart := parts[height] // 下一部分路由
	child := n.matchChild(nextPart)
	if child == nil {
		child = &node{
			part:   nextPart,
			isWild: nextPart[0] == ':' || nextPart[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// search
// parts 拆分后的路由，例如 [p, :lang]
// height 当前递归到哪一层
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	nextPart := parts[height]
	children := n.matchChildren(nextPart)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
