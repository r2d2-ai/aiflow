package action

import (
	"fmt"

	"github.com/r2d2-ai/aiflow/support/log"

	"github.com/r2d2-ai/aiflow/support"
)

var (
	actionFactories = make(map[string]Factory)
)

func Register(action Action, f Factory) error {

	if action == nil {
		return fmt.Errorf("'action' must be specified when registering")
	}

	if f == nil {
		return fmt.Errorf("cannot register action with 'nil' action factory")
	}

	ref := support.GetRef(action)

	if _, dup := actionFactories[ref]; dup {
		return fmt.Errorf("action already registered: %s", ref)
	}

	log.RootLogger().Debugf("Registering action: %s", ref)

	actionFactories[ref] = f

	return nil
}

func GetFactory(ref string) Factory {
	return actionFactories[ref]
}

func Factories() map[string]Factory {
	//todo return copy
	return actionFactories
}
