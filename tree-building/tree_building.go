package tree

import (
	"errors"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func FindIdInTree(node *Node, id int) bool {
	if node.ID == id {
		return true
	}
	for _, children := range node.Children {
		if FindIdInTree(children, id) {
			return true
		}
	}
	return false
}

func InsertInTree(node *Node, record Record) (bool, error) {
	if record.Parent == node.ID {
		k := len(node.Children)
		for i := 0; i < len(node.Children); i++ {
			if record.ID == node.Children[i].ID {
				return false, errors.New("duplicate node")
			}
			if record.ID < node.Children[i].ID {
				k = i
				break
			}
		}
		// insert node to children on k-position
		node.Children = append(node.Children[:k], append([]*Node{{ID: record.ID}}, node.Children[k:]...)...)
		return true, nil
	}
	for _, children := range node.Children {
		result, err := InsertInTree(children, record)
		if result || err != nil {
			return result, err
		}
	}
	return false, nil
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	root := &Node{ID: 0}
	n := len(records)
	hasRoot := false

	for i := 0; i < n; i++ {
		for j := 0; j < len(records); j++ {
			if records[j].ID < records[j].Parent {
				return nil, errors.New("higher id parent of lower id")
			}
			if records[j].ID == 0 {
				if hasRoot {
					return nil, errors.New("duplicate root")
				}
				hasRoot = true
				if records[j].Parent > 0 {
					return nil, errors.New("root node has no parent")
				}
				// delete j-element
				records = append(records[:j], records[j+1:]...)
				continue
			}
			q, err := InsertInTree(root, records[j])
			if err != nil {
				return nil, err
			}
			if q {
				// delete j-element
				records = append(records[:j], records[j+1:]...)
				break
			}
		}
	}
	if !hasRoot {
		return nil, errors.New("no root node")
	}
	for i := 0; i < n; i++ {
		if !FindIdInTree(root, i) {
			return nil, errors.New("non-continuous")
		}
	}

	return root, nil
}
