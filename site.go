package site

const (
	MessageAPI = "/lan/user/msg/add"
)

type Site struct {
	UserId        uint
	UserName      string
	APIUriPrefix  string
	APISignSecret string
}

// setAPIUrlPrefix 设置api schema和host，默认是官网st环境
func (s *Site) setAPIUrlPrefix() {
	if s.APIUriPrefix == "" {
		s.APIUriPrefix = "https://www.st.landui.cn"
	}
}
