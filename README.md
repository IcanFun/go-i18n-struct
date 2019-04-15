go-i18n-struct 
=======

[nicksnyder/go-i18n](https://github.com/nicksnyder/go-i18n)
是以文件来实现多语言，但是语言类型多的时候，维护和更新多个文件很麻烦，
所以做了这么一个功能，自动生成多语言文件。

目前只支持json，欢迎大家fork

Installation 
---
go get -u github.com/IcanFun/go-i18n-struct

Workflow
---
实现接口
```
type I18n interface {
	ID()string
}
```

I18n 的tag `file:name` 将生成文件name.json

最后调用```i18nfile.WriteToFile(i18ns []I18n, path string)```写入文件

Example
---
```
type myI18n struct {
	Id string
	Zh_CN string `file:"zh-cn"`
	En_US string `file:"en-us"`
}

func (m myI18n)ID()string  {
	return m.Id
}

func TestI18ns_ToFile(t *testing.T) {
	i18ns := make([]I18n,2)
	i18ns[0] = myI18n{Id: "1", Zh_CN: "你好", En_US: "hello"}
	i18ns[1] = myI18n{Id: "2", Zh_CN: "世界", En_US: "world"}

	err := WriteToFile(i18ns,".")
	if err != nil {
		t.Error(err.Error())
	}
}
 ```
