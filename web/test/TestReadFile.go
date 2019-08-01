package main

import (
	"fmt"
	"github.com/markusleevip/taostorage/utils"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"log"
	"os"
	"strconv"
)

func main() {
	albumPath := "/data/album/"
	fileName := albumPath+"1564471473020780500.jpg"
	fmt.Println(fileName)
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}


	// Two convenience functions exist for date/time taken and GPS coords:
	tm, _ := x.DateTime()
	fmt.Println("Taken: ", tm)
	fmt.Println(strconv.FormatInt(tm.Unix(),10))

	fmt.Println("year=",utils.GetDateYYYYMM(tm))
}
