package handlerfunctions

import (
	"fmt"
	"net/http"

	"github.com/krishnadulla/gosqlCrudMux/entities"
	"github.com/krishnadulla/gosqlCrudMux/models"
	"github.com/krishnadulla/sqlCrud/config"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server Started")
	w.Header().Set("Content-Type", "application/json")
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DB Connection Formed")
		userModel := models.UserModel{
			Db: db,
		}
		user := entities.User{
			ID: 10,
			FirstName:       "Krishna-Updated",
			LastName:        "Dulla-4",
			PhoneNum:        "1234567",
			Email:           "testbasicuser@gmail.com",
			Password:        "testq",
			ConfirmPassword: "test123",
			CreateDate:      "1/10/2020",
			CreateBy:        "Krishna Dulla",
		}

		rows, err := userModel.UpdateUser(user)

		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			if rows >0 {
				fmt.Println("Done")
			}
		}

	}

}
