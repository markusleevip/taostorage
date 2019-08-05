package utils

import (
	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"image/jpeg"
	"os"
	"time"
)

const (
	PhotoPreviewSize    =300
	PhotoPreviewSizeStr ="300"
)

type Photo struct{}

func (Photo) GetDate(fileName string) (time.Time, error) {
	f, err := os.Open(fileName)
	var dt time.Time
	if err != nil {
		return dt, err
	}
	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		return dt, err
	}
	f.Close()
	return x.DateTime()

}

func (Photo) CreatePreviewImg(imgName string ,prevFileName string) error {
	f, _ := os.Open(imgName)
	srcImage, _ := jpeg.Decode(f)
	dstImage := imaging.Resize(srcImage, PhotoPreviewSize, 0, imaging.Lanczos)
	prevName, _ := os.Create(prevFileName)
	return jpeg.Encode(prevName, dstImage, &jpeg.Options{50})
}
