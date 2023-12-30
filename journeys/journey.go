package journeys

type Container struct {
	Container any      `json:"container"`
	Name      string   `json:"name"`
	Key       string   `json:"key"`
	Who       []string `json:"who"`
	Approve   []string `json:"approve,omitempty"`
	Consult   []string `json:"consult,omitempty"`
	Inform    []string `json:"inform,omitempty"`
	Needs     []string `json:"needs"`
	Produces  []string `json:"produces"`
	Script    string   `json:"script"`
}

type Waypoint struct {
	Port  string   `json:"port"`
	Done  []string `json:"done,omitempty"`
	Loads struct {
		Containers []Container `json:"containers"`
	} `json:"loads"`
	Services []struct {
		Tugs []struct {
			Tug      any      `json:"tug"`
			Name     string   `json:"name"`
			Who      []string `json:"who"`
			Needs    []string `json:"needs"`
			Produces []string `json:"produces"`
			Script   string   `json:"script"`
		} `json:"tugs"`
	} `json:"services,omitempty"`
}
type Journey struct {
	Journey  string `json:"journey"`
	Metadata struct {
		App         string `json:"app"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"metadata"`
	Triggers []struct {
		Trigger any      `json:"trigger"`
		Name    string   `json:"name"`
		Key     string   `json:"key"`
		Details []string `json:"details"`
	} `json:"triggers"`
	Waypoints []Waypoint `json:"waypoints"`
}

type TravelPlan struct {
	Journey  string `json:"journey"`
	Metadata struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"metadata"`
	Triggers []struct {
		Trigger any      `json:"trigger"`
		Name    string   `json:"name"`
		Needs   []string `json:"needs"`
	} `json:"triggers"`
	Waypoints []struct {
		Port  string   `json:"port"`
		Done  []string `json:"done,omitempty"`
		Loads struct {
			Containers []struct {
				Container any      `json:"container"`
				Name      string   `json:"name"`
				Who       []string `json:"who"`
				Approve   []string `json:"approve,omitempty"`
				Consult   []string `json:"consult,omitempty"`
				Inform    []string `json:"inform,omitempty"`
				Needs     []string `json:"needs"`
				Produces  []string `json:"produces"`
				Powerapp  struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"powerapp,omitempty"`
				Script string `json:"script"`
			} `json:"containers"`
		} `json:"loads"`
	} `json:"waypoints"`
}

type Entity struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Attributes []string `json:"attributes"`
}

type ContainerDetails struct {
	Container any      `json:"container"`
	Name      string   `json:"name"`
	Who       []string `json:"who"`
	Approve   []string `json:"approve,omitempty"`
	Consult   []string `json:"consult,omitempty"`
	Inform    []string `json:"inform,omitempty"`
	Needs     []Entity `json:"needs"`
	Produces  []Entity `json:"produces"`
	Powerapp  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"powerapp,omitempty"`
	Script string `json:"script"`
}
type Loads struct {
	Containers []ContainerDetails `json:"containers"`
}
type WaypointDetails struct {
	Port  string   `json:"port"`
	Done  []string `json:"done,omitempty"`
	Loads Loads    `json:"loads"`
}
type DetailedJourney struct {
	Journey  string `json:"journey"`
	Metadata struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"metadata"`
	Triggers []struct {
		Trigger any      `json:"trigger"`
		Name    string   `json:"name"`
		Needs   []string `json:"needs"`
	} `json:"triggers"`
	Waypoints []WaypointDetails `json:"waypoints"`
}
