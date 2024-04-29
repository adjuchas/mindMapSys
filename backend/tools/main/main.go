package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
)

type Draft struct {
	DraftID      int    `gorm:"column:draft_ID;primaryKey" json:"draftID"`
	TITLE        string `gorm:"column:TITLE" json:"title"`
	KsID         string `gorm:"column:ks_ID" json:"ksID"`
	STATE        int    `gorm:"column:STATE" json:"state"`
	TAGS         string `gorm:"column:TAGS" json:"tags"`
	DESCRIPTION  string `gorm:"column:DESCRIPTION" json:"description"`
	NodeTreePath string `gorm:"column:nodeTreePath" json:"nodeTreePath"`
	CreateTime   string `gorm:"column:CREATETIME" json:"createTime"`
	UpdateTime   string `gorm:"column:UPDATETIME" json:"updateTime"`
}

func (receiver Draft) TableName() string {
	return "draftinfo"
}

func main() {
	sourceFile, err := os.Open("a.json")
	if err != nil {
		fmt.Println("open source err")
	}
	defer sourceFile.Close()
	fileName := filepath.Base("./a.json")
	tar := filepath.Join("./dot", fileName)
	destFile, _ := os.Create(tar)
	_, err = io.Copy(destFile, sourceFile)
}

func mapToStruct(m map[string]string, s interface{}) error {
	// 获取结构体的反射值
	structValue := reflect.ValueOf(s)
	if structValue.Kind() != reflect.Ptr || structValue.IsNil() {
		return errors.New("s must be a non-nil pointer to struct")
	}
	// 获取结构体的类型
	structType := structValue.Elem().Type()

	// 确保结构体是指针指向的结构体类型
	if structType.Kind() != reflect.Struct {
		return errors.New("s must be a pointer to struct")
	}

	// 遍历 map 中的键值对
	for k, v := range m {
		// 获取结构体字段
		field, ok := structType.FieldByName(k)
		if !ok {
			continue
		}
		// 如果字段存在且可设置，则设置其值
		if structValue.Elem().FieldByName(k).CanSet() {
			// 将字符串类型的值转换为字段的类型
			fieldValue := reflect.ValueOf(v)
			fieldType := field.Type
			if fieldValue.Type().ConvertibleTo(fieldType) {
				structValue.Elem().FieldByName(k).Set(fieldValue.Convert(fieldType))
			} else {
				return errors.New("value type does not match field type")
			}
		}
	}

	return nil
}

//func connredis() {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "117.72.14.167:6379",
//		Password: "adjuchas12", // 没有密码，默认值
//		DB:       0,            // 默认DB 0
//	})
//	val, err := rdb.LRange(context.Background(), "zz", int64(0), int64(-1)).Result()
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(val)
//	fmt.Println(reflect.TypeOf(val))
//}

//func connmysql()  {
//	dsn := "root:adjuchas@tcp(127.0.0.1:3306)/backend?charset=utf8mb4&parseTime=True&loc=Local"
//	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	var tag []taginfo
//	db.Select("tag_ID", "NAME", "dot_NUM").Find(&tag)
//	for _, user := range tag {
//		fmt.Println(user)
//	}
//}
