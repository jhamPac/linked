package main

import (
	"errors"
	"fmt"
	"time"
)

// Post represents a tweet, FB post etc
type Post struct {
	body        string
	publishDate int64
	next        *Post
}

// NewPost factory function to create a new Post
func NewPost(b string, t int64) *Post {
	return &Post{body: b, publishDate: t}
}

// Feed is like a news feed on social media
type Feed struct {
	length int
	start  *Post
	end    *Post // big-O: O(1) constant time
}

// Append adds a post to the feed
func (f *Feed) Append(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
		f.end = newPost
	} else {
		lastPost := f.end
		lastPost.next = newPost
		f.end = newPost
	}
	f.length++
}

// Remove deletes a post from the feed
func (f *Feed) Remove(publishDate int64) {
	if f.length == 0 {
		panic(errors.New("Feed is empty"))
	}

	var previousPost *Post
	currentPost := f.start

	for currentPost.publishDate != publishDate {
		if currentPost.next == nil {
			panic(errors.New("No such Post found"))
		}

		previousPost = currentPost
		currentPost = currentPost.next
	}
	previousPost.next = currentPost.next
	f.length--

}

func main() {
	f := &Feed{}
	post := NewPost("You have gone incognito", time.Now().Unix())

	f.Append(post)

	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First post: %v\n", f.start)

	p2 := NewPost("No one can see you", time.Now().Unix())

	f.Append(p2)

	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First post with next pointer: %v\n", f.start)
	fmt.Printf("Second post: %v\n", f.start.next)
}
