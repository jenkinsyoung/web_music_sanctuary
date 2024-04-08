package imgMethods

import (
	"encoding/base64"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/database"
	"log"
)

func ImgDecode(encodedString string, idFlow chan int64) {
	decoded, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Printf("error decode image %s", err)
	}
	idFlow <- database.DB.ImgInsert(decoded)
}
