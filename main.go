package main

import (
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
}

// Append adds a post to the feed
func (f *Feed) Append(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
	} else {
		currentPost := f.start
		// keep looping through until you get to the end
		for currentPost.next != nil {
			currentPost = currentPost.next
		}
		// by the time its get here next is going to be nil ready for the newPost
		currentPost.next = newPost
	}
	f.length++
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
