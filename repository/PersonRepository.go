package repository

import (
	"database/sql"
	"simpleDDDServer/model"
)

type UserRepository struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{database: database}
}

func (r *UserRepository) FindAll() [] *model.User {

	rows, _ := r.database.Query(`SELECT id, name, age FROM users;`)
	defer rows.Close()

	users := []*model.User{}

	for rows.Next() {
		var (
			id int
			name string
			age int
		)

		rows.Scan(&id, &name, &age)

		users = append(users, &model.User{
			Id:   id,
			Name: name,
			Age:  age,
		})

	}

	return users

}