package huffman

import (
	"sort"
)

func max(i, j uint32) uint32 {
	if i > j {
		return i
	} else {
		return j
	}
}

type node struct {
	letter    string
	frequency float32
	depth     uint32
	zero      *node
	one       *node
}

type byFrequency []node

func (b byFrequency) Len() int      { return len(b) }
func (b byFrequency) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byFrequency) Less(i, j int) bool {
	if b[i].frequency == b[j].frequency {
		return b[i].depth < b[j].depth
	} else {
		return b[i].frequency < b[j].frequency
	}
}

func createTree(input string) node {
	counts := make(map[string]int)
	for _, c := range input {
		counts[string(c)] = counts[string(c)] + 1
	}

	length := len(input)
	nodes := make([]node, 0)
	for k, v := range counts {
		leaf := node{k, float32(v) / float32(length), 1, nil, nil}
		nodes = append(nodes, leaf)
	}

	for len(nodes) > 1 {
		sort.Sort(byFrequency(nodes))
		newNode := node{"", nodes[0].frequency + nodes[1].frequency, max(nodes[0].depth, nodes[1].depth) + 1, &nodes[0], &nodes[1]}
		nodes = append(nodes[2:], newNode)
	}

	return nodes[0]
}
