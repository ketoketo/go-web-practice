package main

import (
	"fmt"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post){
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main(){
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id:1, Content: "hello", Author: "test1"}
	post2 := Post{Id:2, Content: "hello2", Author: "test13"}
	post3 := Post{Id:3, Content: "hello3", Author: "test14"}
	post4 := Post{Id:4, Content: "hello4", Author: "test13"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])
	fmt.Println(&post1)

	for _, post := range PostsByAuthor["test13"] {
		fmt.Println(post)
	}

	for _, post := range PostsByAuthor["test14"] {
		fmt.Println(post)
	}
}
