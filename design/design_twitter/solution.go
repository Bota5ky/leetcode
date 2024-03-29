package design_twitter

import "sort"

// Twitter 355. 设计推特 https://leetcode.cn/problems/design-twitter/
type Twitter struct {
	post    map[int][]tweets
	follows map[int][]int
	time    int
}
type tweets struct {
	tweetID int
	time    int
}

func Constructor() Twitter {
	var t Twitter
	t.follows = make(map[int][]int)
	t.post = make(map[int][]tweets)
	return t
}

func (t *Twitter) PostTweet(userID int, tweetID int) {
	t.time++
	t.post[userID] = append(t.post[userID], tweets{tweetID, t.time})
}

func (t *Twitter) GetNewsFeed(userID int) []int {
	var res []tweets
	followList := t.follows[userID]
	followList = append(followList, userID)
	for i := 0; i < len(followList); i++ {
		PostTweets := t.post[followList[i]]
		if len(PostTweets) > 10 {
			res = append(res, PostTweets[len(PostTweets)-10:]...)
		} else {
			res = append(res, PostTweets...)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].time > res[j].time
	})
	if len(res) > 10 {
		res = res[:10]
	}
	ret := make([]int, len(res))
	for i := 0; i < len(res); i++ {
		ret[i] = res[i].tweetID
	}
	return ret
}

func (t *Twitter) Follow(followerID int, followeeID int) {
	if followerID == followeeID {
		return
	}
	for i := 0; i < len(t.follows[followerID]); i++ {
		if t.follows[followerID][i] == followeeID {
			return
		}
	}
	t.follows[followerID] = append(t.follows[followerID], followeeID)
}

func (t *Twitter) Unfollow(followerID int, followeeID int) {
	for i := 0; i < len(t.follows[followerID]); i++ {
		if t.follows[followerID][i] == followeeID {
			t.follows[followerID] = append(t.follows[followerID][:i], t.follows[followerID][i+1:]...)
		}
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userID,tweetID);
 * param_2 := obj.GetNewsFeed(userID);
 * obj.Follow(followerID,followeeID);
 * obj.Unfollow(followerID,followeeID);
 */
