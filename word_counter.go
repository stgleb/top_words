package top_words

import (
	"github.com/streamrail/concurrent-map"
	"container/heap"
)


var wordsMap = cmap.New()

func TopN(n int) []string {
	var words [n]string
	pq := make(PriorityQueue, wordsMap.Count())
	i := 0

	for iter := range wordsMap.Iter() {
		item := &Item{
			value:    iter.Key,
			priority: iter.Val,
			index:    i,
		}
		pq[i] = item
		i++
	}
	heap.Init(&pq)

	for i := 0; i < n; i++ {
		item := heap.Pop(&pq).(*Item)
		words[i] = item.value
	}

	return words
}
