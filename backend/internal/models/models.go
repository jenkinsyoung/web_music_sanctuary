package models

import "time"

type Category struct {
	Id   int64  `json:"category_id"`
	Name string `json:"category_name"`
}

type Microcategory struct {
	Id             int64 `json:"microcategory_id"`
	ParentCategory int64
	//TODO: здесь должна быть data json
}

type Advertisement struct {
	Id              int64  `json:"advertisement_id"`
	UserId          int64  `json:"user_id"`
	MicrocategoryId int64  `json:"microcategory_id"`
	Description     string `json:"description"`
	Name            string `json:"advertisement_name"`
	Cost            int64  `json:"cost"`
	//TODO: здесь должны быть images json

}

type Token struct {
	UserId    int64     `json:"user_id"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Id       int64  `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
