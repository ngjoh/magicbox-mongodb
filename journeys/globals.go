package journeys

import "strings"

func getEntity(entityString string) (string, string) {
	s := strings.Split(entityString, ",")
	e := strings.TrimSpace(s[0])
	es := strings.Split(e, ".")
	if (len(es)) > 1 {
		return strings.TrimSpace(es[0]), es[1]
	}

	return e, ""

}

func journeyKey(journeyName string, id string) string {
	return "@" + journeyName + ":" + id
}
