package destination

import "k8s.io/apimachinery/pkg/labels"

func List(selector labels.Selector) (ret []Destination, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*Destination))
	})
	return ret, err
}
