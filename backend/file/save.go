package file

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"strings"

	"github.com/google/uuid"
)

func SaveImage(base string) (string, error) {
	base = strings.Replace(base,"data:image/png;base64,", "", -1)
	//fmt.Println(base)
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base))
	img, format, err := image.Decode(reader)
	if err != nil {
		return "", err
	}

	filename := uuid.New().String() + "." + format

	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}

	defer f.Close()

	switch format {
	case "jpeg":
		fallthrough
	case "jpg":
		jpeg.Encode(f, img, nil)
	case "png":
		png.Encode(f, img)
	default:
		fmt.Println(format)
	}

	return filename, nil
}

func GetFile(link string) ([]byte, error) {
	adr := os.Getenv("HOST_ADR")
	filename := strings.Replace(link, adr, "", -1)
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return raw, nil
}