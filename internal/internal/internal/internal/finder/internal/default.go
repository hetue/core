package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/hetue/boot/internal/internal/internal/constant"
)

type Default struct {
	// 无字段
}

func newDefault() *Default {
	return new(Default)
}

func (d *Default) Find(key string) (value string) {
	value = os.Getenv(fmt.Sprintf("%s_%s", constant.PrefixCi, key))
	if "" == value {
		value = os.Getenv(fmt.Sprintf("%s_%s", constant.PrefixPlugin, key))
	}

	// 修复一些特殊配置项
	value = d.fixDrone(key, value) // Drone系统配置项操蛋的注入方式，部分用JSON部分用环境变量直接注入

	return
}

func (d *Default) fixDrone(key string, from string) (value string) {
	if "" == os.Getenv(constant.PlatformDrone) {
		return
	}

	if "" == strings.TrimSpace(from) { // 修正空值
		from = os.Getenv(fmt.Sprintf("%s_%s", constant.PrefixDrone, key))
	}
	if "" == strings.TrimSpace(from) { // 及时回退，如果确实没有配置值
		return
	}

	size := len(from)
	if constant.JsonObjectStart == (from)[0:1] && constant.JsonObjectEnd == (from)[size-1:size] {
		value = d.fixJsonObject(value)
	} else if constant.JsonArrayStart == (from)[0:1] && constant.JsonArrayEnd == (from)[size-1:size] {
		value = d.fixJsonArray(value)
	} else {
		value = from
	}

	return
}

func (d *Default) fixJsonObject(from string) (to string) {
	object := make(map[string]any)
	if ue := json.Unmarshal([]byte(from), &object); nil != ue {
		to = from
	} else {
		d.fixObjectExpr(object)
	}

	if from == to {
		// 不需要进行转换
	} else if bytes, me := json.Marshal(object); nil != me {
		to = from
	} else {
		to = string(bytes)
	}

	return
}

func (d *Default) fixJsonArray(from string) (to string) {
	array := make([]any, 0)
	if ue := json.Unmarshal([]byte(from), &array); nil != ue {
		to = from
	} else {
		d.fixArrayExpr(&array)
	}

	if from == to {
		// 不需要进行转换
	} else if bytes, me := json.Marshal(array); nil != me {
		to = from
	} else {
		to = string(bytes)
	}

	return
}

func (d *Default) fixArrayExpr(array *[]any) {
	for _, value := range *array {
		switch vt := value.(type) {
		case []any:
			d.fixArrayExpr(&vt)
		case map[string]any:
			d.fixObjectExpr(vt)
		}
	}
}

func (d *Default) fixObjectExpr(object map[string]any) {
	for _, value := range object {
		switch vt := value.(type) {
		case []any:
			d.fixArrayExpr(&vt)
		case map[string]any:
			d.fixObjectExpr(vt)
		}
	}
}
