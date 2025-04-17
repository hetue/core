package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/hetue/boot/internal/internal/internal"
	"github.com/hetue/boot/internal/internal/internal/constant"
	"github.com/hetue/boot/internal/internal/param"
	"github.com/harluo/di"
)

type Bootstrap struct {
	param *param.Bootstrap
}

func NewBootstrap(param *param.Bootstrap) *Bootstrap {
	return &Bootstrap{
		param: param,
	}
}

func (b *Bootstrap) Boot(constructor any) {
	application := di.New()
	if "" != b.param.Name {
		application.Name(b.param.Name)
	}
	if "" != b.param.Usage {
		application.Usage(b.param.Usage)
	}
	if "" != b.param.Copyright {
		application.Copyright(b.param.Copyright)
	}
	for key, value := range b.param.Metadata {
		application.Metadata(key, value)
	}

	application = application.Config().Getter(b).Build()                    // 环境变量
	application.Get().Dependency().Put(constructor).Build().Build().Apply() // 注入所有步骤
	application.Get().Run(internal.NewBootstrap)                            // 执行逻辑
}

func (b *Bootstrap) Get(key string) (value string) {
	value = os.Getenv(fmt.Sprintf("%s_%s", constant.PrefixCi, key))
	if "" == value {
		value = os.Getenv(fmt.Sprintf("%s_%s", constant.PrefixPlugin, key))
	}

	// 修复一些特殊配置项
	value = b.fixDrone(key, value) // Drone系统配置项操蛋的注入方式，部分用JSON部分用环境变量直接注入

	return
}

func (b *Bootstrap) fixDrone(key string, from string) (value string) {
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
		value = b.fixJsonObject(value)
	} else if constant.JsonArrayStart == (from)[0:1] && constant.JsonArrayEnd == (from)[size-1:size] {
		value = b.fixJsonArray(value)
	} else {
		value = from
	}

	return
}

func (b *Bootstrap) fixJsonObject(from string) (to string) {
	object := make(map[string]any)
	if ue := json.Unmarshal([]byte(from), &object); nil != ue {
		to = from
	} else {
		b.fixObjectExpr(object)
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

func (b *Bootstrap) fixJsonArray(from string) (to string) {
	array := make([]any, 0)
	if ue := json.Unmarshal([]byte(from), &array); nil != ue {
		to = from
	} else {
		b.fixArrayExpr(&array)
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

func (b *Bootstrap) fixArrayExpr(array *[]any) {
	for _, value := range *array {
		switch vt := value.(type) {
		case []any:
			b.fixArrayExpr(&vt)
		case map[string]any:
			b.fixObjectExpr(vt)
		}
	}
}

func (b *Bootstrap) fixObjectExpr(object map[string]any) {
	for _, value := range object {
		switch vt := value.(type) {
		case []any:
			b.fixArrayExpr(&vt)
		case map[string]any:
			b.fixObjectExpr(vt)
		}
	}
}
