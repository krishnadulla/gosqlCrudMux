package handlerfunctions

import (
	"fmt"
	"net/http"

	
	"github.com/krishnadulla/gosqlCrudMux/models"
	"github.com/krishnadulla/sqlCrud/config"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
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
		

		rows, err := userModel.DeleteUser(10)

		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("Done", rows)
		}

	}

}
