package repositories

import (
	"assyifa/models"
	"database/sql"
)

type UserRepoImpl struct {
	DB *sql.DB
}

func (user UserRepoImpl) GetData() ([]*models.Users, error) {
	dataUsers := []*models.Users{}
	query := "select * from users"

	data, err := user.DB.Query(query)

	if err != nil {
		return nil, err
	}

	for data.Next() {
		users := models.Users{}
		var err = data.Scan(&users.ID, &users.Username, &users.Password, &users.NamaLengkap, &users.TanggalLahir, &users.Alamat, &users.Photo)
		if err != nil {
			return nil, err
		}
		dataUsers = append(dataUsers, &users)
		println("ini user", dataUsers)
	}
	return dataUsers, nil
}

func (user *UserRepoImpl) AddUser(Users models.Users) error {
	tx, err := user.DB.Begin()
	if err != nil {
		return err
	}
	query := "INSERT INTO users (Username,Password,NamaLengkap,TanggalLahir,Alamat,Photo) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&Users.Username, &Users.Password, &Users.NamaLengkap, &Users.TanggalLahir, &Users.Alamat, &Users.Photo)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func InitUserRepoImpl(DB *sql.DB) UserRepositories {
	return &UserRepoImpl{DB}
}
