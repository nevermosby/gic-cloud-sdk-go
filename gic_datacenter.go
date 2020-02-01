package gic

// "status": "success",
//     "code_msg": "success",
//     "message": "查询app信息成功",
//     "code": 200,
type DataCenterData struct {
	Status string `json:"status"`
	CodeMsg string `json:"code_msg"`
	Message string `json:"message"`
	Code int `json:"code"`
	Data []DataCenter `json:"data"`
}
// for gic datacenter
// "resource": {
// 	"name": "tkvdc",
// 	"gic_count": 0,
// 	"wan_count": 1,
// 	"lan_count": 1,
// 	"id": "ad7ed58b-5011-4645-acf6-cf68623b92e4",
// 	"vm_count": 0
// },
// "name": "tkvdc",
//             "customer_user": "13917107484",
//             "site_id": "31f105b5-389e-4989-9944-8ecf76e9d764",
//             "app_id": "ad7ed58b-5011-4645-acf6-cf68623b92e4",
//             "site_name": "东京1"
type DataCenter struct {
	Resource     DataCenterResource `json:"resource"`
	Name         string             `json:"name"`
	CustomerUser string             `json:"customer_user"`
	SiteID       string             `json:"site_id"`
	AppID        string             `json:"app_id"`
	SiteName     string             `json:"site_name"`
}

// for gic datacenter
// "resource": {
// 	"name": "tkvdc",
// 	"gic_count": 0,
// 	"wan_count": 1,
// 	"lan_count": 1,
// 	"id": "ad7ed58b-5011-4645-acf6-cf68623b92e4",
// 	"vm_count": 0
// },
type DataCenterResource struct {
	Name     string `json:"name"`
	GicCount int    `json:"gic_count"`
	WanCount int    `json:"wan_count"`
	LanCount int    `json:"lan_count"`
	ID       string `json:"id"`
	VMCount  int    `json:"vm_count"`
}
