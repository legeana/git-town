package dialog

import (
	"fmt"

	"github.com/git-town/git-town/v20/internal/cli/dialog/components"
	"github.com/git-town/git-town/v20/internal/cli/dialog/components/list"
	"github.com/git-town/git-town/v20/internal/config/configdomain"
	"github.com/git-town/git-town/v20/internal/messages"
)

const (
	syncPrototypeStrategyTitle = `Sync-prototype strategy`
	SyncPrototypeStrategyHelp  = `
Choose how Git Town should synchronize prototype branches.

Prototype branches are local-only feature branches.
They are useful for reducing load on CI systems
and limiting the sharing of confidential changes.

`
)

func SyncPrototypeStrategy(existing configdomain.SyncPrototypeStrategy, inputs components.TestInput) (configdomain.SyncPrototypeStrategy, bool, error) {
	entries := list.Entries[configdomain.SyncPrototypeStrategy]{
		{
			Data: configdomain.SyncPrototypeStrategyMerge,
			Text: "merge updates from the parent and tracking branch",
		},
		{
			Data: configdomain.SyncPrototypeStrategyRebase,
			Text: "rebase branches against their parent and tracking branch",
		},
		{
			Data: configdomain.SyncPrototypeStrategyCompress,
			Text: "compress the branch after merging parent and tracking",
		},
	}
	defaultPos := entries.IndexOf(existing)
	selection, aborted, err := components.RadioList(entries, defaultPos, syncPrototypeStrategyTitle, SyncPrototypeStrategyHelp, inputs)
	if err != nil || aborted {
		return configdomain.SyncPrototypeStrategyMerge, aborted, err
	}
	fmt.Printf(messages.SyncPrototypeBranches, components.FormattedSelection(selection.String(), aborted))
	return selection, aborted, err
}
