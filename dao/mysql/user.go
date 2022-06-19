package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math/rand"
	"strings"
	"time"
	"web_app/models"
)

//CheckUserExist 检查用户是否存在
func CheckUserExist(username string) error {

	sqlstr := "select count(user_id) from user where username=?"
	var count int
	if err := db.Get(&count, sqlstr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已经存在")
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

func encryptPassword(pwd string) []string {
	h := md5.New()
	salt := getRandstring()
	h.Write([]byte(salt))
	return []string{hex.EncodeToString(h.Sum([]byte(pwd))), salt}
}

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
