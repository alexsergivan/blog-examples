package main

import (
	"fmt"
	"sort"
)

func main() {

	items := make(map[int]string)
	for i := 0; i < 10; i++ {
		items[i] = fmt.Sprintf("This is item %d", i)
	}
	// Generates keys map.
	keys := make([]int, len(items))
	for k := range items {
		keys[k] = k
	}
	sort.Ints(keys)
	// Iterates over items map, using sorted keys.
	for _, k := range keys {
		performItem(items[k])
	}
}

func performItem(item string) {
	fmt.Println(item)
}
