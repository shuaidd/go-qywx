package auth

var (
	CorpConfig *CorpInfo
)

// Callback 企业微信回调配置
type Callback struct {
	Name           string
	Token          string
	EncodingAesKey string
	ReceiveId      string
}

// Application 企业微信应用配置
type Application struct {
	Name    string
	AgentId int
	Secret  string
}

// CorpInfo 企业微信主体信息配置
type CorpInfo struct {
	CorpId    string
	Url       string
	DebugMode bool
	Apps      []Application
	Callbacks []Callback
}
