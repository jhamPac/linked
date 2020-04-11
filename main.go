package main

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

}
