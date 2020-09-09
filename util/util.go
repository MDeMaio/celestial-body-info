package util

import (
	"bufio"
	"encoding/base64"
	"log"
	"os"
)

// EncodeImgToBase64 will take a file path and convert the image into a base64 string.
func EncodeImgToBase64(path string) (string, error) {
	imgFile, err := os.Open(path)

	if err != nil {
		log.Fatalf("An error occured when opening the file: %v", err)
		return "", err
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, err := imgFile.Stat()
	if err != nil {
		log.Fatalf("An error occured when fetching file info: %v", err)
		return "", err
	}

	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	return imgBase64Str, nil
}
