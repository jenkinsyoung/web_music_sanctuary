package models

type Category struct {
	Id   int64  `json:"category_id"`
	Name string `json:"category_name"`
}

type Microcategory struct {
	Id int64 `json:"microcategory_id"`

	//TODO: здесь должна быть data json
}

type Advertisement struct {
	Id              int64   `json:"advertisement_id"`
	UserId          int64   `json:"user_id"`
	Description     string  `json:"description"`
	Name            string  `json:"advertisement_name"`
	Cost            int64   `json:"cost"`
	TypeId          int64   `json:"type_id"`
	Microcategories []int64 `json:"microcategories"`

	//TODO: здесь должны быть images json

}

type User struct {
	Id       int64  `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type Photo struct {
	Id    string `json:"id"`
	Photo string `json:"photo"`
}
