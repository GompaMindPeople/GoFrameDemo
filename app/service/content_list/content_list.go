package content_list

import (
	"gf-app/app/model/content_list"
)

type StuContentListServiceImp struct {
	IContentListService
}


type IContentListService interface{
	GetByLabelId(string,string)([]*content_list.Entity,error)
}

func (s *StuContentListServiceImp)GetByLabelId(id,pageNumber string)([]*content_list.Entity,error){
	result, err := content_list.FindAll("rt_id = ? order by create_time desc limit ?,10",id,pageNumber)
	if err != nil {
		return nil,err
	}
	return result,nil
}

func (s *StuContentListServiceImp)GetByContentId(id string)(*content_list.Entity,error){
	result, err := content_list.FindOne("id = ? ",id)
	if err != nil {
		return nil,err
	}
	return result,nil
}

func (s *StuContentListServiceImp)InsertLabel(entity *content_list.Entity)(error){
	_, err := content_list.Insert(entity)
	if err != nil {
		return err
	}
	return nil
}