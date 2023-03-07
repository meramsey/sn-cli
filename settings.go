package sncli

import (
	"github.com/meramsey/gosn-v2"
	"github.com/meramsey/gosn-v2/cache"
)

func (i *GetSettingsConfig) Run() (settings gosn.Items, err error) {
	getItemsInput := cache.SyncInput{
		Session: i.Session,
	}

	var so cache.SyncOutput

	so, err = Sync(getItemsInput, true)
	if err != nil {
		return nil, err
	}

	var allPersistedItems cache.Items

	err = so.DB.All(&allPersistedItems)
	if err != nil {
		return
	}

	var items gosn.Items

	items, err = allPersistedItems.ToItems(i.Session)
	if err != nil {
		return
	}

	items.Filter(i.Filters)

	return
}
