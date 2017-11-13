package expo

import (
	"strings"
)

// NoExpand means no expansion.
var NoExpand = ExpandOptions([]string{})

// ExpandOptions represent options of reference need to expand.
// The key "product" means expand the product object.
// The key "product.advertiser" means expand the advertiser, the parent also ("product").
type ExpandOptions []string

// Expand check whether the reference based on key need to expand.
func (o ExpandOptions) Expand(key string) bool {
	for _, v := range o {
		s := strings.SplitN(v, ".", 2)
		if s[0] == key {
			return true
		}
	}
	return false
}

// Sub is the sub options based on key.
// Sub("product") to "product.advertiser" and "product.charging" will result "advertiser" and "charging".
func (o ExpandOptions) Sub(key string) ExpandOptions {
	var subKeys []string
	for _, v := range o {
		s := strings.SplitN(v, ".", 2)
		if len(s) == 1 {
			continue
		}

		if s[0] == key {
			subKeys = append(subKeys, s[1])
		}
	}

	if len(subKeys) == 0 {
		return NoExpand
	}

	return subKeys
}
