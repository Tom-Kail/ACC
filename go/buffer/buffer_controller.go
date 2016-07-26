package controllers

import (
	"bufio"
	"errors"
	"io"
	"os"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func ReadFile(src string, c *MainController) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()
	buf := make([]byte, 102400) //一次读取多少个字节
	bfRd := bufio.NewReader(sf)

	for {
		n, err := bfRd.Read(buf)
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
		_, err3 := c.Ctx.ResponseWriter.Write(buf[:n])
		if err3 != nil {
			return errors.New("Err: " + err3.Error() + "\ndata: " + string(buf[:n]))
		}
		c.Ctx.ResponseWriter.Flush()
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/octet-stream")
	c.Ctx.ResponseWriter.Header().Set("Content-Disposition", "attachment;filename="+src)
	return nil
}
func (this *MainController) Get() {
	src := "G:\\iso\\CentOS-6.5-i386-LiveDVD.iso"
	ReadFile(src, this)
}

