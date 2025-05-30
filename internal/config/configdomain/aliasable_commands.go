package configdomain

import . "github.com/git-town/git-town/v20/pkg/prelude"

type AliasableCommands []AliasableCommand

// provides the AliasKey matching the given key name
func (self AliasableCommands) LookupKey(name string) Option[AliasKey] {
	for _, aliasableCommand := range self {
		keyOfCommand := aliasableCommand.Key()
		if keyOfCommand.String() == name {
			return Some(keyOfCommand)
		}
	}
	return None[AliasKey]()
}

func (self AliasableCommands) Strings() []string {
	result := make([]string, len(self))
	for c, command := range self {
		result[c] = command.String()
	}
	return result
}

// AllAliasableCommands provides all AliasType values.
func AllAliasableCommands() AliasableCommands {
	return AliasableCommands{
		AliasableCommandAppend,
		AliasableCommandCompress,
		AliasableCommandContribute,
		AliasableCommandDiffParent,
		AliasableCommandHack,
		AliasableCommandDelete,
		AliasableCommandObserve,
		AliasableCommandPark,
		AliasableCommandPrepend,
		AliasableCommandPropose,
		AliasableCommandRename,
		AliasableCommandRepo,
		AliasableCommandSetParent,
		AliasableCommandShip,
		AliasableCommandSync,
	}
}
