# Core Simulation Engine

This directory will contain files copied from WoWSims core engine.

## Files to Copy from WoWSims

From `https://github.com/wowsims/wotlk/tree/master/sim/core`:

### Essential Files:
- [ ] `sim.go` - Main simulation loop
- [ ] `spell.go` - Spell casting system
- [ ] `character.go` - Character base class
- [ ] `unit.go` - Unit interface
- [ ] `aura.go` - Buff/debuff system
- [ ] `dot.go` - DOT mechanics
- [ ] `stats.go` - Stat calculations
- [ ] `pending_action.go` - Event queue
- [ ] `environment.go` - Raid environment
- [ ] `cast.go` - Cast time & GCD

### Import Path Updates Required:
After copying, update all imports:
- FROM: `github.com/wowsims/wotlk/sim/core`
- TO: `github.com/gefthi/bb-lock/core`

## Instructions

1. Clone WoWSims: `git clone https://github.com/wowsims/wotlk.git`
2. Copy files from `wotlk/sim/core/` to this directory
3. Update import statements in copied files
4. Test compilation: `go build`

## Note

These files are copied under MIT license from WoWSims project.
Attribution preserved in LICENSE file.
