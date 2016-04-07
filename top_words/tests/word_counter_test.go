package tests


import (
	"testing"
	"container/heap"
	"github.com/stgleb/concurrent-map"
	"github.com/stretchr/testify/assert"
	"github.com/top_words/top_words"
)


func TestTopN(t *testing.T){
	wordsMap := cmap.New()
	wordsMap.Set("pear", 1)
	wordsMap.Set("banana", 3)
	wordsMap.Set("apple", 2)
	wordsMap.Set("cherry", 4)

	pq := make(top_words.PriorityQueue, wordsMap.Count())
	i := 0

	for iter := range wordsMap.Iter() {
		item := top_words.NewItem(iter.Key, iter.Val.(int), i)
		pq[i] = item
		i++
	}
	heap.Init(&pq)
	expected_len := 2
	result := top_words.TopN(expected_len)
	assert.Equal(t, len(result), expected_len)
}
