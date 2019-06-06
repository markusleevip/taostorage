package model

import (
	"encoding/json"
	"github.com/markusleevip/taostorage/db"
)

type Resource struct {
	FileSize   int64
	FileName   string
	FilePath   string
	FileType   string
	NameSha256 string
	CTime      string
}

func (r *Resource) Save() error{
	db := db.GetDb()
	obj, err := json.Marshal(r)
	if err != nil {
		return err
	}
	return db.Set(r.NameSha256, obj)
}
