package simutil

import "fmt"

type Node interface {
	Value() interface{}
	Children() []Node
}

func _printNode(node Node, prefix string, tail bool) {
	var line string

	if tail {
		line = "└── "
	} else {
		line = "├── "
	}

	fmt.Printf("%s%s%s\n", prefix, line, node.Value())

	if len(node.Children()) > 0 {
		for i := 0; i < len(node.Children()) - 1; i++ {
			var childNode = node.Children()[i]

			if tail {
				line = "    "
			} else {
				line = "│   "
			}

			_printNode(childNode, fmt.Sprintf("%s%s", prefix, line), false)
		}
		if len(node.Children()) >= 1 {
			var lastNode = node.Children()[len(node.Children()) - 1]

			if tail {
				line = "    "
			} else {
				line = "│   "
			}

			_printNode(lastNode, fmt.Sprintf("%s%s", prefix, line), true)
		}
	}
}

func PrintNode(node Node) {
	_printNode(node, "", true)
}
