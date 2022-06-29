package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
	"web_app/models"
)

var secret = []byte("雪下的是盐")

//CheckUserExist 检查用户是否存在
func CheckUserExist(username string) error {

	sqlstr := "select count(user_id) from user where username=?"
	var count int
	if err := db.Get(&count, sqlstr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrUserExist
	}
	return nil
}

//InsertUser 向用户表中插入一条记录
func InsertUser(user *models.User) error {

	//对密码经行处理
	//密码和盐,密码在前,盐在后
	pwdandsalt := encryptPassword(user.Password)
	sqlstr := "insert into user(user_id,username,password,salt)values(?,?,?,?)"

	_, err := db.Exec(sqlstr, user.UserId, user.Username, pwdandsalt[0], pwdandsalt[1])

	return err
}

//登入函数
func Login(p *models.User) error {
	u := &models.User{}
	sqlstr := "select user_id,username,password,salt from user where username=?"
	err := db.Get(u, sqlstr, p.Username)
	if err == sql.ErrNoRows {
		return ErrUserNotExist
	}
	if err != nil {
		//用户名查询错误
		return err
	}

	//密码的判断
	p.Password = loginEncryptPassword(p.Password, u.Salt)
	if p.Password != u.Password {
		return ErrInvalidParam
	}
	p.UserId = u.UserId

	return nil
}
func GetUserById(id int64) (string, error) {

	sqlstr := "select username from user where user_id=?"
	var name string
	err := db.Get(&name, sqlstr, id)
	if err != nil {
		return "", err
	}
	return name, nil
}

//注册时候返回md5加密和随机的盐
func encryptPassword(pwd string) []string {
	h := md5.New()
	salt := getRandstring()
	h.Write([]byte(salt))
	h.Write(secret)
	return []string{hex.EncodeToString(h.Sum([]byte(pwd))), salt}
}

//loginEncryptPassword 登入时候校验密码使用数据库中的盐
func loginEncryptPassword(pwd, salt string) string {

	h := md5.New()
	h.Write([]byte(salt))
	h.Write(secret)
	return hex.EncodeToString(h.Sum([]byte(pwd)))
}

//获取随机的字符串
func getRandstring() string {

	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))
	var rchar string = ""
	for i := 1; i <= 6; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar
}
