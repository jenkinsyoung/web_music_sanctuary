package imgMethods

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"os"
)

func SaveImageBase64(encodedString string) error {
	decoded, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%x.jpeg", md5.Sum(decoded))
	err = os.WriteFile(filename, decoded, 0644)
	if err != nil {
		return err
	}
	return nil
	//TODO: Написать получение айдишника и создание папки с этим айдишником объявления
}
