package model

import "github.com/hatena/ipdrawer/gen/go/model"

func NetworkHasTag(n *model.Network, tag *model.Tag) bool {
	for _, t := range n.Tags {
		if t.Key == tag.Key && t.Value == tag.Value {
			return true
		}
	}
	return false
}
