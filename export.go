package sncli

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/meramsey/gosn-v2"
	"github.com/meramsey/gosn-v2/cache"
	"os"
	"time"
)

type ExportConfig struct {
	Session   *cache.Session
	Decrypted bool
	File      string
	UseStdOut bool
}

type ImportConfig struct {
	Session   *cache.Session
	File      string
	Format    string
	UseStdOut bool
	Debug     bool
}

// Run will retrieve all items from SN directly, re-encrypt them with a new ItemsKey and write them to a file.
func (i ExportConfig) Run() error {
	if !i.Session.Debug {
		prefix := HiWhite("exporting ")

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithWriter(os.Stdout))
		if i.UseStdOut {
			s = spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
		}

		s.Prefix = prefix
		s.Start()

		err := i.Session.Export(i.File)

		s.Stop()

		return err
	}

	return i.Session.Export(i.File)
}

func (i *ImportConfig) Run() (imported int, err error) {
	// populate DB
	gii := cache.SyncInput{
		Session: i.Session,
	}

	gio, err := Sync(gii, true)
	if err != nil {
		return imported, err
	}

	var syncTokens []cache.SyncToken
	if err = gio.DB.All(&syncTokens); err != nil {
		return imported, err
	}
	syncToken := ""
	if len(syncTokens) > 0 {
		syncToken = syncTokens[0].SyncToken
	}
	if err = gio.DB.Close(); err != nil {
		return imported, err
	}

	// request all items from SN
	var iItemsKey gosn.ItemsKey
	var iItems gosn.EncryptedItems

	iItems, iItemsKey, err = i.Session.Session.Import(i.File, syncToken, "")
	if err != nil {
		return 0, err
	}

	if iItemsKey.ItemsKey == "" {
		panic(fmt.Sprintf("iItemsKey.ItemsKey is empty for: '%s'", iItemsKey.UUID))
	}

	// push item and close db
	pii := cache.SyncInput{
		Session: i.Session,
		Close:   true,
	}

	_, err = Sync(pii, true)
	imported = len(iItems)

	return imported, err
}
