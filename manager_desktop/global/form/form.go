package form

type AdminLoginInput struct {
	Username string `json:"username" comment:"管理员用户名" form:"username" binding:"required,is-validuser"`
	Password string `json:"password" comment:"管理员密码" form:"password" binding:"required"`
}

type ChangePasswordReq struct {
	Password string `json:"password" comment:"管理员密码" form:"password" binding:"required"`
}

type AdminInfoReq struct {
	Token string `json:"token" comment:"token" form:"token" binding:"required"`
}

type AddHTTPServiceReq struct {
	ServiceName string `json:"service_name" form:"service_name" binding:"required,match_service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc" binding:"required,min=0,max=255"`
	// 不用传入，但是这是代表是HTTP/TCP/GRPC的服务的代码
	//LoadType               int    `json:"load_type" form:"load_type"`
	Rule                   string `json:"rule" form:"rule" binding:"required,is_valid_rule"`
	RuleType               int    `json:"rule_type" form:"rule_type" binding:"min=0,max=1"`
	NeedHttps              int    `json:"need_https" form:"need_https" binding:"min=0,max=1"`
	NeedWebsocket          int    `json:"need_websocket" form:"need_websocket" binding:"min=0,max=1"`
	NeedStripUrl           int    `json:"need_strip_uri" form:"need_strip_uri" binding:"min=0,max=1"`
	RoundType              int    `json:"round_type" form:"round_type" binding:"min=0,max=3"`
	UrlRewrite             string `json:"url_rewrite" form:"url_rewrite" binding:"is_valid_url_rewrite"`
	HeaderTransfor         string `json:"header_transfor" form:"header_transfor" binding:"is_valid_header_transfor"`
	IPList                 string `json:"ip_list" form:"ip_list" binding:"required,is_valid_ip_list"`
	ForbidList             string `json:"forbid_list" form:"forbid_list" binding:"required,is_valid_ip_list"`
	WeightList             string `json:"weight_list" form:"weight_list" binding:"required,is_valid_weight_list"`
	BlackList              string `json:"black_list" form:"black_list"`
	WhiteList              string `json:"white_list" form:"white_list"`
	ServiceFlowLimit       int    `json:"service_flow_limit" form:"service_flow_limit" binding:"min=0"`
	ClientFlowLimit        int    `json:"clientip_flow_limit" form:"clientip_flow_limit" binding:"min=0"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" binding:"min=0"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" binding:"min=0"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" binding:"min=0"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" binding:"min=0"`
	CheckMethod            int    `json:"check_method" form:"check_method" binding:"min=0,max=1"`
	CheckTimeout           int    `json:"check_timeout" form:"check_timeout" binding:"min=0"`
	CheckInterval          int    `json:"check_interval" form:"check_interval" binding:"min=0"`
	OpenAuth               int    `json:"open_auth" form:"open_auth" binding:"min=0,max=1"`
}

type UpdateHTTPServiceReq struct {
	Id          int    `json:"id" form:"id"  binding:"required"`
	ServiceName string `json:"service_name" form:"service_name" binding:"required,match_service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc" binding:"required,min=0,max=255"`
	// 不用传入，但是这是代表是HTTP/TCP/GRPC的服务的代码
	//LoadType               int    `json:"load_type" form:"load_type"`
	Rule           string `json:"rule" form:"rule"`
	RuleType       int    `json:"rule_type" form:"rule_type"`
	NeedHttps      int    `json:"need_https" form:"need_https"`
	NeedWebsocket  int    `json:"need_websocket" form:"need_websocket"`
	NeedStripUrl   int    `json:"need_strip_uri" form:"need_strip_uri"`
	RoundType      int    `json:"round_type" form:"round_type"`
	UrlRewrite     string `json:"url_rewrite" form:"url_rewrite"`
	HeaderTransfor string `json:"header_transfor" form:"header_transfor"`
	IPList         string `json:"ip_list" form:"ip_list" binding:"required"`
	ForbidList     string `json:"forbid_list" form:"forbid_list"`

	WeightList             string `json:"weight_list" form:"weight_list"`
	BlackList              string `json:"black_list" form:"black_list"`
	WhiteList              string `json:"white_list" form:"white_list"`
	ServiceFlowLimit       int    `json:"service_flow_limit" form:"service_flow_limit"`
	ClientFlowLimit        int    `json:"clientip_flow_limit" form:"clientip_flow_limit"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle"`
	CheckMethod            int    `json:"check_method" form:"check_method"`
	CheckTimeout           int    `json:"check_timeout" form:"check_timeout"`
	CheckInterval          int    `json:"check_interval" form:"check_interval"`
	OpenAuth               int    `json:"open_auth" form:"open_auth"`
}

type AddTCPServiceReq struct {
	ServiceName string `json:"service_name" form:"service_name" binding:"required,match_service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc" binding:"required,min=0,max=255"`
	// 不用传入，但是这是代表是HTTP/TCP/GRPC的服务的代码
	//LoadType               int    `json:"load_type" form:"load_type"`
	Port int `json:"port" form:"form" binding:"required,min=8001,max=8999"`

	RoundType              int    `json:"round_type" form:"round_type"`
	IPList                 string `json:"ip_list" form:"ip_list" binding:"required,is_valid_ip_list"`
	WeightList             string `json:"weight_list" form:"weight_list" binding:"required,is_valid_weight_list"`
	BlackList              string `json:"black_list" form:"black_list"`
	WhiteList              string `json:"white_list" form:"white_list"`
	ServiceFlowLimit       int    `json:"service_flow_limit" form:"service_flow_limit" binding:"min=0"`
	ClientFlowLimit        int    `json:"clientip_flow_limit" form:"clientip_flow_limit" binding:"min=0"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" binding:"min=0"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" binding:"min=0"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" binding:"min=0"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" binding:"min=0"`
	CheckMethod            int    `json:"check_method" form:"check_method" binding:"min=0,max=1"`
	CheckTimeout           int    `json:"check_timeout" form:"check_timeout" binding:"min=0"`
	CheckInterval          int    `json:"check_interval" form:"check_interval" binding:"min=0"`
	OpenAuth               int    `json:"open_auth" form:"open_auth" binding:"min=0,max=1"`
}

type UpdateTCPServiceReq struct {
	Id          int    `json:"id" form:"id"  binding:"required"`
	ServiceName string `json:"service_name" form:"service_name" binding:"required,match_service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc" binding:"required,min=0,max=255"`
	// 不用传入，但是这是代表是HTTP/TCP/GRPC的服务的代码
	//LoadType               int    `json:"load_type" form:"load_type"`
	Port int `json:"port" form:"form" binding:"required,min=8001,max=8999"`

	RoundType              int    `json:"round_type" form:"round_type"`
	IPList                 string `json:"ip_list" form:"ip_list" binding:"required,is_valid_ip_list"`
	WeightList             string `json:"weight_list" form:"weight_list" binding:"required,is_valid_weight_list"`
	BlackList              string `json:"black_list" form:"black_list"`
	WhiteList              string `json:"white_list" form:"white_list"`
	ServiceFlowLimit       int    `json:"service_flow_limit" form:"service_flow_limit" binding:"min=0"`
	ClientFlowLimit        int    `json:"clientip_flow_limit" form:"clientip_flow_limit" binding:"min=0"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" binding:"min=0"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" binding:"min=0"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" binding:"min=0"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" binding:"min=0"`
	CheckMethod            int    `json:"check_method" form:"check_method" binding:"min=0,max=1"`
	CheckTimeout           int    `json:"check_timeout" form:"check_timeout" binding:"min=0"`
	CheckInterval          int    `json:"check_interval" form:"check_interval" binding:"min=0"`
	OpenAuth               int    `json:"open_auth" form:"open_auth" binding:"min=0,max=1"`
}

type AddGRPCServiceReq struct {
	ServiceName string `json:"service_name" form:"service_name" binding:"required,match_service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc" binding:"required,min=0,max=255"`
	// 不用传入，但是这是代表是HTTP/TCP/GRPC的服务的代码
	//LoadType               int    `json:"load_type" form:"load_type"`
	Port int `json:"port" form:"form" binding:"required,min=8001,max=8999"`

	HeaderTransfor         string `json:"header_transfor" form:"header_transfor" binding:"is_valid_header_transfor"`
	RoundType              int    `json:"round_type" form:"round_type"`
	IPList                 string `json:"ip_list" form:"ip_list" binding:"required,is_valid_ip_list"`
	WeightList             string `json:"weight_list" form:"weight_list" binding:"required,is_valid_weight_list"`
	BlackList              string `json:"black_list" form:"black_list"`
	WhiteList              string `json:"white_list" form:"white_list"`
	ServiceFlowLimit       int    `json:"service_flow_limit" form:"service_flow_limit" binding:"min=0"`
	ClientFlowLimit        int    `json:"clientip_flow_limit" form:"clientip_flow_limit" binding:"min=0"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" binding:"min=0"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" binding:"min=0"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" binding:"min=0"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" binding:"min=0"`
	CheckMethod            int    `json:"check_method" form:"check_method" binding:"min=0,max=1"`
	CheckTimeout           int    `json:"check_timeout" form:"check_timeout" binding:"min=0"`
	CheckInterval          int    `json:"check_interval" form:"check_interval" binding:"min=0"`
	OpenAuth               int    `json:"open_auth" form:"open_auth" binding:"min=0,max=1"`
}

type UpdateGRPCServiceReq struct {
	Id          int    `json:"id" form:"id"  binding:"required"`
	ServiceName string `json:"service_name" form:"service_name" binding:"required,match_service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc" binding:"required,min=0,max=255"`
	// 不用传入，但是这是代表是HTTP/TCP/GRPC的服务的代码
	//LoadType               int    `json:"load_type" form:"load_type"`
	Port int `json:"port" form:"form" binding:"required,min=8001,max=8999"`

	HeaderTransfor         string `json:"header_transfor" form:"header_transfor" binding:"is_valid_header_transfor"`
	RoundType              int    `json:"round_type" form:"round_type"`
	IPList                 string `json:"ip_list" form:"ip_list" binding:"required,is_valid_ip_list"`
	WeightList             string `json:"weight_list" form:"weight_list" binding:"required,is_valid_weight_list"`
	BlackList              string `json:"black_list" form:"black_list"`
	WhiteList              string `json:"white_list" form:"white_list"`
	ServiceFlowLimit       int    `json:"service_flow_limit" form:"service_flow_limit" binding:"min=0"`
	ClientFlowLimit        int    `json:"clientip_flow_limit" form:"clientip_flow_limit" binding:"min=0"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" binding:"min=0"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" binding:"min=0"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" binding:"min=0"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" binding:"min=0"`
	CheckMethod            int    `json:"check_method" form:"check_method" binding:"min=0,max=1"`
	CheckTimeout           int    `json:"check_timeout" form:"check_timeout" binding:"min=0"`
	CheckInterval          int    `json:"check_interval" form:"check_interval" binding:"min=0"`
	OpenAuth               int    `json:"open_auth" form:"open_auth" binding:"min=0,max=1"`
}

type DeleteHTTPServiceReq struct {
	Id int `json:"id" query:"id" form:"id"`
}

type ServiceStatReq struct {
	Id int `json:"id" query:"id" form:"id"`
}

type ServiceDetailReq struct {
	Id int `json:"id" query:"id" form:"id" binding:"required"`
}

type ServiceListReq struct {
	PageNum  int    `json:"page_no" query:"page_no" form:"page_no"`
	PageSize int    `json:"page_size" query:"page_size" form:"page_size"`
	Info     string `json:"info" query:"info" form:"info"`
}

type AppListReq struct {
	PageNum  int    `json:"page_no" query:"page_no" form:"page_no"`
	PageSize int    `json:"page_size" query:"page_size" form:"page_size"`
	Info     string `json:"info" query:"info" form:"info"`
}

type AddUserReq struct {
	AppId  string `json:"app_id" form:"app_id" binding:"required,match_service_name"`
	Name   string `json:"name" form:"name" binding:"required,max=255"`
	Secret string `json:"secret" form:"secret"`
	QPS    int    `json:"qps" form:"qps" binding:"min=0"`
	QPD    int    `json:"qpd" form:"qpd" binding:"min=0"`
}

type UpdateUserReq struct {
	Id     int    `json:"id" query:"id" form:"id" binding:"required"`
	AppId  string `json:"app_id" form:"app_id" binding:"required,match_service_name"`
	Name   string `json:"name" form:"name" binding:"required,max=255"`
	Secret string `json:"secret" form:"secret" binding:"min=32,max=32"`
	QPS    int    `json:"qps" form:"qps" binding:"min=0"`
	QPD    int    `json:"qpd" form:"qpd" binding:"min=0"`
}

type UserDetailReq struct {
	Id int `json:"id" query:"id" form:"id" binding:"required"`
}

type DeleteUserReq struct {
	Id int `json:"id" query:"id" form:"id" binding:"required"`
}

type UserStatReq struct {
	Id int `json:"id" query:"id" form:"id" binding:"required"`
}
