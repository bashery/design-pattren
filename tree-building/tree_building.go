package tree

import (
	"errors"
)

// Define a function Build(records []Record) (*Node, error)
// where Record is a struct containing int fields ID and Parent
// and Node is a struct containing int field ID and []*Node field Children.

type Node struct {
	ID       int
	Children []*Node
}

type Record struct {
	ID     int
	Parent int
}

func Build(records []Record) (*Node, error) {
	nodes := make([]Node, len(records))

	if len(records) == 0 {
		return nil, nil
	}

	for i, r := range records {
		if r.ID < 0 || r.ID >= len(records) {
			return nil, errors.New("Out of bounds")
		}
		if r.ID != i {
			records[i], records[r.ID] = records[r.ID], records[i]
		}

	}

	for i, r := range records {
		if r.ID != i {
			return nil, errors.New("Not in order")
		}

		// All records have parent ID lower their own ID
		if r.ID == 0 {
			if r.Parent != 0 {
				return nil, errors.New("error root id")
			}
		} else {
			if r.Parent >= r.ID {
				return nil, errors.New("Bad Parent - non root")
			}
		}
		nodes[i].ID = i
		if i != 0 {
			p := &nodes[r.Parent]
			p.Children = append(p.Children, &nodes[i])
		}
	}

	return &nodes[0], nil
}
