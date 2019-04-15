package i18nfile

import "testing"

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