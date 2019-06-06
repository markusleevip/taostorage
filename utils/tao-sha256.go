package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func GetFileSha256(filePath string ) (string, error){
	result :=""
	fileInfo, err := os.Stat(filePath)
	if err != nil{
		panic(err)
	}

	if !fileInfo.IsDir() {
		fmt.Println(fileInfo.Name())
		fmt.Println(fileInfo.Size()/1024/1024, "MB")
		fmt.Println(fileInfo.ModTime())
		file, err := os.Open(filePath)
		if err != nil {
			return result, err
		}
		defer file.Close()
		hash := sha256.New()
		if _, err := io.Copy(hash,file); err != nil {
			return result, err
		}

		hashInBytes := hash.Sum(nil)
		result = hex.EncodeToString(hashInBytes)

	}
	return result,err

}

func GetTxtSha256(str string) (string ){
	bs := sha256.Sum256([]byte(str))
	return fmt.Sprintf("%x", bs)
}

func GetByteSha256(bytes []byte) (string ){
	bs := sha256.Sum256(bytes)
	return fmt.Sprintf("%x", bs)
}

