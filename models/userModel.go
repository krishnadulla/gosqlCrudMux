package models

import (
	"database/sql"

	"github.com/krishnadulla/gosqlCrudMux/entities"
)

type UserModel struct {
	Db *sql.DB
}

func (userModel UserModel) FindAll() ([]entities.User, error) {
	rows, err := userModel.Db.Query("SELECT * FROM user_details")

	if err != nil {
		return nil, err
	} else {
		users := []entities.User{}

		for rows.Next() {
			var id int64
			var firstName string
			var lastName string
			var phoneNum string
			var email string
			var password string
			var confirmPassword string
			var createDate string
			var createBy string
			err2 := rows.Scan(&id, &firstName, &lastName, &phoneNum, &email, &password, &confirmPassword, &createDate, &createBy)
			if err2 != nil {
				return nil, err2
			} else {
				user := entities.User{
					ID:              id,
					FirstName:       firstName,
					LastName:        lastName,
					PhoneNum:        phoneNum,
					Email:           email,
					Password:        password,
					ConfirmPassword: confirmPassword,
					CreateDate:      createDate,
					CreateBy:        createBy}
				users = append(users, user)
			}
		}
		return users, nil
	}

}

func (userModel UserModel) FindUserById(id int64) (entities.User, error) {
	rows, err := userModel.Db.Query("SELECT * FROM user_details WHERE ID = ?", id)

	if err != nil {
		return entities.User{}, err
	} else {
		users := entities.User{}

		for rows.Next() {
			var id int64
			var firstName string
			var lastName string
			var phoneNum string
			var email string
			var password string
			var confirmPassword string
			var createDate string
			var createBy string
			err2 := rows.Scan(&id, &firstName, &lastName, &phoneNum, &email, &password, &confirmPassword, &createDate, &createBy)
			if err2 != nil {
				return entities.User{}, err2
			} else {
				userSel := entities.User{
					ID:              id,
					FirstName:       firstName,
					LastName:        lastName,
					PhoneNum:        phoneNum,
					Email:           email,
					Password:        password,
					ConfirmPassword: confirmPassword,
					CreateDate:      createDate,
					CreateBy:        createBy}
				users = userSel
			}
		}
		return users, nil
	}

}

func (userModel UserModel) CreateUserModel(user *entities.User) error {
	result, err := userModel.Db.Exec(
		"INSERT INTO user_details (FIRST_NAME, LAST_NAME,PHONE_NUM, EMAIL,PASSWORD, CONFIRM_PASSWORD, CREATE_DATE, CREATE_BY) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)",
		user.FirstName, user.LastName, user.PhoneNum, user.Email, user.Password, user.ConfirmPassword, user.CreateDate, user.CreateBy)

	if err != nil {
		return err
	} else {
		user.ID, _ = result.LastInsertId()
		return nil
	}
}



func (userModel UserModel) UpdateUser(user entities.User) (int64,error)  {
	result, err := userModel.Db.Exec(
		"UPDATE user_details SET FIRST_NAME =? WHERE ID=?",user.FirstName, user.ID)

	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}



func (userModel UserModel) DeleteUser(id int64) (int64,error)  {
	result, err := userModel.Db.Exec(
		"DELETE FROM  user_details WHERE ID=?",id)

	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}



