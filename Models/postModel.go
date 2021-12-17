package models

type Post struct {
	Id       string `json:"id"`
	Author   string `json:"author "`
	PostedOn string `json:"postedOn"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

//Check if user has entered valid data or not
func IsValidPost(post Post) map[string]string {

	var errs = make(map[string]string)
	// check if the name empty
	if post.Author == "" {
		errs["Author1"] = "The Author is required!"
	}
	if len(post.Author) < 2 || len(post.Author) > 25 {
		errs["Title"] = "The Author field must be between 2-25 chars!"
	}
	if len(post.Title) < 2 || len(post.Title) > 25 {
		errs["Title"] = "The Title field must be between 2-25 chars!"
	}
	if len(post.Body) < 2 || len(post.Body) > 25 {
		errs["name3"] = "The Body field must be between 2-25 chars!"
	}

	return errs
}
