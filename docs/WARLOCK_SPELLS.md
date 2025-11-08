# Warlock Spell Implementations

## Overview
This document describes the three basic Warlock spells implemented for the BB-Lock simulator, based on WotLK (Wrath of the Lich King) mechanics.

## Implemented Spells

### 1. Immolate
**File**: `warlock/immolate.go`

**Spell Details**:
- **Spell ID**: 47811 (Rank 11)
- **School**: Fire
- **Type**: DoT (Damage over Time) with initial direct damage
- **Mana Cost**: 17% of base mana
- **Cast Time**: 2.0 seconds
- **GCD**: Standard (1.5 seconds)

**Damage Components**:
- **Direct Damage**: 505-621 Fire damage on cast
- **DoT Damage**: 785 total Fire damage over 15 seconds
  - 5 ticks
  - 3 seconds per tick
  - ~157 damage per tick

**Mechanics**:
- Applies a DoT debuff to the target
- DoT snapshots spell power and other modifiers on application
- Can be refreshed/reapplied
- Used as a prerequisite for Incinerate bonus damage

**Implementation Notes**:
- Uses snapshot mechanics for DoT ticks
- Properly handles spell hit and crit for initial damage
- DoT ticks use snapshotted crit chance

---

### 2. Incinerate
**File**: `warlock/incinerate.go`

**Spell Details**:
- **Spell ID**: 47838 (Rank 4)
- **School**: Fire
- **Type**: Direct damage nuke
- **Mana Cost**: 14% of base mana
- **Cast Time**: 2.5 seconds
- **GCD**: Standard (1.5 seconds)

**Damage Components**:
- **Base Damage**: 583-676 Fire damage
- **Conditional Bonus**: 25% increased damage if target has Immolate DoT active

**Mechanics**:
- Checks if target has active Immolate DoT
- Applies 1.25x damage multiplier if Immolate is present
- Standard spell hit and crit mechanics
- Primary single-target nuke for Destruction spec

**Implementation Notes**:
- Queries the Immolate DoT status on target
- Dynamically adjusts damage based on Immolate presence
- Efficient for spell power scaling

---

### 3. Chaos Bolt
**File**: `warlock/chaos_bolt.go`

**Spell Details**:
- **Spell ID**: 59172 (Rank 4)
- **School**: Fire + Chaos
- **Type**: Direct damage nuke (always crits)
- **Mana Cost**: 7% of base mana
- **Cast Time**: 2.5 seconds
- **Cooldown**: 12 seconds
- **GCD**: Standard (1.5 seconds)

**Damage Components**:
- **Base Damage**: 1429-1813 Fire damage
- **Special**: Always critically strikes

**Mechanics**:
- **Cannot be resisted**: Bypasses partial resistance mechanics
- **Always crits**: Forces critical strike on hit
- Uses both Fire and Chaos spell schools
- High mana efficiency for burst damage
- Cooldown-gated to prevent spam

**Implementation Notes**:
- Uses `OutcomeMagicCrit` to force critical strikes
- Spell school flags prevent resistance
- Tracks both hits and crits in metrics
- Cooldown timer prevents casting more than once per 12 seconds

---

## Spell Synergies

### Immolate + Incinerate
- **Immolate** should be maintained on the target
- **Incinerate** deals 25% more damage while Immolate is active
- Optimal rotation: Apply Immolate, then spam Incinerate

### Chaos Bolt Usage
- Use on cooldown for maximum DPS
- High burst damage with guaranteed crit
- Low mana cost makes it very efficient
- 12-second cooldown requires timing in rotation

---

## Rotation Example

**Basic Single-Target Rotation**:
1. Cast **Immolate** on target (if not already active)
2. Cast **Chaos Bolt** (if off cooldown)
3. Cast **Incinerate** as filler
4. Refresh **Immolate** before it expires
5. Repeat

**Priority System**:
1. Chaos Bolt (if available)
2. Immolate (if not active or about to expire)
3. Incinerate (filler)

---

## Implementation Status

### Completed
- ✅ Basic spell structure for all three spells
- ✅ Mana costs
- ✅ Cast times
- ✅ Damage calculations
- ✅ Immolate DoT mechanics
- ✅ Incinerate conditional damage bonus
- ✅ Chaos Bolt guaranteed crit mechanic
- ✅ Chaos Bolt cooldown

### Pending
- ⏳ Talent modifications (e.g., reduced cast times, increased damage)
- ⏳ Glyph effects
- ⏳ Set bonuses
- ⏳ Integration with full Warlock rotation system
- ⏳ Spell power coefficient scaling (may need adjustment)
- ⏳ Integration tests

### Future Enhancements
- Backdraft (Conflagrate talent) interaction
- Molten Core proc (Fire and Brimstone talent)
- Improved Immolate talent
- Empowered Imp talent (Incinerate bonus)
- Spell power scaling adjustments
- Multiple target scenarios

---

## Technical Notes

### Dependencies
- `github.com/gefthi/bb-lock/sim/core` - Core simulation framework
- Warlock base class structure (assumed to exist)

### Required Warlock Fields
The spell implementations assume the Warlock struct has:
```go
type Warlock struct {
    core.Character
    BaseMana float64
    Immolate *core.Spell
    Incinerate *core.Spell
    ChaosBolt *core.Spell
    // ... other fields
}
```

### Registration
These spells need to be registered in the Warlock initialization:
```go
func (warlock *Warlock) Initialize() {
    warlock.registerImmolateSpell()
    warlock.registerIncinerateSpell()
    warlock.registerChaosBoltSpell()
    // ... other spells
}
```

---

## Testing Checklist

- [ ] Verify spell IDs match WotLK values
- [ ] Test Immolate initial damage
- [ ] Test Immolate DoT ticks (count and damage)
- [ ] Test Incinerate base damage
- [ ] Test Incinerate bonus damage with Immolate active
- [ ] Test Incinerate without Immolate (no bonus)
- [ ] Test Chaos Bolt guaranteed crit
- [ ] Test Chaos Bolt cooldown enforcement
- [ ] Verify mana costs are correct
- [ ] Test spell power scaling
- [ ] Test spell hit/miss mechanics
- [ ] Test spell crit mechanics
- [ ] Integration test: Full rotation DPS

---

## References
- WoWHead: Immolate (Rank 11) - https://www.wowhead.com/wotlk/spell=47811
- WoWHead: Incinerate (Rank 4) - https://www.wowhead.com/wotlk/spell=47838
- WoWHead: Chaos Bolt (Rank 4) - https://www.wowhead.com/wotlk/spell=59172
- WotLK Warlock Mechanics: Various community resources

---

## Version History
- **v0.1** - Initial implementation of three core spells
  - Basic damage calculations
  - Core mechanics (DoT, conditional bonus, guaranteed crit)
  - Mana costs and cast times
