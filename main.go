package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Post struct {
	Title       string
	Description string
}

// TodoPageData list of page
type TodoPageData struct {
	Posts []Post
}

func main() {
	css := http.FileServer(http.Dir("./resources/stylesheets"))
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets", css))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./resources/pages/index.html"))
		data := TodoPageData{
			Posts: []Post{
				{Title: "Task 1 is here", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique, nesciunt autem corrupti laborum, perferendis eos nemo corporis ducimus velit! Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique"},
				{Title: "Task 2", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique"},
				{Title: "Task 1 is here", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique, nesciunt autem corrupti laborum, perferendis eos nemo corporis ducimus velit! Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique"},
				{Title: "Task 2", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique"},
				{Title: "Task 1 is here", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique, nesciunt autem corrupti laborum, perferendis eos nemo corporis ducimus velit!"},
				{Title: "Task 2", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique"},
				{Title: "Task 1 is here", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique, nesciunt autem corrupti laborum, perferendis eos nemo corporis ducimus velit!"},
				{Title: "Task 2", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique"},
				{Title: "Task 1 is here", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique, nesciunt autem corrupti laborum, perferendis eos nemo corporis ducimus velit!"},
				{Title: "Task 2", Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ea itaque repudiandae ut maxime tenetur id non iste libero consequatur similique"},
			},
		}
		tmpl.Execute(w, data)
	})

	fmt.Println("Listening on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
