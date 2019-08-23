package upload

import (
	"github.com/julienschmidt/httprouter"
	"github.com/markusleevip/taostorage/config"
	"github.com/markusleevip/taostorage/db/model"
	"github.com/markusleevip/taostorage/log"
	"github.com/markusleevip/taostorage/utils"
	"github.com/markusleevip/taostorage/web/common"
	"github.com/markusleevip/taostorage/web/core/kit"
	"github.com/markusleevip/taostorage/web/core/render"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const baseFormat = "2006-01-02 15:04:05"

var albumPath = ""

type Controller struct {
}

func (Controller) Upload(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	albumPath = config.PFile.AlbumPath + "/"
	file, fHead, err := r.FormFile("uploadFile") //
	log.Info("albumPath:%s", albumPath)
	// 读文件错误
	if err != nil {
		common.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error.")
		return
	}
	defer file.Close()

	// 获取文件的扩展名
	extName := utils.GetFileExt(fHead.Filename)

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error("Read file error : %v", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error.")
		return
	}

	// 生成文件sha256码
	sha256Value := utils.GetByteSha256(data)
	log.Info(sha256Value)

	// 获取DB中是否已经保存该文件
	temp := model.Resource{}
	temp.NameSha256 = sha256Value
	temp.Get()
	if temp.FileName != "" {
		log.Error(" 文件已经存在，文件名=%s\n", temp.FileName)
	}

	fileName := fHead.Filename

	tempFile := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	tempFileName := tempFile + "." + extName
	previewFileName := tempFile + "_" + utils.PhotoPreviewSizeStr + "." + extName
	tempStoreFile := albumPath + tempFile + "." + extName
	log.Info("storeFile:%s", tempStoreFile)
	err = ioutil.WriteFile(tempStoreFile, data, 0664)

	if err != nil {
		log.Error("Write file error: %v", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error.")
		return
	}
	fileInfo, err := os.Stat(tempStoreFile)
	if err != nil {
		log.Error(" get fileInfo error: %v", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error.")
		return
	}
	res := model.Resource{}
	cTime, err := utils.Photo{}.GetDate(tempStoreFile)
	res.FileName = tempFileName
	res.FileSize = fileInfo.Size()
	res.NameSha256 = sha256Value
	res.FileType = utils.GetFileType(extName)
	if err == nil {
		log.Info("ctime:%d", cTime.Unix())
		res.CTime = strconv.FormatInt(cTime.Unix(), 10)
		res.FilePath = utils.GetDateYYYYMM(cTime)
		netFileName := albumPath + res.FilePath + "/" + tempFileName
		_, err := os.Stat(albumPath + res.FilePath)
		if err != nil {
			log.Error("Read dir error:%v", err)
			err = os.Mkdir(albumPath+res.FilePath,0765)
			//err = os.MkdirAll(albumPath+res.FilePath, 0664)
			if err != nil {
				log.Error("Mkdir error:%v", err)
			} else {
				log.Info("Mkdir success:%s", albumPath+res.FilePath)
			}
		}
		err = os.Rename(tempStoreFile, netFileName)
		if err != nil {
			log.Error("rename error:%v", err)
		}
		// previewFileName
		err = utils.Photo{}.CreatePreviewImg(netFileName, albumPath+res.FilePath+"/"+previewFileName)
		if err != nil {
			log.Error("rename error:%v", err)
		} else {
			res.Preview = previewFileName
		}

	}
	log.Info("save res.")
	// save to taodb
	res.Save()
	w.WriteHeader(http.StatusOK)
	ret := kit.GetCommonRet()
	ret.State = kit.RetStateOk
	log.Info("===fileName:%s", fileName)
	bean := Bean{}
	bean.FileName = fileName
	bean.State = 1
	ret.Data = bean
	render.RenderJson(w, ret)

}
