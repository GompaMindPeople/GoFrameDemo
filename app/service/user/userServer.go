package user

import (
	"errors"
	"gf-app/app/model/user"
)

type StructUserServer struct {

}


var  dao = user.Model

func (s *StructUserServer) Login(account,password string)(*user.Entity,error){

	entity, e := dao.FindOne("account = ? and password = ?", account, password)

	if e != nil{
		return nil,errors.New("查询数据库的时候发生错误"+e.Error())
	}
	if entity == nil{
		return nil,errors.New("帐号或密码错误")
	}

	return entity,nil
}
