package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	data "github.com/pradyumna001/blog-app/Data"
	models "github.com/pradyumna001/blog-app/Models"
)

//PostsController.go
var DB = *data.DB

type PostsController struct{}

//Create Post
func (pc *PostsController) CreateRecord(w http.ResponseWriter, r *http.Request) {

	vals := r.URL.Query()

	author, authorOk := vals["author"]
	postedOnStr, postedOnOk := vals["posted_on"]
	title, titleOk := vals["title"]
	body, bodyOk := vals["body"]

	post := models.Post{
		Author:   author[0],
		PostedOn: postedOnStr[0],
		Title:    title[0],
		Body:     body[0],
	}

	validationErr := models.IsValidPost(post)
	//date validation

	layout := "2006-01-02"
	postedOn, err := time.Parse(layout, postedOnStr[0])
	if err != nil {
		validationErr["Date Of Birth"] = "Date Should be Valid"
	}
	if len(validationErr) > 0 {

		usersJson, err := json.Marshal(validationErr)
		if err != nil {
			log.Fatal("Error  :: ", err)
			fmt.Fprint(w, "Error  :: ", err)
		}
		_, err = w.Write(usersJson)
		if err != nil {
			log.Fatal("Error  :: ", err)
			fmt.Fprint(w, "Error  :: ", err)
		}
		return
	}

	if authorOk && postedOnOk && titleOk && bodyOk {

		id := 0
		err := DB.QueryRow("INSERT into post_table(author,posted_on,title,Body) values( $1 , $2 , $3 ,  $4  ) RETURNING pid", string(author[0]), postedOn, string(title[0]), string(body[0])).Scan(&id)
		if err != nil {
			log.Print("error executing query :: ", err)
			return
		}
		fmt.Fprint(w, "Record Successfully Inserted.\nLast Inserted Record is  :: ", id)

	} else {
		fmt.Fprintf(w, "Error occurred while creating record in database for name :: %s", author[0])
	}
}

//Read All Posts
func (pc *PostsController) ReadRecords(w http.ResponseWriter, r *http.Request) {
	log.Print("reading records from database")
	rows, err := DB.Query("SELECT * FROM post_table")

	if err != nil {
		log.Print("error executing select query :: ", err)
		return
	}
	posts := []models.Post{}
	for rows.Next() {
		var id string
		var author string
		var postedOn string
		var title string
		var body string
		err = rows.Scan(&id, &author, &postedOn, &title, &body)
		if err != nil {
			panic(err)
		}
		newpost := models.Post{Id: id, Author: author, PostedOn: postedOn, Title: title, Body: body}
		posts = append(posts, newpost)
	}
	json.NewEncoder(w).Encode(posts)
}

//Update Posts
func (pc *PostsController) UpdateRecord(w http.ResponseWriter, r *http.Request) {

	vals := r.URL.Query()

	id, idOk := vals["id"]
	author, authorOk := vals["author"]
	postedOnStr, postedOnOk := vals["posted_on"]
	title, titleOk := vals["title"]
	body, bodyOk := vals["body"]

	post := models.Post{
		Author:   author[0],
		PostedOn: postedOnStr[0],
		Title:    title[0],
		Body:     body[0],
	}

	validationErr := models.IsValidPost(post)
	//date validation

	layout := "2006-01-02"
	postedOn, err := time.Parse(layout, postedOnStr[0])
	if err != nil {
		validationErr["Date Of Birth"] = "Date Should be Valid"
	}
	if len(validationErr) > 0 {

		usersJson, err := json.Marshal(validationErr)
		if err != nil {
			log.Fatal("Error  :: ", err)
			fmt.Fprint(w, "Error  :: ", err)
		}
		_, err = w.Write(usersJson)
		if err != nil {
			log.Fatal("Error  :: ", err)
			fmt.Fprint(w, "Error  :: ", err)
		}
		return
	}

	if idOk && authorOk && postedOnOk && titleOk && bodyOk {
		log.Print("going to update record in database for id :: ", id)

		result, err := DB.Exec("UPDATE post_table SET author = $1 , posted_on = $2 , title = $3 , body = $4 where pid = $5", author[0], postedOn, title[0], body[0], id[0])
		if err != nil {
			log.Print("error occurred while executing query :: ", err)
			return
		}
		rowsAffected, _ := result.RowsAffected()

		fmt.Fprintf(w, "Number of rows updated in database are :: %d", rowsAffected)
	} else {
		fmt.Fprintf(w, "Error occurred while updating record in database for id :: %s", id)
	}
}

//Delete Post
func (pc *PostsController) DeleteRecord(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	id, idok := vals["id"]
	if idok {
		log.Print("going to delete record in database for name :: ", id[0])

		result, err := DB.Exec("DELETE from post_table where pid=$1", id[0])
		if err != nil {
			log.Print("error occurred while executing query :: ", err)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Print("error occurred while stroeing Rowaffected :: ", err)
			return
		}
		fmt.Println(w, "Number of rows deleted in database are :: ", rowsAffected)
		fmt.Fprint(w, id[0])
	} else {
		fmt.Fprint(w, "Error occurred while deleting record in database for id ", id[0])
	}
}
