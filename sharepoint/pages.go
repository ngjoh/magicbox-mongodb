package sharepoint

func Pages(sharePointSiteUrl string) (interface{}, error) {
	_, err := GetClient(sharePointSiteUrl)

	if err != nil {
		return nil, err
	}

	return nil, nil

}
