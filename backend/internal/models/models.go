package models

type Listing struct {
	Id          int64  `json:"advertisement_id"`
	UserId      int64  `json:"user_id"`
	GuitarId    int64  `json:"guitar_id"`
	GuitarName  string `json:"guitar_name"`
	Cost        int64  `json:"cost"`
	Description string `json:"description"`

	ImgList []ImgJSON `json:"img_list"`
}

type ListingFullInfo struct {
	Id     int64 `json:"advertisement_id"`
	UserId int64 `json:"user_id"`

	GuitarId     int64  `json:"guitar_id"`
	Form         string `json:"guitar_form"`
	PickupConfig string `json:"pickup_config"`
	Category     string `json:"category"`

	GuitarName  string `json:"guitar_name"`
	Cost        int64  `json:"cost"`
	Description string `json:"description"`

	ImgList []ImgJSON `json:"img_list"`
}

type User struct {
	Id         int64  `json:"user_id"`
	Name       string `json:"user_name"`
	Surname    string `json:"user_surname"`
	Patronymic string `json:"user_patronymic"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
}

type Picture struct {
	Id    int64  `json:"id"`
	Image []byte `json:"image"`
}

type ImgJSON struct {
	Image string `json:"image"`
}

type Guitar struct {
	Id           int64  `json:"guitar_id"`
	Form         string `json:"guitar_form"`
	PickupConfig string `json:"pickup_config"`
	Category     string `json:"category"`
}
