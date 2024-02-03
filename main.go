package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Post struct {
	Title       string
	Description string
}

type IndexPage struct {
	Posts map[string]Post
}

func main() {
	posts := map[string]Post{
		"a": {Title: "post 1", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique, nesciunt autem corrupti laborum, perferendis eos nemo corporis ducimus velit! Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique"},
		"b": {Title: "post 2", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique"},
		"c": {Title: "post 1 is here", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique, nesciunt autem corrupti laborum, perferendis eos nemo corporis ducimus velit! Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique"},
		"d": {Title: "Task 2", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique"},
		"e": {Title: "Task 1 is here", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique, nesciunt autem corrupti laborum, perferendis eos nemo corporis ducimus velit!"},
		"f": {Title: "Task 2", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique"},
		"g": {Title: "Task 1 is here", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique, nesciunt autem corrupti laborum, perferendis eos nemo corporis ducimus velit!"},
		"h": {Title: "Task 2", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique"},
		"i": {Title: "Task 1 is here", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique, nesciunt autem corrupti laborum, perferendis eos nemo corporis ducimus velit!"},
		"j": {Title: "Task 2", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique"},
	}
	css := http.FileServer(http.Dir("./resources/stylesheets"))
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets", css))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./resources/pages/index.html"))
		data := IndexPage{Posts: posts}
		tmpl.Execute(w, data)
	})
	http.HandleFunc("/post/", func(w http.ResponseWriter, r *http.Request) {
		postId := strings.TrimPrefix(r.URL.Path, "/post/")
		fmt.Println(postId)
		tmpl := template.Must(template.ParseFiles("./resources/pages/post.html"))
		data := posts[postId]
		tmpl.Execute(w, data)
	})

	fmt.Println("Listening on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
