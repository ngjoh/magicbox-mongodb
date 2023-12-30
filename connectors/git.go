package connectors

import (
	"strings"
)

func GitStatus() ([]Connector, error) {
	mateContext, err := GetMateContext()
	if err != nil {
		return nil, err
	}
	bytes, err := Execute("git", *&Options{Dir: mateContext.Current.Path}, "status")
	if err != nil {
		return nil, err
	}
	connectors := make([]Connector, 0)
	orgs := strings.Split(string(bytes), "\n")
	for _, org := range orgs {
		name := strings.TrimSpace(org)
		if name == "" {
			continue
		}
		connector := Connector{
			Name: name,
			Url:  "https://github.com/" + name,
		}
		connectors = append(connectors, connector)
	}

	return connectors, nil
}
