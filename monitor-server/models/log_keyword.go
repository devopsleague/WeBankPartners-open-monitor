package models

type LogKeywordMonitorTable struct {
	Guid string `json:"guid"`
	ServiceGroup string `json:"service_group"`
	LogPath string `json:"log_path"`
	MonitorType string `json:"monitor_type"`
	UpdateTime string `json:"update_time"`
}

type LogKeywordConfigTable struct {
	Guid string `json:"guid"`
	LogKeywordMonitor string `json:"log_keyword_monitor"`
	Keyword string `json:"keyword"`
	Regulative int `json:"regulative"`
	NotifyEnable int `json:"notify_enable"`
	Priority string `json:"priority"`
	UpdateTime string `json:"update_time"`
}

type LogKeywordEndpointRelTable struct {
	Guid string `json:"guid"`
	LogKeywordMonitor string `json:"log_keyword_monitor"`
	SourceEndpoint string `json:"source_endpoint"`
	TargetEndpoint string `json:"target_endpoint"`
}

type LogKeywordServiceGroupObj struct {
	ServiceGroupTable
	Config []*LogKeywordMonitorObj `json:"config"`
}

type LogKeywordMonitorObj struct {
	Guid string `json:"guid"`
	ServiceGroup string `json:"service_group"`
	LogPath string `json:"log_path"`
	MonitorType string `json:"monitor_type"`
	KeywordList []*LogKeywordConfigTable `json:"keyword_list"`
	EndpointRel []*LogKeywordEndpointRelTable `json:"endpoint_rel"`
}