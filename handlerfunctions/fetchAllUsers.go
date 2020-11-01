package handlerfunctions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/krishnadulla/gosqlCrudMux/models"
	"github.com/krishnadulla/sqlCrud/config"
)

func FetchAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DB Connection Successful")
		userModel := models.UserModel{
			Db: db,
		}
		users, err := userModel.FindAll()

		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(users)
		}
	}

}

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Server Started")
// 	w.Header().Set("Content-Type", "application/json")
// 	db, err := config.GetMySQLDB()

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("DB Connection Formed")
// 		userModel := models.UserModel{
// 			Db: db,
// 		}
// 		user := entities.User{
// 			FirstName:       "Krishna-4",
// 			LastName:        "Dulla-4",
// 			PhoneNum:        "1234567",
// 			Email:           "testbasicuser@gmail.com",
// 			Password:        "testq",
// 			ConfirmPassword: "test123",
// 			CreateDate:      "31/10/2020",
// 			CreateBy:        "Krishna Dulla",
// 		}

// 		err := userModel.CreateUserModel(&user)

// 		if err != nil {
// 			fmt.Println("Error: ", err)
// 		} else {
// 			fmt.Println("Latest Id: ", user.ID)
// 		}

// 	}

// }
