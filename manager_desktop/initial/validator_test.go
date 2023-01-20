package initial

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"simple_gateway/global"
	"simple_gateway/global/form"
	"strings"
	"testing"
)

func TestInitValidator(t *testing.T) {
	ReplaceGinBinding("zh")
	var invalid = form.AdminLoginInput{
		Username: "aaa",
		Password: "aaa",
	}
	var valid = form.AdminLoginInput{
		Username: "admin",
		Password: "aaa",
	}
	v := binding.Validator
	err := v.ValidateStruct(invalid)
	errs := err.(validator.ValidationErrors)
	errMap := errs.Translate(global.Trans)
	graceMap := make(map[string]string)
	for k, v := range errMap {
		graceMap[k[strings.Index(k, ".")+1:]] = v
	}
	fmt.Println(graceMap)
	if err == nil {
		t.Fatal("校验错误")
	}
	err = v.ValidateStruct(valid)
	if err != nil {
		t.Fatal("校验错误")
	}
}
