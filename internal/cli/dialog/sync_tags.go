package dialog

import (
	"fmt"

	"github.com/git-town/git-town/v20/internal/cli/dialog/components"
	"github.com/git-town/git-town/v20/internal/cli/dialog/components/list"
	"github.com/git-town/git-town/v20/internal/cli/format"
	"github.com/git-town/git-town/v20/internal/config/configdomain"
	"github.com/git-town/git-town/v20/internal/messages"
)

const (
	syncTagsTitle = `Sync-tags strategy`
	SyncTagsHelp  = `
Should "git town sync" sync Git tags with origin?

`
)

func SyncTags(existing configdomain.SyncTags, inputs components.TestInput) (configdomain.SyncTags, bool, error) {
	entries := list.Entries[configdomain.SyncTags]{
		{
			Data: true,
			Text: "yes, sync Git tags",
		},
		{
			Data: false,
			Text: "no, don't sync Git tags",
		},
	}
	defaultPos := entries.IndexOf(existing)
	selection, aborted, err := components.RadioList(entries, defaultPos, syncTagsTitle, SyncTagsHelp, inputs)
	if err != nil || aborted {
		return true, aborted, err
	}
	fmt.Printf(messages.SyncTags, components.FormattedSelection(format.Bool(selection.IsTrue()), aborted))
	return selection, aborted, err
}
