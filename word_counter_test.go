package top_words

import (
	"container/heap"
	"github.com/stgleb/concurrent-map"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopN(t *testing.T) {
	wordsMap := cmap.New()
	wordsMap.Set("pear", 1)
	wordsMap.Set("banana", 3)
	wordsMap.Set("apple", 2)
	wordsMap.Set("cherry", 4)

	pq := make(PriorityQueue, wordsMap.Count())
	i := 0

	for iter := range wordsMap.Iter() {
		item := NewItem(iter.Key, iter.Val.(int), i)
		pq[i] = item
		i++
	}
	heap.Init(&pq)
	expected_len := 2
	result := TopN(expected_len)
	assert.Equal(t, len(result), expected_len)
}
