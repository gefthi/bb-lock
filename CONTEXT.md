# 🔮 BB-LOCK CONTEXT LOADER
**Optimized for Claude's Efficient Session Bootstrapping**

---

## ⚡ CRITICAL - READ THIS FIRST (Token-Optimized Entry Point)

**Repository:** https://github.com/gefthi/bb-lock  
**Branch:** main  
**Project Type:** WoWSims Warlock Subset (Private Server Simulator)  
**Current Phase:** Phase 1 - Initial Setup  
**Last Updated:** 2025-11-08

---

## 🎯 PROJECT SUMMARY (Quick Context)

### What This Is:
Standalone Warlock DPS simulator extracted from WoWSims (https://github.com/wowsims/wotlk) for a custom private server with:
- **Custom "Mystic Enchants"** - Server-specific mechanics (like glyphs)
- **YAML Configuration** - No web UI, config files only
- **APL System** - Action Priority Lists for rotations
- **Core WoWSims Engine** - Reused simulation core

### What We're Building:
A CLI tool that reads character stats/rotation from YAML, runs simulation, outputs DPS metrics.

---

## 📋 TOKEN ECONOMY DECISIONS (IMPORTANT!)

### ✅ Decisions Made to Save Tokens:

1. **Protobuf Strategy: Option 2 (Vendored)**
   - Copy generated `.pb.go` files from WoWSims
   - NO protoc compiler needed
   - NO build step required
   - Location: `proto/*.pb.go`

2. **Development Pattern: Diff/Patch After Bootstrap**
   - Phase 1: Generate complete files once (~10K tokens)
   - Phase 2+: Only generate diffs/patches (~100-300 tokens each)
   - Never re-paste full files in chat

3. **Context Loading Strategy:**
   - New session: Share this file's URL
   - Claude fetches this file (~2K tokens)
   - Claude knows repo structure from this doc
   - Claude only fetches specific files needed for task

4. **File Reference Method:**
   - Always use GitHub URLs to reference files
   - Claude caches within session (no re-fetch)
   - Example: "Update https://github.com/gefthi/bb-lock/blob/main/warlock/chaos_bolt.go"

---

## 📁 PROJECT STRUCTURE (Real Paths)

```
bb-lock/
├── main.go                           # Entry point - runs simulation
├── go.mod                            # Go module definition
├── go.sum                            # Dependency checksums
├── README.md                         # User-facing documentation
│
├── docs/                             # Documentation
│   ├── CONTEXT.md                    # This file! Context loader
│   ├── PROJECT-PLAN.md               # Full development plan
│   └── DECISIONS.md                  # Architectural decisions log
│
├── proto/                            # Vendored protobuf files (from WoWSims)
│   ├── api.pb.go                     # API definitions
│   ├── apl.pb.go                     # APL system definitions
│   ├── common.pb.go                  # Common types
│   └── warlock.pb.go                 # Warlock-specific (if exists)
│
├── core/                             # Core simulation engine (from WoWSims)
│   ├── sim.go                        # Main simulation loop & Step()
│   ├── spell.go                      # Spell casting system
│   ├── character.go                  # Character base class
│   ├── unit.go                       # Unit interface & base
│   ├── aura.go                       # Buff/debuff system
│   ├── dot.go                        # DOT mechanics
│   ├── stats.go                      # Stat system & calculations
│   ├── pending_action.go             # Event queue (priority queue)
│   ├── environment.go                # Raid environment setup
│   └── cast.go                       # Cast time & GCD handling
│
├── warlock/                          # Warlock-specific implementation
│   ├── warlock.go                    # Main Warlock struct & methods
│   ├── talents.go                    # Talent definitions & modifiers
│   ├── chaos_bolt.go                 # Chaos Bolt spell
│   ├── immolate.go                   # Immolate (direct + DOT)
│   ├── conflagrate.go                # Conflagrate spell
│   ├── incinerate.go                 # Incinerate filler
│   ├── corruption.go                 # Corruption DOT
│   ├── life_tap.go                   # Life Tap mana generation
│   ├── backdraft.go                  # Backdraft aura handling
│   └── rotation.go                   # Hardcoded rotation (Phase 1 only)
│
├── apl/                              # APL compiler (Phase 3)
│   ├── compiler.go                   # YAML → Proto converter
│   ├── parser.go                     # YAML parsing logic
│   └── conditions.go                 # Condition compilation
│
├── mystic/                           # Custom enchant system (Phase 4)
│   ├── enchant.go                    # Enchant interface & registry
│   ├── backdraft_chaos.go            # Backdraft + Chaos Bolt enchant
│   ├── immolate_haste.go             # Immolate haste scaling
│   └── conflag_no_consume.go         # Conflagrate doesn't consume Immolate
│
├── loader/                           # Configuration loaders (Phase 2)
│   ├── character.go                  # Load character.yml
│   ├── spells.go                     # Load spells.yml
│   ├── rotation.go                   # Load rotation.yml
│   └── enchants.go                   # Load mystic_enchants.yml
│
├── config/                           # User configuration files (Phase 2+)
│   ├── character.yml                 # Character stats, talents, gear
│   ├── spells.yml                    # Spell database (damage, coefs, etc)
│   ├── rotation.yml                  # APL in human-readable format
│   └── mystic_enchants.yml           # Custom enchant definitions
│
└── output/                           # Results formatting (Phase 4)
    ├── formatter.go                  # Format results for display
    └── exporter.go                   # Export to JSON/CSV
```

---

## 🔗 QUICK FILE ACCESS (Direct Links)

### Core Files (Fetch These First):
- **Main Entry:** https://github.com/gefthi/bb-lock/blob/main/main.go
- **Warlock Struct:** https://github.com/gefthi/bb-lock/blob/main/warlock/warlock.go
- **Sim Engine:** https://github.com/gefthi/bb-lock/blob/main/core/sim.go

### Spell Implementations:
- **Chaos Bolt:** https://github.com/gefthi/bb-lock/blob/main/warlock/chaos_bolt.go
- **Immolate:** https://github.com/gefthi/bb-lock/blob/main/warlock/immolate.go
- **Conflagrate:** https://github.com/gefthi/bb-lock/blob/main/warlock/conflagrate.go

### Configuration:
- **Character:** https://github.com/gefthi/bb-lock/blob/main/config/character.yml
- **Rotation:** https://github.com/gefthi/bb-lock/blob/main/config/rotation.yml

---

## 📊 PROJECT STATUS

### ✅ Completed:
- [x] Project plan created
- [x] Repository initialized
- [x] Context document created (this file)

### 🚧 In Progress:
- [ ] Phase 1: Bootstrap skeleton

### ⏳ Planned:
- [ ] Phase 2: YAML configuration system
- [ ] Phase 3: Human-readable APL compiler
- [ ] Phase 4: Mystic Enchants system

---

## 🎨 CODE PATTERNS & CONVENTIONS

### Go Package Structure:
```go
// Each spell in separate file
// warlock/chaos_bolt.go
package warlock

import (
    "time"
    "github.com/gefthi/bb-lock/core"
)

func (warlock *Warlock) registerChaosBolt() {
    warlock.ChaosBolt = warlock.RegisterSpell(core.SpellConfig{
        ActionID: core.ActionID{SpellID: 50796},
        // ... config
    })
}
```

### Naming Conventions:
- **Spells:** `ChaosBolt`, `Immolate` (PascalCase)
- **Files:** `chaos_bolt.go`, `immolate.go` (snake_case)
- **Configs:** `character.yml`, `rotation.yml` (lowercase)
- **Functions:** `registerChaosBolt()` (camelCase)

### Talent Access:
```go
// Always access via warlock.Talents struct
baseDamage *= 1.0 + 0.04 * float64(warlock.Talents.FireAndBrimstone)
```

### Spell Registration Pattern:
```go
func (warlock *Warlock) registerSpells() {
    warlock.registerChaosBolt()
    warlock.registerImmolate()
    warlock.registerConflagrate()
    // ... etc
}
```

---

## 🚀 COMMON TASKS & PATTERNS

### For Claude in New Sessions:

#### Task: Add New Spell
```
1. Fetch: warlock/warlock.go (to see spell registration pattern)
2. Generate: warlock/new_spell.go (complete file)
3. Generate: Patch for warlock.go (add registration call)
```

#### Task: Update Existing Spell
```
1. Fetch: warlock/spell_name.go
2. Generate: Diff/patch with changes
3. User applies patch locally
```

#### Task: Add Mystic Enchant (Phase 4)
```
1. Fetch: mystic/enchant.go (registry pattern)
2. Generate: mystic/new_enchant.go
3. Generate: Patch for registry
```

#### Task: Review/Debug
```
1. User provides GitHub URL to problematic file
2. Claude fetches & analyzes
3. Claude suggests fixes as diff/patch
```

---

## 🧠 CONTEXT LOADING PROTOCOL

### Start of New Session - Optimal Pattern:

**User Says:**
```
Context: https://github.com/gefthi/bb-lock/blob/main/docs/CONTEXT.md

Task: [specific request]
```

**Claude Does:**
1. Fetch CONTEXT.md (~2K tokens) - gets full project understanding
2. Identify which files needed for task
3. Fetch ONLY those files (~2-4K tokens)
4. Generate response (diff/patch or new file)

**Total:** ~4-6K tokens (vs ~15K if explaining from scratch)

---

## 📖 REFERENCE DOCUMENTATION

### Key WoWSims Files for Reference:
- **Core Engine:** https://github.com/wowsims/wotlk/blob/master/sim/core/sim.go
- **Spell System:** https://github.com/wowsims/wotlk/blob/master/sim/core/spell.go
- **APL System:** https://github.com/wowsims/wotlk/blob/master/sim/core/apl.go

### Private Server Info:
- **Server:** Custom WotLK private server
- **Special Feature:** Mystic Enchants (glyph-like system)
- **Examples:**
  - Backdraft + Chaos Bolt: CB during Backdraft doesn't consume charges for 4s
  - Immolate haste: DOT ticks affected by haste (normally aren't)
  - Conflagrate: Doesn't consume Immolate

---

## 🔧 DEVELOPMENT PHASES

### Phase 1: Bootstrap (Week 1) - Current
**Goal:** Working sim with hardcoded rotation
- Generate skeleton files
- Implement 3-5 basic spells
- Hardcoded simple rotation
- Output DPS

**Files to Generate:**
- main.go
- go.mod
- warlock/warlock.go
- warlock/chaos_bolt.go
- warlock/immolate.go
- warlock/conflagrate.go
- README.md

### Phase 2: YAML Configs (Week 2)
**Goal:** Config-driven instead of hardcoded
- character.yml (stats, talents, glyphs)
- spells.yml (spell database)
- Loader infrastructure

### Phase 3: APL Compiler (Week 3)
**Goal:** Human-readable rotations
- rotation.yml (YAML APL format)
- Compiler (YAML → Proto)
- Integration with sim

### Phase 4: Mystic Enchants (Week 4)
**Goal:** Custom server mechanics
- mystic_enchants.yml
- Enchant system
- 5+ example enchants

---

## 🐛 KNOWN ISSUES & GOTCHAS

### Current Known Issues:
- None yet (Phase 1 not started)

### Potential Gotchas:
1. **Proto imports:** Must update import paths after vendoring
   - Change: `github.com/wowsims/wotlk/sim/core/proto`
   - To: `github.com/gefthi/bb-lock/proto`

2. **WoWSims dependencies:** Some core files import other core files
   - Must copy dependency tree correctly
   - Solution: Start with minimal set, add as needed

3. **Spell interactions:** Complex interactions between spells/auras
   - Example: Backdraft consumed by Incinerate/Chaos Bolt
   - Solution: Implement step-by-step, test thoroughly

---

## 💡 TIPS FOR EFFICIENT COLLABORATION

### For User:
1. **Always link to this CONTEXT.md at start of new sessions**
2. **Reference files by GitHub URL**
3. **Ask for diffs/patches for updates (not full rewrites)**
4. **Batch related questions together**
5. **Update this file when project state changes**

### For Claude:
1. **Fetch CONTEXT.md first in new sessions**
2. **Only fetch files relevant to current task**
3. **Generate diffs/patches after Phase 1**
4. **Cache fetched files within session**
5. **Reference back to this doc for patterns/conventions**

---

## 📝 VERSION HISTORY

### v1.0 - 2025-11-08
- Initial context document created
- Project structure defined
- Token optimization strategies documented
- Ready for Phase 1 implementation

---

## 🎯 QUICK REFERENCE - SPELL IDS

### Destruction Spells:
- **Chaos Bolt:** 50796 (Rank 4)
- **Immolate:** 47811 (Rank 11)
- **Conflagrate:** 17962 (Rank 8)
- **Incinerate:** 47838 (Rank 4)

### Affliction Spells:
- **Corruption:** 47813 (Rank 10)
- **Unstable Affliction:** 47843 (Rank 5)
- **Haunt:** 59164 (Rank 4)

### Utility:
- **Life Tap:** 57946 (Rank 8)

### Auras:
- **Backdraft:** 54274 (3 stacks, 30% haste)
- **Decimation:** 63167 (execute phase buff)

---

## 🔮 FUTURE ENHANCEMENTS (Post-Phase 4)

- Multi-target support
- Pandemic mechanics
- Stat weight calculator
- Rotation optimizer
- Comparison tool (different gear sets)
- Web UI (optional, later)

---

## 📞 CONTACT & REPO INFO

**Repository:** https://github.com/gefthi/bb-lock  
**Issues:** https://github.com/gefthi/bb-lock/issues  
**License:** MIT (inherited from WoWSims)

**Attribution:** Based on WoWSims (https://github.com/wowsims/wotlk)

---

## ⚡ TLDR FOR NEW SESSIONS

```
Repo: https://github.com/gefthi/bb-lock
Type: Warlock DPS sim (WoWSims subset)
Phase: 1 (Bootstrap)
Strategy: Vendored protos, diff-based updates
Pattern: CONTEXT.md → fetch needed files → generate diffs
Goal: YAML-configured sim with custom enchants
```

---

**END OF CONTEXT DOCUMENT**

*This file is designed to be the SINGLE source of truth for bootstrapping Claude in new sessions with maximum token efficiency.*
