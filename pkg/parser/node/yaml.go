package node

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
)

var ErrFailedToMarshal = errors.New("failed to marshal to YAML")

func (node *Node) MarshalYAML() (interface{}, error) {
	switch obj := node.Value.(type) {
	case *MapValue:
		var result yaml.MapSlice

		for _, child := range node.Children {
			marshalledItem, err := child.MarshalYAML()
			if err != nil {
				return nil, err
			}

			result = append(result, yaml.MapItem{
				Key:   child.Name,
				Value: marshalledItem,
			})
		}

		return result, nil
	case *ListValue:
		var result []interface{}

		for _, child := range node.Children {
			marshalledItem, err := child.MarshalYAML()
			if err != nil {
				return nil, err
			}

			result = append(result, marshalledItem)
		}

		return result, nil
	case *ScalarValue:
		return obj.Value, nil
	default:
		return nil, fmt.Errorf("%w: unknown node type: %T", ErrFailedToMarshal, node.Value)
	}
}