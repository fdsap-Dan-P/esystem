package images

import (
	"image"
	"image/gif"
	"log"
	"os"
	"path"
	"path/filepath"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

type GifImageData struct {
	UUID           uuid.UUID
	HomePath       string
	SourceFile     FileSpecs
	ThumbImageFile FileSpecs
	NewFile        FileSpecs
	ThumbImage     image.Image
}

func NewGif(fileSpecs FileSpecs, homePath string, uuid uuid.UUID) *GifImageData {
	// id := uuid.New().String()
	CopyFile(fileSpecs.FullPath(), path.Join(homePath, TargetPath, uuid.String()+".Gif"))
	return &GifImageData{
		UUID:       uuid,
		HomePath:   homePath,
		SourceFile: fileSpecs,
		ThumbImageFile: FileSpecs{
			Path:      TargetPath,
			BaseName:  uuid.String() + "_thumb",
			Extension: ".gif",
		},
		NewFile: FileSpecs{
			Path:      TargetPath,
			BaseName:  uuid.String(),
			Extension: ".gif",
		},
	}
}

func NewGifFromURL(fileSpecs FileSpecs, homePath string, uuid uuid.UUID) *GifImageData {
	return &GifImageData{
		UUID:       uuid,
		HomePath:   homePath,
		SourceFile: fileSpecs,
		ThumbImageFile: FileSpecs{
			Path:      TargetPath,
			BaseName:  uuid.String() + "_thumb",
			Extension: ".gif",
		},
		NewFile: fileSpecs,
	}
}

func (t *GifImageData) Thumbnail() image.Image {
	return t.ThumbImage
}

func (t *GifImageData) ThumbnailBytes() []byte {
	return util.ImageToBytes(t.ThumbImage)
}

func (t *GifImageData) ImageData() ImageData {
	return ImageData{
		UUID:           t.UUID,
		SourceFile:     t.SourceFile,
		ThumbImageFile: t.ThumbImageFile,
		NewFile:        t.NewFile,
		ThumbImage:     t.ThumbImage,
	}
}

func (t *GifImageData) CreateThumbnail() error {

	originalimagefile, err := os.Open(t.SourceFile.FullPath())

	if err != nil {
		log.Println(err)
		return err
	}

	img, err := gif.Decode(originalimagefile)

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

	op := gif.Options{NumColors: 256, Quantizer: nil, Drawer: nil}
	gif.Encode(thumbImageFile, thumbImage, &op)

	if err != nil {
		log.Println("Encountered error while resizing image:", err)
		return err
	}

	return nil
}
