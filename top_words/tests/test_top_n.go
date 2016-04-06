package tests


import (
	"testing"
	"github.com/top-words"
	"github.com/streamrail/concurrent-map"
	"github.com/stretchr/testify/assert"
	"container/heap"
)


func TestTopN(t *testing.T){
	wordsMap := cmap.New()
	wordsMap.Set("pear", 1)
	wordsMap.Set("banana", 3)
	wordsMap.Set("apple", 2)
	wordsMap("cherry", 4)

	pq := make(top_words.PriorityQueue, wordsMap.Count())
	i := 0

	for iter := range wordsMap.Iter() {
		item := &top_words.Item{
			value:    iter.Key,
			priority: iter.Val,
			index:    i,
		}
		pq[i] = item
		i++
	}
	heap.Init(&pq)
	expected_len := 2
	result := top_words.TopN(expected_len, cmap)
	assert.Equal(t, len(result), expected_len)
}
