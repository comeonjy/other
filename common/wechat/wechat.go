package wechat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Wechat struct {
	accessToken *accessToken
}
type accessToken struct {
	access_token string
	expires_in   int32
}

func (w *Wechat) GetAccessToken() (*accessToken,error) {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=$this->appId&secret=$this->appSecret"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, w.accessToken)
	if err != nil {
		return err
	}
	if w.accessToken.access_token != "" {

	}
}

func httpGet() {

}
