package album

import (
	"encoding/json"
	"fmt"
	"github.com/markusleevip/taostorage/db/model"
)

func mapsToList(maps map[string]string) (interface{}, error){

	if maps!=nil && len(maps)>0 {
		list := make([]model.Resource,0)
		for _,value := range maps {
			fmt.Println(value)
			if value !="" {
				res :=model.Resource{}
				json.Unmarshal([]byte(value),&res)
				list = append(list,res)
			}
		}
		return list,nil

	}
	return nil, nil
}