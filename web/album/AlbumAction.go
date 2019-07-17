package album

import (
	mux "github.com/julienschmidt/httprouter"
	"github.com/markusleevip/taostorage/config"
	"github.com/markusleevip/taostorage/db"
	"github.com/markusleevip/taostorage/web/common"
	"github.com/markusleevip/taostorage/web/core/kit"
	"github.com/markusleevip/taostorage/web/core/render"
	"log"
	"net/http"
	"os"
	"time"
)

func List(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	db := db.GetDb()
	dataMap, _ := db.Iterator("album")
	ret := kit.GetCommonRet()
	if dataMap != nil{
		list ,_:= mapsToList(dataMap)
		ret.Data = list
		ret.State =  kit.RetStateOk
	}
	render.RenderJson(w, ret)
}

func Show(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	fileName := ps.ByName("fileName")
	if fileName !=""{
		albumPath := config.PFile.AlbumPath+"/"
		data, err := os.Open(albumPath+fileName)
		if err !=nil{
			log.Printf("Read file error : %v", err)
			common.SendErrorResponse(w, http.StatusNotFound, "Not found file.")
			return
		}
		if data!=nil{
			w.Header().Set("Content-Type", "image/jpeg")
			http.ServeContent(w, r, "", time.Now().UTC(), data)
		}
	}

}
