package imgMethods

import (
	"log"
	"os"
	"path/filepath"
)

func CreatePhotoDir() {
	homeDir, _ := os.Getwd()
	photoDir := filepath.Join(homeDir, "backend", "photo")
	if _, err := os.Stat(photoDir); os.IsNotExist(err) {
		err = os.Mkdir(photoDir, os.ModeDir)
		if err != nil {
			log.Printf("Error occured creating photo dir: %s", err)
		}
	}
}
