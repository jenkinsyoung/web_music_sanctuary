package database

import (
	"fmt"
)

//const Connection, err := ConnectToDB()

func (c *DBConnection) GetUserID(email, hashedPassword string) (int, error) {
	id := 0
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND pass_hash=$2")
	err := c.db.QueryRow(query, email, hashedPassword).Scan(&id)
	return id, err
}

func (c *DBConnection) CreateUser(email, password string) {

}
