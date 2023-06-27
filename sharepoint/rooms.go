package sharepoint

type Room struct {
	Canbeusedforreceptions any    `json:"Canbeusedforreceptions"`
	Capacity               int    `json:"Capacity"`
	CiscoVideo             any    `json:"CiscoVideo"`
	DeviceSerialNumber     any    `json:"DeviceSerialNumber"`
	Email                  string `json:"Email"`
	ID                     int    `json:"ID"`
	ID0                    int    `json:"Id"`
	PriceList              struct {
		Deliverto any `json:"Deliverto"`
		Metadata  struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"__metadata"`
	} `json:"Price_x0020_List"`
	Production         any    `json:"Production"`
	ProvisioningStatus string `json:"Provisioning_x0020_Status"`
	RestrictedTo       string `json:"RestrictedTo"`
	TeamsMeetingRoom   any    `json:"TeamsMeetingRoom"`
	Title              string `json:"Title"`
	Metadata           struct {
		Etag string `json:"etag"`
		ID   string `json:"id"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"__metadata"`
}

type Floor struct {
	ID    int `json:"ID"`
	ID0   int `json:"Id"`
	Rooms struct {
		Results []struct {
			ID       int    `json:"Id"`
			Title    string `json:"Title"`
			Metadata struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"__metadata"`
		} `json:"results"`
	} `json:"Rooms"`
	Title    string `json:"Title"`
	Metadata struct {
		Etag string `json:"etag"`
		ID   string `json:"id"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"__metadata"`
}

type Building struct {
	Floors struct {
		Results []struct {
			ID       int    `json:"Id"`
			Title    string `json:"Title"`
			Metadata struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"__metadata"`
		} `json:"results"`
	} `json:"Floors"`
	ID         int    `json:"ID"`
	ID0        int    `json:"Id"`
	Title      string `json:"Title"`
	Wheelchair bool   `json:"Wheelchair"`
	Metadata   struct {
		Etag string `json:"etag"`
		ID   string `json:"id"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"__metadata"`
}

type Location struct {
	Buildings struct {
		Results []struct {
			ID       int    `json:"Id"`
			Title    string `json:"Title"`
			Metadata struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"__metadata"`
		} `json:"results"`
	} `json:"Buildings"`
	ID       int    `json:"ID"`
	ID0      int    `json:"Id"`
	Title    string `json:"Title"`
	Metadata struct {
		Etag string `json:"etag"`
		ID   string `json:"id"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"__metadata"`
}

type Country struct {
	Countrycode string `json:"Countrycode"`
	ID          int    `json:"ID"`
	ID0         int    `json:"Id"`
	Locations   struct {
		Results []struct {
			ID       int    `json:"Id"`
			Title    string `json:"Title"`
			Metadata struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"__metadata"`
		} `json:"results"`
	} `json:"Locations"`
	Title    string `json:"Title"`
	Metadata struct {
		Etag string `json:"etag"`
		ID   string `json:"id"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"__metadata"`
}

func RoomsList() (*[]Room, error) {
	items, err := GetListItems[Room](
		"https://christianiabpos.sharepoint.com/sites/Cava3",
		"Rooms",
		"Id,Title,Capacity,Provisioning_x0020_Status,Email,RestrictedTo,TeamsMeetingRoom,Canbeusedforreceptions,DeviceSerialNumber,Price_x0020_List/Deliverto,CiscoVideo,Production",
		"Price_x0020_List/Id")

	if err != nil {
		return nil, err
	}
	return &items, nil
}

func FloorsList() (*[]Floor, error) {
	items, err := GetListItems[Floor](
		"https://christianiabpos.sharepoint.com/sites/Cava3",
		"Floors",
		"Id,Title,Rooms/Id,Rooms/Title",
		"Rooms/Id")

	if err != nil {
		return nil, err
	}
	return &items, nil
}

func BuildingsList() (*[]Building, error) {
	items, err := GetListItems[Building](
		"https://christianiabpos.sharepoint.com/sites/Cava3",
		"Buildings",
		"Id,Title,Wheelchair,Floors/Id,Floors/Title",
		"Floors/Id")

	if err != nil {
		return nil, err
	}
	return &items, nil
}
func LocationsList() (*[]Location, error) {
	items, err := GetListItems[Location](
		"https://christianiabpos.sharepoint.com/sites/Cava3",
		"Locations",
		"Id,Title,Buildings/Id,Buildings/Title",
		"Buildings/Id")

	if err != nil {
		return nil, err
	}
	return &items, nil
}
func CountriesList() (*[]Country, error) {
	items, err := GetListItems[Country](
		"https://christianiabpos.sharepoint.com/sites/Cava3",
		"Countries",
		"Id,Title,Countrycode,Locations/Id,Locations/Title",
		"Locations/Id")

	if err != nil {
		return nil, err
	}
	return &items, nil
}
