package stores

type Store struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Url         string `json:"url"`
	JSON        any    `json:"json"`
	IsCurrent   bool   `json:"isCurrent"`

	// Name of the connector
}
