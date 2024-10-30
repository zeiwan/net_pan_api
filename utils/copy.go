package utils

import (
	"github.com/jinzhu/copier"
	"log"
)

// Copy 拷贝结构体
func Copy(toValue interface{}, fromValue interface{}) any {
	if err := copier.Copy(toValue, fromValue); err != nil {
		log.Fatalf("Copy err: err=[%+v]", err)
		panic("拷贝结构体出错")
	}
	return toValue
}
