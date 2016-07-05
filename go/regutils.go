package utils

import (
	"fmt"
	// "fmt"
	// "fmt"
	"html/template"
	// "io/ioutil"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type TaskOption struct {
	Children []Child `json:"children"`
	Folder   string  `json:"folder"`
	Name     string  `json:"name"`
}
type Child struct {
	Children []Control `json:"children"`
	Group    string    `json:"group"`
	Name     string    `json:"name"`
}
type Control struct {
	Default string      `json:"default"`
	Display string      `json:"display"`
	Hint    string      `json:"hint"`
	Name    string      `json:"name"`
	Type    string      `json:"type"`
	Value   interface{} `json:"value"`
}

func CompJson(value string, bkstr string) (result []TaskOption, err error) {
	var m []TaskOption
	err = json.Unmarshal([]byte(bkstr), &m)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	var v map[string]string
	err = json.Unmarshal([]byte(value), &v)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	for i, x := range m {
		for j, y := range x.Children {
			for k, z := range y.Children {
				if va, ok := v[z.Name]; ok {
					m[i].Children[j].Children[k].Default = va
				} else {
					m[i].Children[j].Children[k].Default = ""
				}
			}
		}
	}
	return m, err
}

//T interface{}  un []string 未通过验证字段对应的错误信息
func Tag(T interface{}) (err []string) {
	st := reflect.TypeOf(T)
	sv := reflect.ValueOf(T)
	var v, r, e string
	for i := 0; i < st.NumField(); i++ {
		v = sv.Field(i).String()
		r = st.Field(i).Tag.Get("reg")
		e = st.Field(i).Tag.Get("err")
		if r != "" && v != "" {
			b1, _ := regexp.MatchString(r, v)
			if b1 == false {
				err = append(err, e)
			}
		}
	}
	return
}
func TagTask(T interface{}) (err []string) {
	st := reflect.TypeOf(T)
	sv := reflect.ValueOf(T)
	var v, r, e string
	for i := 0; i < st.NumField(); i++ {
		v = sv.Field(i).String()
		r = st.Field(i).Tag.Get("reg")
		e = st.Field(i).Tag.Get("err")
		if r != "" {
			b1, _ := regexp.MatchString(r, v)
			if b1 == false {
				err = append(err, e)
			}
		}
	}
	return
}

func XSSEscape(str string) string {
	tmp := template.JSEscapeString(str)
	final := template.HTMLEscapeString(tmp)
	return final
}

func TagS(T interface{}) (err []string) {
	val := reflect.ValueOf(T)
	if val.Kind() != reflect.Ptr {
		panic(fmt.Errorf("arg must be use ptr"))
	}
	ind := reflect.Indirect(val)
	etyp := ind.Type()
	typ := etyp
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	eval := val.Elem()
	var v, r, e string
	for i := 0; i < ind.NumField(); i++ {
		if eval.Field(i).Kind() != reflect.String || etyp.Field(i).Tag.Get("type") == "json" {
			continue
		}
		tmp := eval.Field(i).String()
		v = XSSEscape(tmp)
		eval.Field(i).SetString(v)
		r = etyp.Field(i).Tag.Get("reg")
		e = etyp.Field(i).Tag.Get("err")
		if r != "" {
			b1, _ := regexp.MatchString(r, v)
			if b1 == false {
				err = append(err, e)
			}
		}
	}
	return
}

// func InterfaceAppend(s interface{}, s2 interface{}) (sreturn interface{}) {
// 	for _, x := range s2 {
// 		s = append(s, x)
// 	}
// 	return s
// }
func StringAppend(s []string, s2 []string) (sreturn []string) {
	for _, x := range s2 {
		s = append(s, x)
	}
	return s
}
func TemplateParse(data map[string]interface{}) (filepath string, err error) {
	tpath := beego.AppConfig.String("template")
	dpath := beego.AppConfig.String("downloadpath")
	f, err := os.Create(dpath + "output" + time.Now().Format("20060102150405") + ".html")
	if err != nil {
		return "", err
	}
	t3, err := template.ParseFiles(tpath) //将一个文件读作模板
	if err != nil {
		return "", err
	}
	t3.Execute(f, data)
	return f.Name(), err
}

//utils.DeWeight
func DeWeight(sl []string) (s []string) {
	c := true
	for _, x := range sl {
		c = true
		for _, y := range s {
			if x == y {
				c = false
			}
		}
		if c {
			s = append(s, x)

		}
	}
	return
}

func JsonToString(in string) (result string) {
	var s []string
	var re string
	err := json.Unmarshal([]byte(in), &s)
	if err != nil {
		beego.Error(err)
		return
	}
	for i, x := range s {
		re += x
		if i != (len(s) - 1) {
			re += ","
		}
	}
	return re
}

func StringToJson(in string) (result string) {
	sl := strings.Split(in, ",")
	a, err := json.Marshal(&sl)
	if err != nil {
		beego.Error(err)
	}
	return string(a)
}
func CheckUrl(target string) (result string, b bool, err error) {
	for i := 0; i < 3; i++ {
		url := target
		if i >= 1 {
			url = AddUrl(target)
		}
		b, er := SiteCanUse(url)
		if er != nil {
			err = er
		}
		if b {
			return url, b, nil
		}
	}
	return target, false, err
}

func AddUrl(url string) (result string) {
	url = strings.ToLower(url)
	b := strings.Contains(url, "http://")
	c := strings.Contains(url, "https://")
	if b == false && c == false {
		url = `http://` + url
		return url
	}
	if b == true && c == false {
		url = strings.Replace(url, "http://", "https://", 1)
		return url
	}
	if b == false && c == true {
		url = strings.Replace(url, "https://", "http://", 1)
		return url
	}
	return url
}
func SiteCanUse(target string) (b bool, err error) {
	resp, err := http.Get(target)
	if err != nil {
		return false, errors.New("检测失败，该网站无法连接。")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, errors.New("检测失败，该网站无法连接。")
	}
	if string(body) == "" {
		return false, errors.New("检测失败，该网站无法连接。")
	}
	return true, nil
}
