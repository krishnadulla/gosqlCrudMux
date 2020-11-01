package handlerfunctions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/krishnadulla/gosqlCrudMux/models"
	"github.com/krishnadulla/sqlCrud/config"
)

func FetchUserByParams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.GetMySQLDB()

	if err != nil {
		fmt.Println(err)
	} else {
		userModelData := models.UserModel{
			Db: db,
		}
		users, err := userModelData.FindUserById(2)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(users)
		}
	}

}
