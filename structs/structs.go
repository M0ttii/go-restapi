package structs

type Video struct {
	Name     string `json:"name"`
	Uploaded string `json:"uploaded"`
	Show     bool   `json:"show"`
}

func (v Video) GetName() string {
	return v.Name
}

func (v Video) GetUploaded() string {
	return v.Uploaded
}

func (v Video) GetShow() bool {
	return v.Show
}
