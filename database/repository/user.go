package repository

import (
	"database/sql"
	"log"

	"gihub.com/team3_qgame/model"

	"github.com/google/uuid"
)

const (
	getOneItem  = "SELECT id, name FROM users WHERE id = $1;"
	addOneItem  = "INSERT INTO users (id, name) VALUES ($1, $2)"
	updateItem  = "UPDATE users SET name=$2 WHERE id=$1;"
	deleteItem  = "DELETE FROM users WHERE id=$1;"
	getAllItems = "SELECT * FROM users;"
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
		return err
	}

	rowAff, _ := result.RowsAffected()
	log.Printf("Affected %d rows\n", rowAff)

	return nil
}

//GetUser sends a query for get certain user from DB
func (p *UserRepository) GetUserByID(id uuid.UUID) (model.User, error) {
	user := model.User{}
	row := p.conn.QueryRow(getOneItem, id)

	err := row.Scan(&user.ID, &user.Name)
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
func (p *UserRepository) DeleteUserByID(id uuid.UUID) error {
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
		err := rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
