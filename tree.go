package huffman

import (
	"fmt"
	"sort"
)

func max(i, j uint32) uint32 {
	if i > j {
		return i
	}
	return j
}

type node struct {
	letter    rune
	frequency float32
	depth     uint32
	zero      *node
	one       *node
}

func (n *node) Leaf() bool {
	return n.one == nil && n.zero == nil
}

func (n *node) String() string {
	if n.one != nil && n.zero != nil {
		return fmt.Sprint("{ \"letter\": \"", n.letter, "\", \"frequency\": ", n.frequency, ", \"depth\": ", n.depth, ", \"one\": ", *n.one, ", \"zero\": ", *n.zero, " }")
	}
	return fmt.Sprint("{ \"letter\": \"", n.letter, "\", \"frequency\": ", n.frequency, ", \"depth\": ", n.depth, " }")
}

func (n *node) less(other *node) bool {
	if n.frequency == other.frequency {
		return n.depth < other.depth
	}
	return n.frequency < other.frequency
}

type byFrequency []node

func (b byFrequency) Len() int           { return len(b) }
func (b byFrequency) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byFrequency) Less(i, j int) bool { return b[i].less(&b[j]) }

func fromInput(input string) *node {
	counts := make(map[rune]int)
	for _, c := range input {
		counts[c] = counts[c] + 1
	}

	length := len(input)
	nodes := make([]node, 0)
	for k, v := range counts {
		leaf := node{k, float32(v) / float32(length), 0, nil, nil}
		nodes = append(nodes, leaf)
	}

	for len(nodes) > 1 {
		sort.Sort(byFrequency(nodes))
		newNode := node{-1, nodes[0].frequency + nodes[1].frequency, max(nodes[0].depth, nodes[1].depth) + 1, &nodes[0], &nodes[1]}
		nodes = append(nodes[2:], newNode)
	}

	return &nodes[0]
}

func fromMapping(mapping map[rune][]byte) node {
	root := node{-1, 1.0, 0, nil, nil}
	for letter, code := range mapping {
		n := &root
		numberOfBits := code[len(code)-1]
		for _, bit := range code[:len(code)-1] {
			mask := byte(1)
			for i := 0; i < 8; i++ {
				if numberOfBits > 0 {
					if (bit & (mask << i)) != 0 {
						if n.one == nil {
							n.one = &node{-1, 0.0, 0, nil, nil}
						}
						n = n.one
					} else {
						if n.zero == nil {
							n.zero = &node{-1, 0.0, 0, nil, nil}
						}
						n = n.zero
					}
					numberOfBits--
				}
			}
		}
		n.letter = letter
	}

	return root
}
