package journeys

type Cargo struct {
	Name  string            `json:"name" bson:"name"`
	Goods map[string]string `json:"goods" bson:"goods"`
}
