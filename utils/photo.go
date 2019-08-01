package utils

import (
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"os"
	"time"
)

type Photo struct{}

func (Photo) GetDate(fileName string) (time.Time, error){
	f, err := os.Open(fileName)
	var dt time.Time
	if err != nil {
		return dt ,err
	}
	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		return dt ,err
	}
	f.Close()
	return  x.DateTime()

}
