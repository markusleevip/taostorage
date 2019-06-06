package utils

import "strings"

var (
	FileTypeMap = make(map[string] string)
)

func init(){
	FileTypeMap["mp4"] = "video/mp4"
	FileTypeMap["doc"] = "application/msword"
	FileTypeMap["docx"] = "application/msword"
	FileTypeMap["pdf"] = "application/pdf"
	FileTypeMap["ppt"] = "application/x-ppt"
	FileTypeMap["pptx"] = "application/x-ppt"
	FileTypeMap["png"] = "image/png"
	FileTypeMap["jpg"] = "image/jpeg"
	FileTypeMap["jpeg"] = "image/jpeg"
}

// 根据文件名获取扩展名
func  GetFileExt(fileName string ) string{
	fileName = strings.ToLower(fileName)
	index :=strings.LastIndex(fileName,".")
	return fileName[index+1:]
}

// 根据文件扩展名获取文件类型
func  GetFileType(ext string) string {
	fileType :=FileTypeMap[ext]
	if fileType==""{
		fileType = "application/"+ext
	}
	return fileType
}

