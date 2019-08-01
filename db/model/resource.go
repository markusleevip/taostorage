package model

import (
	"encoding/json"
	"fmt"
	"github.com/markusleevip/taostorage/db"
)

var PrefixFileName ="album:%s:%s"
var PrefixSha265   ="sha:%s"

type Resource struct {
	FileSize   int64
	FileName   string
	FilePath   string
	FileType   string
	NameSha256 string
	CTime      string
}

type ResourceSort []Resource

func (r ResourceSort) Len() int {
	return len(r)
}

func (r ResourceSort) Less(i, j int) bool{
	return r[i].FileName < r[j].FileName
}

func (r ResourceSort)  Swap(i, j int){
	r[i],r[j] = r[j],r[i]
}


func (r *Resource) Save() error{
	db := db.GetDb()
	obj, err := json.Marshal(r)
	if err != nil {
		return err
	}
	db.Set(fmt.Sprintf(PrefixSha265,r.NameSha256), obj)
	return db.Set(fmt.Sprintf(PrefixFileName,r.FilePath,r.FileName), obj)
}

func  (r *Resource) Get() {
	db := db.GetDb()
	data,err := db.Get(fmt.Sprintf(PrefixSha265,r.NameSha256))
	if err !=nil{
		return
	}
	json.Unmarshal(data,r)

}
