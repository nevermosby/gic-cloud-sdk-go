package gic

// DataCenterList for gic virtual datacenter list
type DataCenterList struct {
	Status  string               `json:"status"`
	CodeMsg string               `json:"code_msg"`
	Message string               `json:"message"`
	Code    int                  `json:"code"`
	Data    []DataCenterListData `json:"data"`
}

// DataCenterListData for the `data` property of gic virtual datacenter list
type DataCenterListData struct {
	Resource     DataCenterListDataResource `json:"resource"`
	Name         string                     `json:"name"`
	CustomerUser string                     `json:"customer_user"`
	SiteID       string                     `json:"site_id"`
	AppID        string                     `json:"app_id"`
	SiteName     string                     `json:"site_name"`
}

// DataCenterListDataResource for the `resource` property of gic virtual datacenter list
type DataCenterListDataResource struct {
	Name     string `json:"name"`
	GicCount int    `json:"gic_count"`
	WanCount int    `json:"wan_count"`
	LanCount int    `json:"lan_count"`
	ID       string `json:"id"`
	VMCount  int    `json:"vm_count"`
}

// DataCenterInfo for single gic virtual datacenter info
type DataCenterInfo struct {
	Status  string               `json:"status"`
	CodeMsg string               `json:"code_msg"`
	Message string               `json:"message"`
	Code    int                  `json:"code"`
	Data    []DataCenterInfoData `json:"data"`
}

type DataCenterInfoData struct {
	VMCount  int    `json:"vms_count"`
	SiteName string `json:"site_name"`
	Name     string `json:"name"`
	// TODO: maybe not []string
	GPNList      []string            `json:"gpns"`
	NetworkList  []DataCenterNetwork `json:"net"`
	CustomerUser string              `json:"customer_user"`
	SiteID       string              `json:"site_id"`
	AppID        string              `json:"app_id"`
	// TODO: need to define vm struct
	// VMList string `json:"vms"`
}

type DataCenterNetwork struct {
	Status string `json:"status"`
	Type   string `json:"type"`
	QoS    int    `json:"qos"`
	Name   string `json:"name"`
	// TODO: maybe not []string
	SegmentIDList []string `json:"segment_id"`
	GPNID         string   `json:"gpn_id"`
	Mask          int      `json:"mask"`
	CreateStatus  int      `json:"create_status"`
	ID            string   `json:"id"`
	VlanID        int      `json:"vlan_id"`
	NetID         string   `json:"net_id"`
}
