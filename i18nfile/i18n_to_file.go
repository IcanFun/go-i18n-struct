package i18nfile

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
)

type I18n interface {
	ID()string
}

type I18ns interface {
	Items()[]I18n
}

func WriteToFile(i18ns I18ns, path string) error {
	exist := checkFileIsExist(path)

	if !exist {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	items := i18ns.Items()
	if len(items) == 0 {
		return nil
	}

	m := make(map[string][]map[string]string)
	t := reflect.TypeOf(items[0])
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if tag := f.Tag.Get("file"); tag != "" {
			m[tag] = make([]map[string]string, len(items))
		}
	}
	for i, item := range items {
		t := reflect.TypeOf(item)
		v := reflect.ValueOf(item)
		for j := 0; j < t.NumField(); j++ {
			f := t.Field(j)
			if tag := f.Tag.Get("file"); tag != "" {
				s := m[tag]
				s[i] = map[string]string{"id": item.ID(), "translation": v.Field(j).String()}
			}
		}
	}

	for key, value := range m {
		data, _ := json.Marshal(value)
		if err := ioutil.WriteFile(path+"/"+key+".json", data, 0644); err != nil {
			return err
		}
	}

	return nil
}