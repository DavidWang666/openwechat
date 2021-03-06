
package openwechat

// WechatStorage
// 可以根据自己的情况来实现该接口
type WechatStorage interface {
	SetLoginInfo(loginInfo LoginInfo)
	SetBaseRequest(baseRequest BaseRequest)
	SetWebInitResponse(webInitResponse WebInitResponse)
	GetLoginInfo() LoginInfo
	GetBaseRequest() BaseRequest
	GetWebInitResponse() WebInitResponse
}

// implement WechatStorage
// WechatStorage接口的实现
type SimpleWechatStorage struct {
	loginInfo       LoginInfo
	baseRequest     BaseRequest
	webInitResponse WebInitResponse
}

func NewSimpleWechatStorage() *SimpleWechatStorage {
	return &SimpleWechatStorage{}
}

func (s *SimpleWechatStorage) SetLoginInfo(loginInfo LoginInfo) {
	s.loginInfo = loginInfo
}

func (s *SimpleWechatStorage) SetBaseRequest(baseRequest BaseRequest) {
	s.baseRequest = baseRequest
}

func (s *SimpleWechatStorage) SetWebInitResponse(webInitResponse WebInitResponse) {
	s.webInitResponse = webInitResponse
}

func (s *SimpleWechatStorage) GetLoginInfo() LoginInfo {
	return s.loginInfo
}

func (s *SimpleWechatStorage) GetBaseRequest() BaseRequest {
	return s.baseRequest
}

func (s *SimpleWechatStorage) GetWebInitResponse() WebInitResponse {
	return s.webInitResponse
}
