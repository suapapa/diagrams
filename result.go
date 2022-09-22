package diagrams

type Result struct {
	Msg  string `json:"msg,omitempty"`
	Err  string `json:"err,omitempty"`
	Name string `json:"name,omitempty"`
	Img  string `json:"img,omitempty"`
}
