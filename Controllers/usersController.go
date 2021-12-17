package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	data "github.com/pradyumna001/blog-app/Data"
	models "github.com/pradyumna001/blog-app/Models"
)

type UsersController struct{}

var db = *data.DB

//Return Handlers After Authentication
func (uc *UsersController) AuthenticateUser(rw http.ResponseWriter, r *http.Request) bool {
	_, pass, ok := r.BasicAuth()
	storedUser := uc.GetUser(r)
	if !ok {
		rw.Header().Set("WWW-Authenticate", `Basic realm="Enter UserName and Password"`)
		rw.WriteHeader(401)
		rw.Write([]byte("You are Unauthorized to access the application.\n"))
		return false
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	if err != nil {
		log.Println("Password Encryption Failed", err)
	}
	if len([]byte(hashedPass)) != len([]byte(storedUser.Password)) {
		log.Println("Wrong Password")
		return false
	}
	return true
}

//Create New User
func (uc *UsersController) CreateRecord(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	name, nameOk := vals["name"]
	email, emailOk := vals["email"]
	userName, userNameOk := vals["userName"]
	password, passwordOk := vals["password"]
	dateOfBirthstr, dateOfBirthOk := vals["dateOfBirth"]
	phoneNumber, phoneNumberOk := vals["phoneNumber"]
	user := models.User{
		Name:        name[0],
		Email:       email[0],
		Username:    userName[0],
		Password:    password[0],
		DateOfBirth: dateOfBirthstr[0],
		PhoneNumber: phoneNumber[0],
	}

	validationErr := models.IsValidUser(user)
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
	//encrypting password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		log.Println("Password Encryption Failed : ", err)
	}
	if nameOk && emailOk && userNameOk && passwordOk && dateOfBirthOk && phoneNumberOk {
		log.Print("Inserting record in database for Name : ", name[0])
		id := 0
		err := db.QueryRow("INSERT into user_table(uname,email,username,userpassword,dateofbirth,phonenumber) values( $1 , $2 , $3 ,  $4 ,  $5 , $6 ) RETURNING uid", string(name[0]), string(email[0]), string(userName[0]), string(hashedPassword), dateOfBirthstr[0], string(phoneNumber[0])).Scan(&id)
		if err != nil {
			log.Print("error executing query :: ", err)
			fmt.Fprint(w, "error executing query :: ", err)
			return
		}
		fmt.Fprint(w, "Record Successfully Inserted.\nLast Inserted Record is  :: ", id)
	} else {
		fmt.Fprintf(w, "Error occurred while creating record in database for name :: %s", name[0])
	}
}

//Read All Users
func (uc *UsersController) ReadRecords(w http.ResponseWriter, r *http.Request) {

	users := uc.GetUsers()
	usersJson, err := json.Marshal(users)
	if err != nil {
		log.Fatal("Error  :: ", err)
		fmt.Fprint(w, "Error  :: ", err)
	}
	_, err = w.Write(usersJson)
	if err != nil {
		log.Fatal("Error  :: ", err)
		fmt.Fprint(w, "Error  :: ", err)
	}
}

//Read Single User
func (uc *UsersController) ReadRecord(w http.ResponseWriter, r *http.Request) {

	userInstance := uc.GetUser(r)
	userJson, err := json.Marshal(userInstance)
	if err != nil {
		log.Fatal("Error  :: ", err)
		fmt.Fprint(w, "Error  :: ", err)
	}
	_, err = w.Write(userJson)
	if err != nil {
		log.Fatal("Error  :: ", err)
		fmt.Fprint(w, "Error  :: ", err)
	}
}

//Get Single User After Authentication Only
func (uc *UsersController) GetUser(r *http.Request) models.User {
	log.Print("reading records from user database")
	userName, _, _ := r.BasicAuth()

	row, err := db.Query("SELECT * FROM user_table where userName = $1 ", userName)

	if err != nil {
		log.Print("error executing select query :: ", err)

	}
	var newUser models.User
	for row.Next() {
		var id string
		var name string
		var email string
		var userNameDB string
		var password string
		var dateOfBirth string
		var phoneNumber string
		err = row.Scan(&id, &name, &email, &userNameDB, &dateOfBirth, &phoneNumber, &password)
		if err != nil {
			log.Fatal("Error  :: ", err)

		}
		newUser = models.User{Id: id, Name: name, Email: email, Username: userNameDB, Password: password, DateOfBirth: dateOfBirth, PhoneNumber: phoneNumber}
	}

	return newUser
}

//Get All Users
func (uc *UsersController) GetUsers() []models.User {
	log.Print("reading records from user database")
	rows, err := db.Query("SELECT * FROM user_table")

	if err != nil {
		log.Print("error executing select query :: ", err)
		return nil
	}
	users := []models.User{}
	for rows.Next() {

		var id string
		var name string
		var email string
		var userName string
		var password string
		var dateOfBirth string
		var phoneNumber string
		err = rows.Scan(&id, &name, &email, &userName, &dateOfBirth, &phoneNumber, &password)
		if err != nil {
			log.Fatal("Error  :: ", err)

		}
		newuser := models.User{Id: id, Name: name, Email: email, Username: userName, Password: password, DateOfBirth: dateOfBirth, PhoneNumber: phoneNumber}

		users = append(users, newuser)
	}
	return users
}

//Update Record
func (uc *UsersController) UpdateRecord(w http.ResponseWriter, r *http.Request) {

	vals := r.URL.Query()

	id, idOk := vals["id"]
	name, nameOk := vals["name"]
	email, emailOk := vals["email"]
	userName, userNameOk := vals["userName"]
	password, passwordOk := vals["password"]
	dateOfBirthstr, dateOfBirthOk := vals["dateOfBirth"]
	phoneNumber, phoneNumberOk := vals["phoneNumber"]

	user := models.User{
		Name:        name[0],
		Email:       email[0],
		Username:    userName[0],
		Password:    password[0],
		DateOfBirth: dateOfBirthstr[0],
		PhoneNumber: phoneNumber[0],
	}

	validationErr := models.IsValidUser(user)
	//date validation
	layout := "2006-01-02"
	dateOfBirth, err := time.Parse(layout, dateOfBirthstr[0])
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

	if idOk && nameOk && emailOk && userNameOk && passwordOk && dateOfBirthOk && phoneNumberOk {
		log.Print("going to update record in database for id :: ", id[0])
		//encrypting password

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password[0]), 8)
		if err != nil {
			log.Fatal("Error while converting ciphertest to hashed password :: ", err)
			fmt.Fprint(w, "Error  :: ", err)
		}

		result, err := db.Exec("UPDATE user_table SET uName = $1 , email = $2 , userName = $3 , userpassword = $4 , dateOfBirth = $5 , phoneNumber = $6  where uid = $7", name[0], email[0], userName[0], hashedPassword, dateOfBirth, phoneNumber[0], id[0])
		if err != nil {
			log.Print("error occurred while executing query :: ", err)
			return
		}
		rowsAffected, _ := result.RowsAffected()
		fmt.Fprintf(w, "Number of rows updated in database are :: %d", rowsAffected)
	} else {
		fmt.Fprint(w, "Error occurred while updating record in database for id :: ", id[0])
	}
}

//Delete Record By ID
func (uc *UsersController) DeleteRecord(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	id, idok := vals["id"]
	if idok {
		log.Print("going to delete record in database for id :: ", id[0])

		result, err := db.Exec("DELETE from user_table where uid=$1", id[0])
		if err != nil {
			log.Print("error occurred while executing query :: ", err)
			return
		}

		rowsAffected, _ := result.RowsAffected()
		fmt.Println("Number of rows deleted in database are :: ", rowsAffected)
		fmt.Fprint(w, id[0])
	} else {
		fmt.Fprint(w, "Error occurred while deleting record in database for id ", id[0])
	}
}

//Read All Searched Users
func (uc *UsersController) SearchRecords(w http.ResponseWriter, r *http.Request) {

	users := uc.GetSearchedUsers(r)
	usersJson, err := json.Marshal(users)
	if err != nil {
		log.Fatal("Error  :: ", err)
		fmt.Fprint(w, "Error  :: ", err)
	}
	_, err = w.Write(usersJson)
	if err != nil {
		log.Fatal("Error  :: ", err)
		fmt.Fprint(w, "Error  :: ", err)
	}
}

//Get Searched Users
func (uc *UsersController) GetSearchedUsers(r *http.Request) []models.User {

	vals := r.URL.Query()

	id, idOk := vals["id"]
	name, nameOk := vals["name"]
	email, emailOk := vals["email"]
	userName, userNameOk := vals["userName"]

	if idOk {
		log.Print("reading records from user database by ID")
		rows, err := db.Query("SELECT * FROM user_table where uid = $1", id[0])

		if err != nil {
			log.Print("error executing select query :: ", err)
			return nil
		}
		users := []models.User{}
		for rows.Next() {

			var id string
			var name string
			var email string
			var userName string
			var password string
			var dateOfBirth string
			var phoneNumber string
			err = rows.Scan(&id, &name, &email, &userName, &dateOfBirth, &phoneNumber, &password)
			if err != nil {
				log.Fatal("Error  :: ", err)

			}
			newuser := models.User{Id: id, Name: name, Email: email, Username: userName, Password: password, DateOfBirth: dateOfBirth, PhoneNumber: phoneNumber}

			users = append(users, newuser)
		}
		return users
	} else if nameOk {
		log.Print("reading records from user database by Name")
		nameSearch := "%" + name[0] + "%"
		rows, err := db.Query("SELECT * FROM user_table where uName LIKE $1", nameSearch)

		if err != nil {
			log.Print("error executing select query :: ", err)
			return nil
		}
		users := []models.User{}
		for rows.Next() {

			var id string
			var name string
			var email string
			var userName string
			var password string
			var dateOfBirth string
			var phoneNumber string
			err = rows.Scan(&id, &name, &email, &userName, &dateOfBirth, &phoneNumber, &password)
			if err != nil {
				log.Fatal("Error  :: ", err)

			}
			newuser := models.User{Id: id, Name: name, Email: email, Username: userName, Password: password, DateOfBirth: dateOfBirth, PhoneNumber: phoneNumber}

			users = append(users, newuser)
		}
		return users
	} else if emailOk {
		log.Print("reading records from user database by Email")
		emailSearch := "%" + email[0] + "%"
		rows, err := db.Query("SELECT * FROM user_table where email LIKE $1", emailSearch)

		if err != nil {
			log.Print("error executing select query :: ", err)
			return nil
		}
		users := []models.User{}
		for rows.Next() {

			var id string
			var name string
			var email string
			var userName string
			var password string
			var dateOfBirth string
			var phoneNumber string
			err = rows.Scan(&id, &name, &email, &userName, &dateOfBirth, &phoneNumber, &password)
			if err != nil {
				log.Fatal("Error  :: ", err)

			}
			newuser := models.User{Id: id, Name: name, Email: email, Username: userName, Password: password, DateOfBirth: dateOfBirth, PhoneNumber: phoneNumber}

			users = append(users, newuser)
		}
		return users
	} else if userNameOk {
		log.Print("reading records from user database by UserName")
		userNameSearch := "%" + userName[0] + "%"
		rows, err := db.Query("SELECT * FROM user_table where userName LIKE $1", userNameSearch)

		if err != nil {
			log.Print("error executing select query :: ", err)
			return nil
		}
		users := []models.User{}
		for rows.Next() {

			var id string
			var name string
			var email string
			var userName string
			var password string
			var dateOfBirth string
			var phoneNumber string
			err = rows.Scan(&id, &name, &email, &userName, &dateOfBirth, &phoneNumber, &password)
			if err != nil {
				log.Fatal("Error  :: ", err)

			}
			newuser := models.User{Id: id, Name: name, Email: email, Username: userName, Password: password, DateOfBirth: dateOfBirth, PhoneNumber: phoneNumber}

			users = append(users, newuser)
		}
		return users
	}
	return uc.GetUsers()
}
