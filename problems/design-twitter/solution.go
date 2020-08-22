package solution

const FEED_LEN = 10

type Twitter struct {
	tweets  []int
	feedLen int
	// userId : indexes from tweets
	userTweets    map[int][]int
	userFollowees map[int]map[int]bool
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		tweets:        make([]int, 0),
		feedLen:       FEED_LEN,
		userTweets:    make(map[int][]int),
		userFollowees: make(map[int]map[int]bool),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweets = append(this.tweets, tweetId)
	if this.userTweets[userId] == nil {
		this.userTweets[userId] = make([]int, 0)
	}
	this.userTweets[userId] = append(this.userTweets[userId], len(this.tweets)-1)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	h := newHeap()
	for i, k := len(this.userTweets[userId])-1, this.feedLen; i >= 0 && k > 0; i, k = i-1, k-1 {
		h.push(
			&heapItem{
				priority: this.userTweets[userId][i],
			},
		)
	}
	for followeeId := range this.userFollowees[userId] {
		for i, k := len(this.userTweets[followeeId])-1, this.feedLen; i >= 0 && k > 0; i, k = i-1, k-1 {
			h.push(
				&heapItem{
					priority: this.userTweets[followeeId][i],
				},
			)
		}
	}

	res := make([]int, 0, this.feedLen)
	for i := 0; i < this.feedLen && !h.isEmpty(); i++ {
		res = append(
			res,
			this.tweets[h.top().priority],
		)
		h.pop()
	}

	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if this.userFollowees[followerId] == nil {
		this.userFollowees[followerId] = make(map[int]bool)
	}
	this.userFollowees[followerId][followeeId] = true
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.userFollowees[followerId] == nil {
		return
	}
	delete(this.userFollowees[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

type heapItem struct {
	priority int
}

type heap struct {
	storage []*heapItem
}

func newHeap() *heap {
	return &heap{
		storage: make([]*heapItem, 0),
	}
}

func (h *heap) isEmpty() bool {
	return len(h.storage) <= 0
}

func (h *heap) push(item *heapItem) {
	h.storage = append(h.storage, item)
	ind := len(h.storage) - 1
	// parent index
	pi := (ind - 1) / 2
	// move item up until it is > then parent
	for ind > 0 && h.storage[ind].priority > h.storage[pi].priority {
		h.storage[ind], h.storage[pi] = h.storage[pi], h.storage[ind]
		ind = pi
		pi = (ind - 1) / 2
	}
}

func (h *heap) top() *heapItem {
	return h.storage[0]
}

func (h *heap) pop() {
	h.storage[0], h.storage[len(h.storage)-1] = h.storage[len(h.storage)-1], h.storage[0]
	h.storage = h.storage[:len(h.storage)-1]
	h._heapify(0)
}

func (h *heap) _heapify(ind int) {
	// left index, right index, max index
	li, ri, mi := 2*ind+1, 2*ind+2, ind

	for _, i := range []int{li, ri} {
		if i < len(h.storage) && h.storage[i].priority > h.storage[mi].priority {
			mi = i
		}
	}

	if mi != ind {
		h.storage[mi], h.storage[ind] = h.storage[ind], h.storage[mi]
		h._heapify(mi)
	}
}
