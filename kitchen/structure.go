package kitchen

type Script struct {
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Environment []string `json:"environment"`
}

type Station struct {
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Path        string   `json:"path"`
	Description string   `json:"description"`
	Readme      string   `json:"readme"`
	Scripts     []Script `json:"scripts"`
}

type Kitchen struct {
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Stations    []Station `json:"stations"`
	Description string    `json:"description"`
	Path        string    `json:"path"`
	Readme      string    `json:"readme"`
}
