package v1

import (
	"bytes"
	"common/base"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Server string
}

func NewClient(s string) *Client {
	return &Client{
		Server: s,
	}
}

//注册服务
func (c *Client) RegistService(s Service) (Service, error) {

	err := s.ValidRegist()
	if err != nil {
		return s, err
	}

	url := fmt.Sprintf("http://%s/platform_registry/v1/service", c.Server)

	bys, err := json.Marshal(s)
	if err != nil {
		return s, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bys))
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return s, err
	}

	if response.StatusCode != 200 {
		return s, fmt.Errorf("注册服务失败，返回状态码为：[ %s ]", response.StatusCode)

	}

	result := base.ApiResult{}

	defer response.Body.Close()

	bys, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(bys, &result)
	if err != nil {
		return s, err
	}

	if !result.Success {
		return s, fmt.Errorf(result.Error.Msg)
	}

	return s, nil
}

//删除服务
func (c *Client) UnRegistService(name, mode string) error {

	if name == "" {
		return fmt.Errorf("name 不能为空")
	}

	url := fmt.Sprintf("http://%s/platform_registry/v1/service/%s/%s", c.Server, name, mode)

	req, err := http.NewRequest("DELETE", url, nil)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("删除注册服务失败，返回状态码为：[ %s ]", response.StatusCode)
	}

	return nil
}

//获取服务信息
func (c *Client) GetService(name, mode string) ([]Service, error) {
	if name == "" {
		return make([]Service, 0), fmt.Errorf("name 不能为空")
	}
	url := fmt.Sprintf("http://%s/platform_registry/v1/service/%s/%s", c.Server, name, mode)

	req, err := http.NewRequest("GET", url, nil)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return make([]Service, 0), err
	}

	if response.StatusCode != 200 {
		return make([]Service, 0), fmt.Errorf("获取注册服务失败，返回状态码为：[ %s ]", response.StatusCode)
	}

	result := base.ApiResult{}

	defer response.Body.Close()

	bys, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return make([]Service, 0), err
	}

	err = json.Unmarshal(bys, &result)
	if err != nil {
		return make([]Service, 0), err
	}

	if !result.Success {
		return make([]Service, 0), fmt.Errorf(result.Error.Msg)
	}

	list := make([]Service, 0)
	bys, _ = json.Marshal(result.Result["service"])
	json.Unmarshal(bys, &list)
	return list, nil
}
