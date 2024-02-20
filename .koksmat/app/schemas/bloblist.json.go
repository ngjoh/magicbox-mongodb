package schemas

type Bloblist []struct {
	LastModified string `json:"lastModified"`
	Name         string `json:"name"`
	Size         int    `json:"size"`
}

