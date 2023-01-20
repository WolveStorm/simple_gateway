package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"gateway_server/cache/model"
	"gateway_server/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func RspData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  Msg(CodeSuccess),
		Data: data,
	})
}

func RspError(c *gin.Context, code ResCode, err error) {
	if err == nil {
		c.JSON(http.StatusOK, Response{
			Code: code,
			Msg:  Msg(code),
			Data: nil,
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Code: code,
			Msg:  err.Error(),
			Data: nil,
		})
	}
}

func GetSaltPassword(salt, password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	shaPass := fmt.Sprintf("%x", hash.Sum(nil))
	hash1 := sha256.New()
	hash1.Write([]byte(shaPass + salt))
	return fmt.Sprintf("%x", hash1.Sum(nil))
}

func MD5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", md5.Sum(nil))
}

func getKind(val reflect.Value) reflect.Kind {
	kind := val.Kind()

	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float32
	default:
		return kind
	}
}

// key为json tag，val为值
// 必须保证没有不可导出的属性或者方法，否则会panic
// 该方法一般用于gorm的更新时map转换，默认删除主键更新
func StructToUpdateMap(arg interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	t := reflect.TypeOf(arg).Elem()
	v := reflect.ValueOf(arg).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("ignore") != "" {
			continue
		}
		val := v.Field(i)
		switch getKind(val) {
		case reflect.Bool:
			m[field.Tag.Get("json")] = val.Bool()
		case reflect.Interface:
			m[field.Tag.Get("json")] = val.Interface()
		case reflect.String:
			m[field.Tag.Get("json")] = val.String()
		case reflect.Int, reflect.Int32, reflect.Int16, reflect.Int64, reflect.Int8:
			m[field.Tag.Get("json")] = val.Int()
		case reflect.Uint, reflect.Uint32, reflect.Uint16, reflect.Uint64, reflect.Uint8:
			m[field.Tag.Get("json")] = val.Uint()
		case reflect.Float32, reflect.Float64:
			m[field.Tag.Get("json")] = val.Float()
		case reflect.Struct:
			// 不可以直接用val.Interface()
			value := reflect.New(field.Type)
			reflect.Indirect(value).Set(val)
			m[field.Tag.Get("json")] = StructToUpdateMap(value.Interface())
		case reflect.Map:
			mj, _ := json.Marshal(val.Interface())
			m[field.Tag.Get("json")] = string(mj)
		case reflect.Slice:
			sj, _ := json.Marshal(val.Interface())
			m[field.Tag.Get("json")] = string(sj)
		case reflect.Array:
			aj, _ := json.Marshal(val.Interface())
			m[field.Tag.Get("json")] = string(aj)
		case reflect.Ptr:
			data := val.Interface()
			isNil := data == nil
			if !isNil {
				switch v := reflect.Indirect(reflect.ValueOf(data)); v.Kind() {
				case reflect.Chan,
					reflect.Func,
					reflect.Interface,
					reflect.Map,
					reflect.Ptr,
					reflect.Slice:
					isNil = v.IsNil()
				}
			}
			if isNil {
				if !val.IsNil() && val.CanSet() {
					nilValue := reflect.New(val.Type()).Elem()
					m[field.Tag.Get("json")] = nilValue.Interface()
				}
			} else {
				m[field.Tag.Get("json")] = val.Interface()
			}
		default:
			// If we reached this point then we weren't able to decode it
			return nil
		}
	}
	delete(m, "id")
	return m
}

func GetServiceDetail(c *gin.Context) (*model.ServiceDetail, error) {
	serviceDetail, ok := c.Get(global.ServiceDetail)
	if !ok {
		return nil, errors.New("error")
	}
	detail := &model.ServiceDetail{}
	serviceDetailStr := serviceDetail.(string)
	err := json.Unmarshal([]byte(serviceDetailStr), detail)
	if err != nil {
		return nil, err
	}
	return detail, nil
}

func GetUser(c *gin.Context) (*model.User, error) {
	userDetail, ok := c.Get(global.AppDetail)
	if !ok {
		return nil, errors.New("error")
	}
	detail := &model.User{}
	serviceDetailStr := userDetail.(string)
	err := json.Unmarshal([]byte(serviceDetailStr), detail)
	if err != nil {
		return nil, err
	}
	return detail, nil
}
