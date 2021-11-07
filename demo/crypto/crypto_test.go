package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// ////////////////////////////////////////////////////////////////////////////
// 哈希算法
//  对用户输入的密码进行加密
//  用户登录时对用户的密码进行比对
func TestHash(t *testing.T) {
	userPassword := "123456"
	passwordbyte, err := GeneratePassword(userPassword)
	if err != nil {
		t.Fatal("generate password failed")
	}

	passwordstring := string(passwordbyte)
	fmt.Println(passwordstring)

	// 模拟这个字符串是从数据库读取出来的 值是12345678
	mysqlPwd := "$2a$10$rWPez92DnGh14R3T0GsMUeUQKETmmUVAJCRgUaqnPu8.6SDnKUy22"
	ok, err := ValidatePassword(userPassword, mysqlPwd)
	if !ok {
		t.Fatal("validate password failed", err)
	}

	mysqlErrorPwd := "$qweqeqweErrPassword.6SDnKUy22"
	ok, err = ValidatePassword(userPassword, mysqlErrorPwd)
	if ok {
		t.Fatal("validate password failed", err)
	}
}

// GeneratePassword 给密码就行加密操作
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// ValidatePassword 密码比对
func ValidatePassword(userPassword string, hashed string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误！")
	}
	return true, nil
}

// ////////////////////////////////////////////////////////////////////////////
// md5
func TestMD5(t *testing.T) {
	str := "www.topgoer.com"

	// 方法一
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) // 将[]byte转成16进制

	// 方法二
	w := md5.New()
	_, _ = io.WriteString(w, str)     // 将str写入到w中
	bw := w.Sum(nil)                  // w.Sum(nil)将w的hash转成[]byte格式
	md5str2 := fmt.Sprintf("%x", bw)  // 将 bw 转成字符串
	md5str3 := hex.EncodeToString(bw) // 将 bw 转成字符串

	if md5str1 != md5str2 && md5str2 != md5str3 {
		t.Fatal("failed")
	}
	fmt.Println(md5str2)
}

func TestBase64(t *testing.T) {
	str := "www.example.com?age=我的"

	encodeString := base64.StdEncoding.EncodeToString([]byte(str))
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("encodeString", encodeString)
	t.Log("decodeString", string(decodeBytes))

	// 如果要用在url中，需要使用URLEncoding
	urlencode := base64.URLEncoding.EncodeToString([]byte(str))
	t.Logf("urlencode : %v\n", urlencode)

	urldecode, err := base64.URLEncoding.DecodeString(urlencode)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("urldecode : %v\n", string(urldecode))
}
