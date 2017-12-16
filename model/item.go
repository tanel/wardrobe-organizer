package model

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/juju/errors"
)

// Item represents a wardrobe item
type Item struct {
	Base
	UserID      string
	Name        string
	Description string
	Color       string
	Size        string
	Brand       string
	Price       float64
	Currency    string
	Category    string
	Season      string
	Formal      bool

	Images []ItemImage
}

func NewItemForm(r *http.Request) (*Item, error) {
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		return nil, errors.Annotate(err, "parsing multipart form failed")
	}

	var item Item

	item.Name = strings.TrimSpace(r.FormValue("name"))
	if item.Name == "" {
		return nil, errors.New("please enter a name")
	}

	s := strings.TrimSpace(r.FormValue("price"))
	if s != "" {
		price, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, errors.New("please enter a valid price or leave it blank")
		}

		item.Price = price
	}

	item.Description = strings.TrimSpace(r.FormValue("description"))
	item.Color = strings.TrimSpace(r.FormValue("color"))
	item.Size = strings.TrimSpace(r.FormValue("size"))
	item.Brand = strings.TrimSpace(r.FormValue("brand"))
	item.Category = strings.TrimSpace(r.FormValue("category"))
	item.Currency = strings.TrimSpace(r.FormValue("currency"))
	item.Season = strings.TrimSpace(r.FormValue("season"))
	item.CreatedAt = time.Now()

	file, _, err := r.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return nil, errors.Annotate(err, "getting form file failed")
	}

	if file != nil {
		defer file.Close()

		b, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, errors.Annotate(err, "reading form file failed")
		}

		item.Images = append(item.Images, ItemImage{
			Body: b,
		})
	}

	return &item, nil
}

func (itemImage ItemImage) DirectoryPath() string {
	return filepath.Join("uploads", itemImage.UserID, "images")
}

func (itemImage ItemImage) FilePath() string {
	directoryPath := itemImage.DirectoryPath()
	return filepath.Join(directoryPath, itemImage.ID)
}

func (itemImage ItemImage) Save() error {
	directoryPath := itemImage.DirectoryPath()
	if err := os.MkdirAll(directoryPath, 0777); err != nil && !strings.Contains(err.Error(), "file exists") {
		return errors.Annotate(err, "creating image directory failed")
	}

	filePath := itemImage.FilePath()
	if err := ioutil.WriteFile(filePath, itemImage.Body, 0644); err != nil {
		return errors.Annotate(err, "writing image failed")
	}

	return nil
}

func (itemImage *ItemImage) Load() error {
	filePath := itemImage.FilePath()
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.Annotate(err, "writing image failed")
	}

	itemImage.Body = b

	return nil
}
