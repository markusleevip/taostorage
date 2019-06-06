package upload

import (
	"github.com/julienschmidt/httprouter"
	"github.com/markusleevip/taostorage/config"
	"github.com/markusleevip/taostorage/db/model"
	"github.com/markusleevip/taostorage/utils"
	"github.com/markusleevip/taostorage/web/common"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const baseFormat = "2006-01-02 15:04:05"

var  albumPath = ""

type Controller struct {

}

func (Controller) Upload(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	albumPath = config.PFile.AlbumPath+"/"
	file, fHead, err := r.FormFile("uploadFile") //
	log.Println("albumPath:",albumPath)
	// 读文件错误
	if err != nil {
		common.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error.")
		return
	}
	// 获取文件的扩展名
	extName := utils.GetFileExt(fHead.Filename)

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error : %v", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error.")
		return
	}

	// 生成文件sha256码
	sha256Value := utils.GetByteSha256(data)
	log.Println(sha256Value)

	fileName := fHead.Filename

	tempFileName := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	storeFile := albumPath + tempFileName+"."+extName
	log.Println("storeFile:",storeFile)
	err = ioutil.WriteFile(storeFile, data, 0644)

	if err != nil {
		log.Printf("Write file error: %v", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error.")
		return
	}
	fileInfo, err := os.Stat(storeFile)
	if err != nil {
		log.Printf(" get fileInfo error: %v", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error.")
		return
	}
	res := model.Resource{}
	res.FileName = tempFileName+"."+extName
	res.FileSize = fileInfo.Size()
	res.NameSha256 = sha256Value
	res.FileType = utils.GetFileType(extName)
	// save to taodb
	res.Save()
	fileSize := fileInfo.Size()

	log.Println("fileName:" + fileName)
	log.Println(fileSize)
	nt := time.Now()
	nowTimeStr := nt.Format(baseFormat)
	log.Println(nowTimeStr)

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded success.")

}

