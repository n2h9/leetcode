package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	obj := Constructor()
	// empty
	assert.Equal(t, []int{}, obj.GetNewsFeed(1))
	obj.PostTweet(1, 5)
	// follow itself
	obj.Follow(1, 1)
	assert.Equal(t, []int{5}, obj.GetNewsFeed(1))
	obj.Follow(1, 2)
	obj.PostTweet(2, 6)
	assert.Equal(t, []int{6, 5}, obj.GetNewsFeed(1))
	obj.Unfollow(1, 2)
	assert.Equal(t, []int{5}, obj.GetNewsFeed(1))
	obj.PostTweet(1, 7)
	obj.PostTweet(1, 8)
	obj.PostTweet(1, 9)
	obj.PostTweet(1, 10)
	obj.PostTweet(1, 11)
	obj.PostTweet(1, 12)
	obj.PostTweet(1, 13)
	obj.PostTweet(1, 14)
	obj.PostTweet(1, 15)
	obj.PostTweet(1, 16)
	obj.PostTweet(1, 17)
	obj.PostTweet(1, 18)
	assert.Equal(t, []int{18, 17, 16, 15, 14, 13, 12, 11, 10, 9}, obj.GetNewsFeed(1))
	obj.Follow(1, 3)
	obj.Follow(1, 4)
	obj.PostTweet(3, 20)
	obj.PostTweet(4, 21)
	assert.Equal(t, []int{21, 20, 18, 17, 16, 15, 14, 13, 12, 11}, obj.GetNewsFeed(1))
	assert.Equal(t, []int{20}, obj.GetNewsFeed(3))
	obj.Follow(3, 1)
	assert.Equal(t, []int{20, 18, 17, 16, 15, 14, 13, 12, 11, 10}, obj.GetNewsFeed(3))
	obj.Unfollow(1, 3)
	assert.Equal(t, []int{20, 18, 17, 16, 15, 14, 13, 12, 11, 10}, obj.GetNewsFeed(3))
	assert.Equal(t, []int{21, 18, 17, 16, 15, 14, 13, 12, 11, 10}, obj.GetNewsFeed(1))
}
