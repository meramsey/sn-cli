package sncli

import (
	"github.com/meramsey/gosn-v2"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/meramsey/gosn-v2/cache"
)

func TestAddDeleteTagByTitle(t *testing.T) {
	testDelay()

	addTagConfig := AddTagsInput{
		Session: testSession,
		Tags:    []string{"TestTagOne", "TestTagTwo"},
	}

	ato, err := addTagConfig.Run()
	require.NoError(t, err)
	require.Contains(t, ato.Added, "TestTagOne")
	require.Contains(t, ato.Added, "TestTagTwo")
	require.Empty(t, ato.Existing)

	deleteTagConfig := DeleteTagConfig{
		Session:   testSession,
		TagTitles: []string{"TestTagOne", "TestTagTwo"},
	}

	var noDeleted int
	noDeleted, err = deleteTagConfig.Run()
	require.Equal(t, 2, noDeleted)
	require.NoError(t, err)
}

func TestGetTag(t *testing.T) {
	testDelay()

	defer cleanUp(*testSession)

	testTagTitles := []string{"TestTagOne", "TestTagTwo"}
	addTagInput := AddTagsInput{
		Session: testSession,
		Tags:    testTagTitles,
	}

	ato, err := addTagInput.Run()
	require.NoError(t, err)
	require.NoError(t, err)
	require.Contains(t, ato.Added, "TestTagOne")
	require.Contains(t, ato.Added, "TestTagTwo")
	require.Empty(t, ato.Existing)

	// create filters
	getTagFilters := gosn.ItemFilters{
		MatchAny: true,
	}

	for _, testTagTitle := range testTagTitles {
		getTagFilters.Filters = append(getTagFilters.Filters, gosn.Filter{
			Key:        "Title",
			Value:      testTagTitle,
			Type:       "Tag",
			Comparison: "==",
		})
	}

	getTagConfig := GetTagConfig{
		Session: testSession,
		Filters: getTagFilters,
	}

	var output gosn.Items
	output, err = getTagConfig.Run()
	require.NoError(t, err)
	require.EqualValues(t, len(output), 2, "expected two items but got: %+v", output)
}

func _addNotes(session cache.Session, i map[string]string) error {
	for k, v := range i {
		addNoteConfig := AddNoteInput{
			Session: &session,
			Title:   k,
			Text:    v,
		}

		err := addNoteConfig.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func _deleteNotesByTitle(session cache.Session, input map[string]string) (noDeleted int, err error) {
	var noteTitles []string
	for k := range input {
		noteTitles = append(noteTitles, k)
	}
	deleteNoteConfig := DeleteNoteConfig{
		Session:    &session,
		NoteTitles: noteTitles,
	}

	noDeleted, err = deleteNoteConfig.Run()
	if err != nil {
		return noDeleted, err
	}

	return noDeleted, deleteNoteConfig.Session.CacheDB.Close()
}

func _deleteTagsByTitle(session cache.Session, input []string) (noDeleted int, err error) {
	deleteTagConfig := DeleteTagConfig{
		Session:   &session,
		TagTitles: input,
	}

	return deleteTagConfig.Run()
}

func TestTaggingOfNotes(t *testing.T) {
	testDelay()
	defer cleanUp(*testSession)

	// create four notes
	notes := map[string]string{
		"noteOneTitle":   "noteOneText example",
		"noteTwoTitle":   "noteTwoText",
		"noteThreeTitle": "noteThreeText",
		"noteFourTitle":  "noteFourText example",
	}

	err := _addNotes(*testSession, notes)
	require.NoError(t, err)
	// tag new notes with 'testTag'
	tags := []string{"testTag"}
	tni := TagItemsConfig{
		Session:  testSession,
		FindText: "example",
		NewTags:  tags,
	}
	err = tni.Run()
	require.NoError(t, err)

	filterNotesByTagName := gosn.Filter{
		Type:       "Note",
		Key:        "TagTitle",
		Comparison: "==",
		Value:      "testTag",
	}
	itemFilters := gosn.ItemFilters{
		Filters:  []gosn.Filter{filterNotesByTagName},
		MatchAny: true,
	}
	gnc := GetNoteConfig{
		Session: testSession,
		Filters: itemFilters,
	}

	var retNotes gosn.Items
	retNotes, err = gnc.Run()
	require.NoError(t, err)

	if len(retNotes) != 2 {
		t.Errorf("expected two notes but got: %d", len(retNotes))
	}
	require.NoError(t, testSession.CacheDB.Close())

	nd, err := _deleteNotesByTitle(*testSession, notes)
	require.NoError(t, err)
	require.Equal(t, 4, nd)

	var deletedTags int
	deletedTags, err = _deleteTagsByTitle(*testSession, tags)
	require.NoError(t, err)
	require.Equal(t, 1, deletedTags)
}
