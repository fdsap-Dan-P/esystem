package images

import (
	"image"
	"image/png"
	"log"
	"os"
	"path"
	"path/filepath"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

type PngImageData struct {
	UUID           uuid.UUID
	HomePath       string
	SourceFile     FileSpecs
	ThumbImageFile FileSpecs
	NewFile        FileSpecs
	ThumbImage     image.Image
}

func NewPng(fileSpecs FileSpecs, homePath string, uuid uuid.UUID) *PngImageData {
	CopyFile(fileSpecs.FullPath(), path.Join(homePath, TargetPath, uuid.String()+".png"))
	// id := uuid.New().String()
	return &PngImageData{
		UUID:       uuid,
		HomePath:   homePath,
		SourceFile: fileSpecs,
		ThumbImageFile: FileSpecs{
			Path:      TargetPath,
			BaseName:  uuid.String() + "_thumb",
			Extension: ".png",
		},
		NewFile: FileSpecs{
			Path:      TargetPath,
			BaseName:  uuid.String(),
			Extension: ".png",
		},
	}
}

func NewPngFromURL(fileSpecs FileSpecs, homePath string, uuid uuid.UUID) *PngImageData {
	return &PngImageData{
		UUID:       uuid,
		HomePath:   homePath,
		SourceFile: fileSpecs,
		ThumbImageFile: FileSpecs{
			Path:      TargetPath,
			BaseName:  uuid.String() + "_thumb",
			Extension: ".png",
		},
		NewFile: fileSpecs,
	}
}

func (t *PngImageData) Thumbnail() image.Image {
	return t.ThumbImage
}

func (t *PngImageData) ThumbnailBytes() []byte {
	return util.ImageToBytes(t.ThumbImage)
}

func (t *PngImageData) ImageData() ImageData {
	return ImageData{
		UUID:           t.UUID,
		SourceFile:     t.SourceFile,
		ThumbImageFile: t.ThumbImageFile,
		NewFile:        t.NewFile,
		ThumbImage:     t.ThumbImage,
	}
}

func (t *PngImageData) CreateThumbnail() error {

	originalimagefile, err := os.Open(t.SourceFile.FullPath())

	if err != nil {
		log.Println(err)
		return err
	}

	img, err := png.Decode(originalimagefile)
	if err != nil {
		log.Println("Encountered Error while decoding image file: ", err)
		return err
	}

	thumbImage := resize.Resize(270, 0, img, resize.Lanczos3)
	t.ThumbImage = thumbImage
	thumbImageFile, err := os.Create(filepath.Join(t.HomePath, t.ThumbImageFile.FullPath()))

	if err != nil {
		log.Println("Encountered error while resizing image:", err)
		return err
	}

	png.Encode(thumbImageFile, thumbImage)

	if err != nil {
		log.Println("Encountered error while resizing image:", err)
		return err
	}
	return nil
}
