package warlock

// TODO: Uncomment when core package is ready
// import (
//     "github.com/gefthi/bb-lock/core"
// )

// Warlock represents a Warlock character in the simulation
type Warlock struct {
	// TODO: Embed core.Character
	// *core.Character
	
	// Spells will be registered here
	// ChaosBolt    *core.Spell
	// Immolate     *core.Spell
	// Conflagrate  *core.Spell
	// Incinerate   *core.Spell
	// Corruption   *core.Spell
	// LifeTap      *core.Spell
	
	// Talents
	Talents WarlockTalents
}

// WarlockTalents represents the warlock's talent tree selections
type WarlockTalents struct {
	// Destruction
	Bane             int32
	Aftermath        int32
	MoltenCore       int32
	Backlash         int32
	ChaosBolt        int32
	FireAndBrimstone int32
	Backdraft        int32
	Emberstorm       int32
	
	// Demonology
	DemonicEmbrace   int32
	
	// Affliction
	Suppression      int32
}

// New creates a new Warlock instance
// func New(char *core.Character) *Warlock {
// 	warlock := &Warlock{
// 		Character: char,
// 	}
// 	
// 	warlock.registerSpells()
// 	
// 	return warlock
// }

// registerSpells registers all warlock spells
// func (warlock *Warlock) registerSpells() {
// 	warlock.registerChaosBolt()
// 	warlock.registerImmolate()
// 	warlock.registerConflagrate()
// 	warlock.registerIncinerate()
// }

// OnGCDReady is called when the GCD is ready (rotation logic)
// Phase 1: Hardcoded simple rotation
// Phase 3: Will be replaced by APL system
// func (warlock *Warlock) OnGCDReady(sim *core.Simulation) {
// 	target := warlock.CurrentTarget
// 	
// 	// Simple priority:
// 	// 1. Immolate if not up
// 	// 2. Conflagrate if ready
// 	// 3. Chaos Bolt
// 	
// 	if !warlock.Immolate.Dot(target).IsActive() {
// 		if warlock.Immolate.Cast(sim, target) {
// 			return
// 		}
// 	}
// 	
// 	if warlock.Conflagrate.IsReady(sim) {
// 		if warlock.Conflagrate.Cast(sim, target) {
// 			return
// 		}
// 	}
// 	
// 	warlock.ChaosBolt.Cast(sim, target)
// }
