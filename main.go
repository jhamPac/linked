package main

import "fmt"

type Post struct {
	body        string
	publishDate int64
	next        *Post
}

type Feed struct {
	length int
	start  *Post
}

func (f *Feed) Append(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
	} else {
		currentPost := f.start
		for currentPost.next != nil {
			currentPost = currentPost.next
		}
		currentPost.next = newPost
	}
	f.length++
}

func main() {
	f := &Feed{}
	post := &Post{body: "You have gone incognito"}

	f.Append(post)

	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First post: %s\n", f.start.body)

	p2 := &Post{body: "You are not seen"}

	f.Append(p2)

	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("Second post: %s\n", f.start.next.body)
}
