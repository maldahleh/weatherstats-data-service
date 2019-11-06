package utils

import (
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func DownloadFile(path string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func DeleteFile(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Error("failed to delete file", path, "error", err)
	}
}
