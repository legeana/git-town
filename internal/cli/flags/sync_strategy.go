package flags

import (
	"github.com/git-town/git-town/v17/internal/config/configdomain"
	. "github.com/git-town/git-town/v17/pkg/prelude"
	"github.com/spf13/cobra"
)

const syncFeatureStrategyLong = "feature-strategy"

// type-safe access to the CLI arguments of type configdomain.SyncFeatureStrategy
func SyncFeatureStrategy() (AddFunc, ReadSyncFeatureStrategyFunc) {
	addFlag := func(cmd *cobra.Command) {
		cmd.Flags().StringP(syncFeatureStrategyLong, "", "", "override the sync-feature-strategy")
	}
	readFlag := func(cmd *cobra.Command) (Option[configdomain.SyncFeatureStrategy], error) {
		value, err := cmd.Flags().GetString(syncFeatureStrategyLong)
		if err != nil {
			return None[configdomain.SyncFeatureStrategy](), err
		}
		strategyOpt, err := configdomain.ParseSyncFeatureStrategy(value)
		if err != nil {
			return None[configdomain.SyncFeatureStrategy](), err
		}
		strategy, hasStrategy := strategyOpt.Get()
		if !hasStrategy {
			return None[configdomain.SyncFeatureStrategy](), nil
		}
		return Some(strategy), nil
	}
	return addFlag, readFlag
}

// ReadSyncStrategyFunc is the type signature for the function that reads the sync "feature-strategy" flag from the args to the given Cobra command.
type ReadSyncStrategyFunc func(*cobra.Command) (Option[configdomain.SyncFeatureStrategy], error)
