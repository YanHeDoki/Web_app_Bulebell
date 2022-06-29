package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/snowflake"
	"web_app/utils"
)

func SignUp(p *models.ParamSignUp) error {

	//判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	//生成uuid
	uuid := snowflake.GenID()

	user := models.User{
		UserId:   uuid,
		Username: p.Username,
		Password: p.Password,
	}

	//用户数据处理

	//保存到数据库

	if err := mysql.InsertUser(&user); err != nil {
		return err
	}
	return nil
}

func Login(p *models.ParamLogin) (user *models.User, err error) {

	//用户实例
	u := &models.User{Username: p.Username, Password: p.Password}

	//登入逻辑
	if err = mysql.Login(u); err != nil {
		//登入失败
		return nil, err
	}

	token, err := utils.GenToken(u.UserId, u.Username)
	if err != nil {
		return nil, err
	}
	u.Token = token
	return u, nil
}
