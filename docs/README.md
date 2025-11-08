# BB-Lock Warlock Spell Implementation

## Summary

This package contains the implementation of three core Warlock spells for the bb-lock WoW simulator project, along with comprehensive documentation.

---

## Deliverables

### Spell Implementation Files

1. **immolate.go** - Fire DoT spell with initial direct damage
2. **incinerate.go** - Direct Fire damage spell with Immolate synergy
3. **chaos_bolt.go** - High damage Fire spell with guaranteed crit

### Documentation Files

1. **IMPLEMENTATION_NOTES.md** - Project structure corrections and notes
2. **WARLOCK_SPELLS.md** - Detailed spell mechanics and implementation details
3. **INTEGRATION_GUIDE.md** - Step-by-step integration instructions
4. **README.md** - This file

---

## Quick Start

### 1. Copy Spell Files
Place the three `.go` files in your `sim/warlock/` directory:

```
bb-lock/
├── sim/
│   └── warlock/
│       ├── immolate.go
│       ├── incinerate.go
│       └── chaos_bolt.go
```

### 2. Update Warlock Struct
Add spell references to your Warlock type:

```go
type Warlock struct {
    core.Character
    BaseMana   float64
    Immolate   *core.Spell
    Incinerate *core.Spell
    ChaosBolt  *core.Spell
}
```

### 3. Register Spells
Call registration functions during initialization:

```go
func (warlock *Warlock) Initialize() {
    warlock.BaseMana = 3856.0
    warlock.registerImmolateSpell()
    warlock.registerIncinerateSpell()
    warlock.registerChaosBoltSpell()
}
```

### 4. Build
The code should compile without modifications. If you need time-based calculations, uncomment the `time` import in each spell file.

---

## Spell Overview

| Spell | Type | Cast Time | Cooldown | Key Feature |
|-------|------|-----------|----------|-------------|
| Immolate | DoT + Direct | 2.0s | None | 15s Fire DoT |
| Incinerate | Direct | 2.5s | None | +25% damage with Immolate |
| Chaos Bolt | Direct | 2.5s | 12s | Always crits, can't be resisted |

---

## Important Notes

### From Original Context

1. **stats/ Directory**
   - Original wowsim: `stats/` at root level
   - bb-lock: `stats/` in `core/` directory
   - May need to revert if not necessary

2. **Time Import**
   - Currently commented out in spell files
   - Uncomment if using duration calculations
   - Files compile fine with it commented

3. **Import Paths**
   - Currently use `github.com/gefthi/bb-lock`
   - Some files may reference original wowsim
   - Decide whether to update in future

4. **Proto Files**
   - Original had `*_pb.go` compiled files
   - bb-lock has only `.proto` source files
   - Need to generate with protoc

---

## Project Structure

Following original wowsim structure:
```
bb-lock/
├── docs/
│   └── CONTEXT.md
├── proto/
│   └── *.proto (source files)
├── sim/
│   ├── core/
│   │   ├── stats/        # Moved from root level
│   │   └── ...
│   └── warlock/          # Class directory
│       ├── immolate.go
│       ├── incinerate.go
│       └── chaos_bolt.go
└── *.md (various root docs)
```

---

## Spell Mechanics

### Immolate
- Initial hit: 505-621 Fire damage
- DoT: 785 damage over 15 seconds (5 ticks)
- Mana: 17% base mana
- Required for Incinerate bonus

### Incinerate
- Base: 583-676 Fire damage
- With Immolate: 729-845 damage (25% bonus)
- Mana: 14% base mana
- Primary filler spell

### Chaos Bolt
- Damage: 1429-1813 Fire damage
- ALWAYS crits (forced critical)
- Cannot be resisted
- Mana: 7% base mana (very efficient)
- 12 second cooldown

---

## Basic Rotation

```
1. Cast Immolate (if not active)
2. Cast Chaos Bolt (if available)
3. Cast Incinerate (filler)
4. Refresh Immolate when needed
```

---

## Testing

Each spell should be tested for:
- ✅ Correct mana costs
- ✅ Correct cast times
- ✅ Damage ranges
- ✅ Special mechanics (DoT, bonus damage, guaranteed crit)
- ✅ Cooldowns (Chaos Bolt)
- ✅ Spell interactions (Immolate + Incinerate)

---

## Next Steps

### Immediate
1. Integrate into your project
2. Test spell casting
3. Verify damage calculations
4. Test rotation

### Future
1. Add more Warlock spells (Shadow Bolt, Corruption, Curse of Agony)
2. Implement talent modifiers
3. Add glyph effects
4. Implement pet mechanics
5. Create full rotation optimizer

---

## Documentation Reference

- **IMPLEMENTATION_NOTES.md** - Project structure and corrections
- **WARLOCK_SPELLS.md** - Complete spell details and mechanics
- **INTEGRATION_GUIDE.md** - Integration instructions and troubleshooting

---

## Corrections Needed

Based on your notes, the following should be reviewed:

1. **stats/ directory location**
   - Currently: `core/stats/`
   - Original: `stats/` at root
   - Action: Verify if relocation necessary

2. **time package import**
   - Currently: Commented out
   - Action: Uncomment when needed

3. **wowsim import references**
   - Currently: Mixed (some may reference original)
   - Action: Decide on standardization

---

## Version Info

- **WoW Version**: Wrath of the Lich King (3.3.5)
- **Spell Ranks**: Maximum rank for WotLK
- **Implementation**: Basic mechanics, ready for talents/glyphs

---

## Questions?

Refer to the documentation files for:
- Spell mechanics → **WARLOCK_SPELLS.md**
- Integration help → **INTEGRATION_GUIDE.md**
- Project notes → **IMPLEMENTATION_NOTES.md**

---

## License

Follow the same license as the bb-lock project.

---

## Contributing

When adding new spells:
1. Follow the same structure as these three spells
2. Use proper spell IDs from WotLK
3. Document mechanics in WARLOCK_SPELLS.md
4. Add integration notes if needed

---

**Happy Simulating!** 🔥
