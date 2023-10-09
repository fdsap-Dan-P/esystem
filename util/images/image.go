package images

import (
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"simplebank/util/file"
	"strings"

	"github.com/google/uuid"
)

var (
	TargetPath string
)

type ImageData struct {
	UUID           uuid.UUID
	HomePath       string
	SourceFile     FileSpecs
	ThumbImageFile FileSpecs
	NewFile        FileSpecs
	ThumbImage     image.Image
}

// const ImagePath = "app/images/"

// type DocsSpecs interface {
// 	FullPath() string
// }

type FileSpecs struct {
	Path      string
	Extension string
	BaseName  string
}

func (f FileSpecs) FullPath() string {
	return filepath.Join(f.Path, f.BaseName+f.Extension)
}

func (f FileSpecs) FileName() string {
	// log.Printf("FullFileName: %v", f.BaseName+f.Extension)
	return f.BaseName + f.Extension
}

type ImageDocument interface {
	CreateThumbnail() error
	ThumbnailBytes() []byte
	Thumbnail() image.Image
	ImageData() ImageData
	// FullPath() string
	// PrepareNewPath(p string) error
}

func NewImageDocument(
	imagePath string, homePath string, uuid uuid.UUID, targetPath string) ImageDocument {
	// uuid := uuid.New()
	TargetPath = targetPath
	p, n, e := file.FileSpecs(imagePath)

	fs := FileSpecs{
		Path:      p,
		BaseName:  n,
		Extension: e,
	}

	pth := path.Join(homePath, targetPath)
	err := os.MkdirAll(pth, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	var img ImageDocument
	switch e {
	case ".jpeg", ".jpg":
		img = NewJpg(fs, homePath, uuid)

	case ".png":
		img = NewPng(fs, homePath, uuid)

	case ".gif":
		img = NewGif(fs, homePath, uuid)
	}
	return img
}

func extractFileName(url string) string {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	return fileName
}

func NewImageDocumentFromURL(
	url string, homePath string, uuid uuid.UUID, targetPath string) ImageDocument {
	TargetPath = targetPath

	// Fetch the image data
	imageResponse, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching image:", err)
	}
	defer imageResponse.Body.Close()

	// Create the output file
	fileName := extractFileName(url)

	e := filepath.Ext(fileName)

	p := path.Join(homePath, targetPath)
	err = os.MkdirAll(p, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	fullfilePath := fmt.Sprintf("%s%s", path.Join(p, uuid.String()), e)

	fs := FileSpecs{
		Path:      path.Join(homePath, targetPath),
		BaseName:  uuid.String(),
		Extension: e,
	}

	file, err := os.Create(fullfilePath)
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	// Copy the image data to the file
	_, err = io.Copy(file, imageResponse.Body)
	if err != nil {
		log.Fatal("Error saving image:", err)
	}

	var img ImageDocument
	switch e {
	case ".jpeg", ".jpg":
		img = NewJpgFromURL(fs, homePath, uuid)

	case ".png":
		img = NewJpgFromURL(fs, homePath, uuid)

	case ".gif":
		img = NewJpgFromURL(fs, homePath, uuid)
	}
	return img
}

func CopyFile(sourceFilePath string, destinationFilePath string) {

	// Open the source file
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// Create the destination file path
	// destinationFilePath := filepath.Join(destinationFolderPath, filepath.Base(sourceFilePath))

	// Create the destination file
	destinationFile, err := os.Create(destinationFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer destinationFile.Close()

	// Copy the contents from the source file to the destination file
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File copied successfully!")
}
