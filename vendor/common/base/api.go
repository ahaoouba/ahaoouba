package base

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type ApiResult struct {
	Apiversion string                 `json:"apiversion"`
	Metadata   MetaData               `json:"metadata"`
	Result     map[string]interface{} `json:"result"`
	Success    bool                   `json:"success"`
	Error      ErrorInfo              `json:"error"`
}

type ErrorInfo struct {
	Msg string `json:"msg"`
}

type MetaData struct {
	Link string `json:"link"`
}

type ResultPage struct {
	Total  int
	Offset int
	Limit  int
}

type Invoker struct {
	InvokerId string `json:"invokerid"`
	Name      string `json:"name"`
}

func ParseApiResult(b []byte) (*ApiResult, error) {
	result := new(ApiResult)
	return result, json.Unmarshal(b, result)
}

func (ar *ApiResult) ParsePage() *ResultPage {
	rp := new(ResultPage)
	bys, _ := json.Marshal(ar.Result["page"])
	json.Unmarshal(bys, rp)
	return rp
}

//func GetSuccessPageResult(data interface{}) map[string]interface{} {
//	result := make(map[string]interface{})
//	result["success"] = true
//	result["data"] = data
//	return result
//}

//func GetErrorPageResult(err string) map[string]interface{} {
//	result := make(map[string]interface{})
//	result["success"] = false
//	result["errmsg"] = err
//	return result
//}

func NewApiResult(version, apiLink string) *ApiResult {
	result := new(ApiResult)
	result.Apiversion = version
	result.Metadata.Link = apiLink
	result.Success = true
	result.Result = make(map[string]interface{})
	return result
}

//生成正确返回结果
func (result *ApiResult) GenerateSuccessApiResult(data map[string]interface{}) {

	for k, v := range data {
		result.Result[k] = v
	}

	result.Success = true

}

//生成错误返回结果
func (result *ApiResult) GenerateErrorApiResult(errmsg string) {
	result.Error.Msg = errmsg
	result.Success = false
	return
}

func GetApiLink(req *http.Request) (string, error) {
	if req != nil {
		//		fmt.Println(*req)

		//				url.Parse()
		//		u, _ := url.Parse(req.RequestURI)
		//		u = u.ResolveReference(u)
		u, err := url.QueryUnescape(req.RequestURI)
		if err != nil {
			return "", err
		}
		return u, nil
	}
	return "", nil
}
