package types

type AntlrColumn struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	Length        int    `json:"length,omitempty"`
	Scale         int    `json:"scale,omitempty"`
	Comment       string `json:"comment,omitempty"`
	AutoIncrement bool   `json:"auto_increment,omitempty"`
}

type AntlrTable struct {
	Database string         `json:"database"`
	Name     string         `json:"name"`
	Columns  []*AntlrColumn `json:"columns"`
	Comment  string         `json:"comment,omitempty"`
}
