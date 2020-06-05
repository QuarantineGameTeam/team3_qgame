package repository

import (
	"database/sql"
	"fmt"
	"log"

	"gihub.com/team3_qgame/model"
)

const (
	getOneItem  = `SELECT * FROM public.users WHERE id = $1;`
	addOneItem  = `INSERT INTO public.users (id, name) VALUES ($1, $2);`
	updateItem  = `UPDATE public.users SET name=$2 WHERE id=$1;`
	deleteItem  = `DELETE FROM public.users WHERE id=$1;`
	getAllItems = `SELECT * FROM public.users;`
)

type UserRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

//NewUser sends a query for creating new one ticket
func (p *UserRepository) NewUser(user model.User) error {
	result, err := p.conn.Exec(addOneItem, user.ID, user.Name)
	if err != nil {
		log.Printf("%v\n", user.Name)
		return err
	}

	rowAff, _ := result.RowsAffected()
	log.Printf("Affected %d rows\n", rowAff)

	return nil
}

//GetUser sends a query for get certain user from DB
func (p *UserRepository) GetUserByID(id int64) (model.User, error) {
	var user model.User
	row := p.conn.QueryRow(getOneItem, id)
	fmt.Println("ROw", row)
	err := row.Scan(&user.ID, &user.Name, &user.Team, &user.Role, &user.Health, &user.Strength, &user.Defence, &user.Intellect, &user.Level)
	if err != nil {
		return user, err
	}

	return user, nil
}

//UpdateUser sends a query for updating one User
func (p *UserRepository) UpdateUser(user model.User) error {
	result, err := p.conn.Exec(updateItem, user.ID, user.Name)
	if err != nil {
		return err
	}

	rowAff, _ := result.RowsAffected()
	log.Printf("Affected %d rows\n", rowAff)

	return nil
}

//DeleteUser sends a query for deleting one User by ID
func (p *UserRepository) DeleteUserByID(id int64) error {
	result, err := p.conn.Exec(deleteItem, id)
	if err != nil {
		return err
	}

	rowAff, _ := result.RowsAffected()
	log.Printf("Affected %d rows\n", rowAff)

	return nil
}

// GetAllUsers sends a query for getting all users
func (p *UserRepository) GetAllUsers() ([]model.User, error) {
	rows, err := p.conn.Query(getAllItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]model.User, 0)
	for rows.Next() {
		u := model.User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Team, &u.Role, &u.Health, &u.Strength, &u.Defence, &u.Intellect, &u.Level)
		if err != nil {
			users = append(users, u)
		}

		log.Printf("\n%v\n", u.Name)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}