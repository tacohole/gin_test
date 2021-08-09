package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "article1", Content: "some content"},
	article{ID: 2, Title: "article2", Content: "some more content"},
}

func getAllArticles() []article {
	return articleList
}
