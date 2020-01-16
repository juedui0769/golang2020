package _map

type OrderedMap struct {
	keys []interface{}
	m map[interface{}]interface{}
}

func (omap *OrderedMap) Len() int {
	return len(omap.keys)
}

func (omap *OrderedMap) Less(i, j int) bool {
	return false
}

func (omap *OrderedMap) Swap(i, j int) {
	omap.keys[i], omap.keys[j] = omap.keys[j], omap.keys[i]
}




