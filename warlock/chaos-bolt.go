package warlock

import (
	// "time" // Uncomment if needed for duration calculations
	"github.com/gefthi/bb-lock/core"
)

// Chaos Bolt is a powerful direct damage Fire spell that cannot be resisted and always crits
func (warlock *Warlock) registerChaosBoltSpell() {
	baseCost := 0.07 * warlock.BaseMana

	warlock.ChaosBolt = warlock.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 59172}, // Rank 4 (WotLK)
		SpellSchool: core.SpellSchoolFire | core.SpellSchoolChaos,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{
			BaseCost: baseCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: 2500, // 2.5 second cast time (in milliseconds)
			},
			CD: core.Cooldown{
				Timer:    warlock.NewTimer(),
				Duration: 12000, // 12 second cooldown (in milliseconds)
			},
		},

		DamageMultiplier: 1,
		CritMultiplier:   warlock.DefaultSpellCritMultiplier(),
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			// Base damage range for Chaos Bolt
			baseDamage := sim.Roll(1429, 1813)

			// Chaos Bolt always crits (special mechanic)
			// We'll force a crit by using OutcomeMagicCrit
			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMagicCrit)

			// Chaos Bolt cannot be resisted/partially resisted (handled by spell school flags)
			if result.Landed() {
				spell.SpellMetrics[target.UnitIndex].Hits++
				spell.SpellMetrics[target.UnitIndex].Crits++
			}

			spell.DealDamage(sim, result)
		},
	})
}
