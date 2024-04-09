package database

import (
	"log"
)

// const Connection, err := ConnectToDB()

func (c *DBConnection) ImgInsert(img []byte) int64 {
	var id int64
	err := c.db.QueryRow(`INSERT INTO "picture" (image) VALUES ($1) RETURNING id`, img).Scan(&id)
	if err != nil {
		log.Printf("Error inserting img to db %s", err)
	}
	return id
}

//func (c *DBConnection) GetGuitarTypeID(name string)(int64, error){
//	var typeID int64
//	err := c.db.QueryRow(`SELECT id FROM "type" WHERE name=$1`, name).Scan(&typeID)
//	if err != nil{
//		log.Printf("error could not find type of guitar %s", err)
//		return 0, err
//	}
//	return typeID, nil
//}
