package main

import "log"

// 게시글 타입
type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 데이터 베이스 hard-coded
var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func getAllArticle() []article {
	log.Println(articleList)
	return articleList
}
