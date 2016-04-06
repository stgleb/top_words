package top_words

import (
	"log"
	"container/heap"
	"github.com/stgleb/concurrent-map"
)

var wordsMap = cmap.New()


func TopN(n int) []string {
	var words = make([]string, n)
	count := wordsMap.Count()

	if count == 0{
		count += count
	}

	pq := make(PriorityQueue, count)
	i := 0

	for iter := range wordsMap.Iter() {
		item := &Item{
			value:    iter.Key,
			priority: iter.Val.(int),
			index:    i,
		}
		pq[i] = item
		i++
	}
	heap.Init(&pq)
	log.Println(pq)

	for i := 0; i < n && i < len(pq); i++ {
		item := heap.Pop(&pq).(*Item)
		words[i] = item.value
	}

	return words
}
