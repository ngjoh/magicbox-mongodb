package kitchen

type Script struct {
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Environment []string `json:"environment"`
	Connection  string   `json:"connection"`
	Input       string   `json:"input"`
	Output      string   `json:"output"`
	Cron        string   `json:"cron"`
	Tag         string   `json:"tag"`
	Trigger     string   `json:"trigger"`
}

type Station struct {
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Path        string   `json:"path"`
	Description string   `json:"description"`
	Readme      string   `json:"readme"`
	Scripts     []Script `json:"scripts"`
	Tag         string   `json:"tag"`
}

type Kitchen struct {
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Stations    []Station `json:"stations"`
	Description string    `json:"description"`
	Path        string    `json:"path"`
	Readme      string    `json:"readme"`
	Tag         string    `json:"tag"`
}
