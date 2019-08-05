package main

import (
	"github.com/disintegration/imaging"
	"image/jpeg"
	"os"
)

func main() {
	albumPath := "/data/album/"
	fileName := albumPath+"1564471473020780500.jpg"

	fileName300 := albumPath+"1564471473020780500_512.jpg"
	f, _ := os.Open(fileName)
	srcImage,_ :=jpeg.Decode(f)
	dstImage128 := imaging.Resize(srcImage, 300, 0, imaging.Lanczos)
	file128 ,_:= os.Create(fileName300)
	jpeg.Encode(file128, dstImage128, &jpeg.Options{50})

}
