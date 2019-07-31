package model

import (
	"encoding/json"
	"github.com/markusleevip/taostorage/db"
)

var prefix ="album:"

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
	return db.Set(prefix+r.NameSha256, obj)
}

func  (r *Resource) Get() {
	db := db.GetDb()

	data,err := db.Get(prefix+r.NameSha256)
	if err !=nil{
		return
	}
	json.Unmarshal(data,r)

}
