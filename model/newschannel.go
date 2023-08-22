package model

import (
	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/officegraph"
	"github.com/koksmat-com/koksmat/officegraph/sites"
)

type NewsChannel struct {
	mgm.DefaultModel `bson:",inline"`
	Item             sites.NewsChannelsListItem `bson:",inline"`
}

func CreateNewsChannel(channel sites.NewsChannelsListItem) (newsChannel *NewsChannel, err error) {

	newRecord := &NewsChannel{}
	newRecord.Item = channel

	err = mgm.Coll(newRecord).Create(newRecord)

	return newRecord, err

}

func ImportNewsChannels() error {

	_, token, err := officegraph.GetClient()
	if err != nil {
		return err
	}

	got, err := sites.GetListItems[sites.NewsChannelsListItem](token, "sites/nexiintra-home", "News Channels", "")
	if err != nil {
		return err
	}

	for _, channel := range *got {

		_, err := CreateNewsChannel(channel)
		if err != nil {
			return err
		}

	}

	return nil
}
