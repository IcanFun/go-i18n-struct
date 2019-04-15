package i18nfile

import "testing"

type myI18n struct {
	Id string
	Zh_CN string `file:"zh-cn"`
	En_US string
}

func (m myI18n)ID()string  {
	return m.Id
}

type myI18ns struct {
	i18ns []I18n
}

func (m *myI18ns)Items() []I18n {
	return m.i18ns
}

func TestI18ns_ToFile(t *testing.T) {
	i18ns := make([]I18n,2)
	i18ns[0] = myI18n{Id: "1", Zh_CN: "你好", En_US: "hello"}
	i18ns[1] = myI18n{Id: "2", Zh_CN: "世界", En_US: "world"}

	err := WriteToFile(&myI18ns{i18ns},".")
	if err != nil {
		t.Error(err.Error())
	}
}