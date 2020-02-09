package gic

type OSList struct {
	Status  string       `json:"status"`
	CodeMsg string       `json:"code_msg"`
	Message string       `json:"message"`
	Code    int          `json:"code"`
	Data    []OSListData `json:"data"`
}

type OSListData struct {
	DisplayName  string `json:"display_name"`
	Name         string `json:"name"`
	TemplateType string `json:"template_type"`
	OSType       string `json:"os_type"`
	Type         string `json:"type"`
	ID           string `json:"id"`
}
