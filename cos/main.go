package main

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main(){
	upload()
	download()

}
func upload() {
	//将<bucketname>、<appid>和<region>修改为真实的信息
	//例如：http://test-1253846586.cos.ap-guangzhou.myqcloud.com
	u, _ := url.Parse("http://yhbaas-1255000078.cos.shanghai.tce.yonghuicloud.cn")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  "AKIDC7RSAZVSw6xXpoOqtfDUSDn40ex06KXR",
			SecretKey: "oJVOs5zZqnQjhLeeKgRM1NSFjumKZV8H",
		},
	})
	//对象键（Key）是对象在存储桶中的唯一标识。
	//例如，在对象的访问域名 ` bucket1-1250000000.cos.ap-guangzhou.myqcloud.com/test/objectPut.go ` 中，对象键为 test/objectPut.go
	name := "test/upload.txt"
	//Local file
	f := strings.NewReader("yes,I know!")
	_, err := c.Object.Put(context.Background(), name, f, nil)
	if err != nil {
		panic(err)
	}
}

func download() {
	//将<bucketname>、<appid>和<region>修改为真实的信息
	//例如：http://test-1253846586.cos.ap-guangzhou.myqcloud.com
	//u, _ := url.Parse("http://<bucketname>-<appid>.cos.<region>.myqcloud.com")
	u, _ := url.Parse("http://yhbaas-1255000078.cos.shanghai.tce.yonghuicloud.cn")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			//SecretID:  os.Getenv("COS_SECRETID"),
			//SecretKey: os.Getenv("COS_SECRETKEY"),
			SecretID:  "AKIDC7RSAZVSw6xXpoOqtfDUSDn40ex06KXR",
			SecretKey: "oJVOs5zZqnQjhLeeKgRM1NSFjumKZV8H",
		},
	})
	//Object key
	name := "test/upload.txt"
	resp, err := c.Object.Get(context.Background(), name, nil)
	if err != nil {
		panic(err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", string(bs))
}
