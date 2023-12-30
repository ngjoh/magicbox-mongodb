package kitchen

type Station struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Readme      string `json:"readme"`
}

type Kitchen struct {
	Name        string    `json:"name"`
	Stations    []Station `json:"stations"`
	Description string    `json:"description"`
	Path        string    `json:"path"`
	Readme      string    `json:"readme"`
}
