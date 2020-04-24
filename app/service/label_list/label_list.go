package label_list

import (
	"gf-app/app/model/router_table"
	"github.com/gogf/gf/os/glog"
)



type StuLabelServiceImp struct {
	ILabelService
}


type ILabelService interface{
		All()([]*router_table.Entity,error)
}


func (i *StuLabelServiceImp)All()([]*router_table.Entity,error){
	result, err := router_table.FindAll()
	if err!=nil {
		glog.Error(err)
		return nil,err
	}
	return result,nil
}
