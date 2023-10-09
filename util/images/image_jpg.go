package images

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"path"
	"path/filepath"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

type JpgImageData struct {
	UUID           uuid.UUID
	HomePath       string
	SourceFile     FileSpecs
	ThumbImageFile FileSpecs
	NewFile        FileSpecs
	ThumbImage     image.Image
}

func NewJpg(fileSpecs FileSpecs, homePath string, uuid uuid.UUID) *JpgImageData {
	// log.Printf("NewJpg: %v---> %v", fileSpecs.FullPath(), path.Join(homePath, TargetPath, uuid.String()+".jpg"))
	CopyFile(fileSpecs.FullPath(), path.Join(homePath, TargetPath, uuid.String()+".jpg"))
	return &JpgImageData{
		UUID:       uuid,
		HomePath:   homePath,
		SourceFile: fileSpecs,
		ThumbImageFile: FileSpecs{
			Path:      TargetPath,
			BaseName:  uuid.String() + "_thumb",
			Extension: ".jpg",
		},
		NewFile: FileSpecs{
			Path:      TargetPath,
			BaseName:  uuid.String(),
			Extension: ".jpg",
		},
	}
}

func NewJpgFromURL(fileSpecs FileSpecs, homePath string, uuid uuid.UUID) *JpgImageData {
	return &JpgImageData{
		UUID:       uuid,
		HomePath:   homePath,
		SourceFile: fileSpecs,
		ThumbImageFile: FileSpecs{
			Path:      TargetPath,
			BaseName:  uuid.String() + "_thumb",
			Extension: ".jpg",
		},
		NewFile: fileSpecs,
	}
}

func (t *JpgImageData) Thumbnail() image.Image {
	return t.ThumbImage
}

func (t *JpgImageData) ThumbnailBytes() []byte {
	return util.ImageToBytes(t.ThumbImage)
}

func (t *JpgImageData) ImageData() ImageData {
	return ImageData{
		UUID:           t.UUID,
		SourceFile:     t.SourceFile,
		ThumbImageFile: t.ThumbImageFile,
		NewFile:        t.NewFile,
		ThumbImage:     t.ThumbImage,
	}
}

func (t *JpgImageData) CreateThumbnail() error {

	log.Printf("CreateThumbnail 1: %v", t.SourceFile.FullPath())
	originalimagefile, err := os.Open(t.SourceFile.FullPath())
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("CreateThumbnail: 2")
	img, err := jpeg.Decode(originalimagefile)
	if err != nil {
		log.Println("Encountered Error while decoding image file: ", err)
		return err
	}

	log.Println("CreateThumbnail: 3")
	thumbImage := resize.Resize(270, 0, img, resize.Lanczos3)
	t.ThumbImage = thumbImage

	log.Printf("filepath.Join(t.HomePath: %v, t.ThumbImageFile.FullPath()): %v",
		t.HomePath, t.ThumbImageFile.FullPath())
	thumbImageFile, err := os.Create(filepath.Join(t.HomePath, t.ThumbImageFile.FullPath()))
	log.Printf("4: %v", filepath.Join(t.HomePath, t.ThumbImageFile.FullPath()))

	if err != nil {
		log.Println("Encountered error while resizing image:", err)
		return err
	}

	err = jpeg.Encode(thumbImageFile, thumbImage, &jpeg.Options{Quality: 75})

	if err != nil {
		log.Println("Encountered error while resizing image:", err)
		return err
	}

	return nil
}
