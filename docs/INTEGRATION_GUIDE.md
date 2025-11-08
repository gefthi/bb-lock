# Integration Guide for Warlock Spells

## Quick Start

This guide explains how to integrate the three implemented Warlock spells into your bb-lock project.

---

## File Placement

Place the spell implementation files in your project:

```
bb-lock/
├── sim/
│   └── warlock/
│       ├── immolate.go
│       ├── incinerate.go
│       └── chaos_bolt.go
```

---

## Warlock Struct Requirements

Ensure your `warlock.go` file has the necessary fields:

```go
package warlock

import (
    "github.com/gefthi/bb-lock/sim/core"
)

type Warlock struct {
    core.Character
    
    BaseMana float64
    
    // Spell references
    Immolate   *core.Spell
    Incinerate *core.Spell
    ChaosBolt  *core.Spell
    
    // Add other fields as needed
}
```

---

## Spell Registration

In your Warlock initialization function, register the spells:

```go
func (warlock *Warlock) Initialize() {
    // Initialize base character
    warlock.Character.Initialize()
    
    // Set base mana (adjust based on level/stats)
    warlock.BaseMana = 3856.0 // Level 80 base mana
    
    // Register spells
    warlock.registerImmolateSpell()
    warlock.registerIncinerateSpell()
    warlock.registerChaosBoltSpell()
    
    // Register other spells...
}
```

---

## Time Package Usage

If you need to use time-based calculations, uncomment the time import in each spell file:

```go
import (
    "time" // Uncomment when needed
    "github.com/gefthi/bb-lock/sim/core"
)
```

Then you can use time.Duration for cooldowns and durations:

```go
Duration: 12 * time.Second,  // Instead of 12000 milliseconds
```

---

## Import Path Updates

### Current State
Files currently import from:
```go
"github.com/gefthi/bb-lock/sim/core"
```

### If You Need to Change
If you need to update import paths (e.g., to reference original wowsim or a different fork):

1. **Find and Replace** in all spell files:
   - Find: `github.com/gefthi/bb-lock`
   - Replace with: `your-repo-path`

2. Or use the command line:
```bash
cd warlock/
sed -i 's|github.com/gefthi/bb-lock|your-repo-path|g' *.go
```

---

## Spell Metrics and Tracking

The spells automatically track:
- **Hits**: Successful spell casts that hit the target
- **Crits**: Critical strikes
- **Damage**: Total damage dealt

Access metrics via:
```go
spell.SpellMetrics[target.UnitIndex].Hits
spell.SpellMetrics[target.UnitIndex].Crits
spell.SpellMetrics[target.UnitIndex].TotalDamage
```

---

## Rotation Integration

### Simple APL (Action Priority List)

```go
func (warlock *Warlock) GetRotation() *core.APL {
    return core.NewAPL(
        // Priority 1: Chaos Bolt if available
        core.NewAPLAction(warlock.ChaosBolt, func(sim *core.Simulation) bool {
            return warlock.ChaosBolt.IsReady(sim)
        }),
        
        // Priority 2: Immolate if not active
        core.NewAPLAction(warlock.Immolate, func(sim *core.Simulation) bool {
            return !warlock.Immolate.Dot(sim.Target).IsActive()
        }),
        
        // Priority 3: Incinerate as filler
        core.NewAPLAction(warlock.Incinerate, func(sim *core.Simulation) bool {
            return true // Always available as filler
        }),
    )
}
```

---

## Debugging

### Enable Spell Logging
Add debug output if needed:

```go
ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
    baseDamage := sim.Roll(583, 676)
    
    // Debug output
    if sim.Debug {
        sim.Log("Incinerate: Base damage %0.0f", baseDamage)
    }
    
    // ... rest of spell logic
}
```

### Common Issues

1. **"time" package unused**
   - Comment out the import if not using time-based calculations
   - Spell durations are in milliseconds by default

2. **Nil pointer on Immolate check**
   - Ensure `warlock.Immolate` is registered before `warlock.Incinerate`
   - Add nil check in Incinerate if registration order is uncertain:
   ```go
   if warlock.Immolate != nil && warlock.Immolate.Dot(target).IsActive() {
       hasImmolate = true
   }
   ```

3. **Spell not casting**
   - Check mana costs
   - Verify GCD is respected
   - Ensure cooldowns are tracked properly

---

## Testing

### Unit Tests Template

```go
package warlock

import (
    "testing"
    "github.com/gefthi/bb-lock/sim/core"
)

func TestImmolate(t *testing.T) {
    warlock := NewWarlock()
    sim := core.NewSimulation()
    target := core.NewTarget()
    
    // Cast Immolate
    warlock.Immolate.Cast(sim, target)
    
    // Verify DoT is active
    if !warlock.Immolate.Dot(target).IsActive() {
        t.Error("Immolate DoT not applied")
    }
    
    // Verify tick count
    expectedTicks := 5
    if warlock.Immolate.Dot(target).NumberOfTicks != expectedTicks {
        t.Errorf("Expected %d ticks, got %d", 
            expectedTicks, 
            warlock.Immolate.Dot(target).NumberOfTicks)
    }
}

func TestIncinerateBonus(t *testing.T) {
    warlock := NewWarlock()
    sim := core.NewSimulation()
    target := core.NewTarget()
    
    // Cast without Immolate
    result1 := warlock.Incinerate.Cast(sim, target)
    damage1 := result1.Damage
    
    // Apply Immolate
    warlock.Immolate.Cast(sim, target)
    
    // Cast with Immolate
    result2 := warlock.Incinerate.Cast(sim, target)
    damage2 := result2.Damage
    
    // Verify bonus (approximately 25% more)
    expectedBonus := 1.25
    actualBonus := damage2 / damage1
    
    if actualBonus < expectedBonus * 0.9 || actualBonus > expectedBonus * 1.1 {
        t.Errorf("Expected ~%0.2fx damage, got %0.2fx", 
            expectedBonus, actualBonus)
    }
}
```

---

## Advanced Configuration

### Talent Modifiers

Add talent-based modifications:

```go
func (warlock *Warlock) applyTalents() {
    // Example: Emberstorm talent (5% more Fire damage)
    if warlock.Talents.Emberstorm > 0 {
        emberstormBonus := 1.0 + (0.01 * float64(warlock.Talents.Emberstorm))
        warlock.Immolate.DamageMultiplier *= emberstormBonus
        warlock.Incinerate.DamageMultiplier *= emberstormBonus
        warlock.ChaosBolt.DamageMultiplier *= emberstormBonus
    }
    
    // Example: Improved Immolate (increases crit chance)
    if warlock.Talents.ImprovedImmolate > 0 {
        bonusCrit := 0.05 * float64(warlock.Talents.ImprovedImmolate)
        warlock.Immolate.BonusCritRating += bonusCrit
    }
}
```

### Glyph Effects

```go
func (warlock *Warlock) applyGlyphs() {
    // Example: Glyph of Incinerate
    if warlock.Glyphs.Incinerate {
        // Increases Incinerate damage by 5%
        warlock.Incinerate.DamageMultiplier *= 1.05
    }
    
    // Example: Glyph of Immolate
    if warlock.Glyphs.Immolate {
        // Increases Immolate duration
        warlock.Immolate.Dot.NumberOfTicks += 2 // +6 seconds
    }
}
```

---

## Performance Optimization

### Spell Power Caching

If spell power doesn't change often:

```go
type Warlock struct {
    core.Character
    
    cachedSpellPower float64
    spellPowerDirty  bool
}

func (warlock *Warlock) GetSpellPower() float64 {
    if warlock.spellPowerDirty {
        warlock.cachedSpellPower = warlock.calculateSpellPower()
        warlock.spellPowerDirty = false
    }
    return warlock.cachedSpellPower
}
```

### Pre-calculated Constants

Store frequently used values:

```go
const (
    ImmolateSpellID     = 47811
    ImmolateBaseCost    = 0.17
    ImmolateCastTime    = 2000 // ms
    ImmolateDirectDmgMin = 505
    ImmolateDirectDmgMax = 621
    ImmolateDotDmg      = 785
    ImmolateDotTicks    = 5
    ImmolateDotInterval = 3000 // ms
)
```

---

## Checklist

Before integrating:
- [ ] Copy spell files to `sim/warlock/` directory
- [ ] Update Warlock struct with spell references
- [ ] Register spells in initialization
- [ ] Set appropriate BaseMana value
- [ ] Decide on time package usage (comment/uncomment)
- [ ] Test each spell individually
- [ ] Test spell interactions (Immolate + Incinerate)
- [ ] Integrate into rotation/APL system
- [ ] Add talent modifiers (if applicable)
- [ ] Add glyph effects (if applicable)
- [ ] Run full simulation tests

---

## Next Steps

1. **Implement additional spells**: Shadow Bolt, Curse of Agony, Corruption, etc.
2. **Add pet mechanics**: Imp, Felguard, etc.
3. **Implement talents**: Full Destruction/Affliction/Demonology trees
4. **Add glyphs**: Major and minor glyph effects
5. **Optimize rotation**: APL system for different scenarios
6. **Multi-target**: AoE spells and cleave mechanics

---

## Support

If you encounter issues:
1. Check the IMPLEMENTATION_NOTES.md for known issues
2. Verify import paths are correct
3. Ensure core simulation framework is properly set up
4. Review the WARLOCK_SPELLS.md for spell mechanics details

---

## Version
- **Document Version**: 1.0
- **Compatible with**: bb-lock initial implementation
- **WoW Version**: WotLK (3.3.5)
