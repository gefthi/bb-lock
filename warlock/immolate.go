package warlock

import (
	// "time" // Uncomment if needed for duration calculations
	"github.com/gefthi/bb-lock/core"
)

// Immolate is a Fire damage over time spell that also deals initial direct damage
func (warlock *Warlock) registerImmolateSpell() {
	baseCost := 0.17 * warlock.BaseMana

	warlock.Immolate = warlock.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 47811}, // Rank 11 (WotLK)
		SpellSchool: core.SpellSchoolFire,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{
			BaseCost: baseCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: 2000, // 2 second cast time (in milliseconds)
			},
		},

		DamageMultiplier: 1,
		CritMultiplier:   warlock.DefaultSpellCritMultiplier(),
		ThreatMultiplier: 1,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: "Immolate",
			},
			NumberOfTicks: 5,
			TickLength:    3000, // 3 seconds per tick (15 second duration total)
			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				// Calculate periodic damage
				baseDamage := 785.0 / 5.0 // Total DoT damage divided by ticks
				dot.SnapshotBaseDamage = baseDamage
				dot.SnapshotAttackerMultiplier = dot.Spell.AttackerDamageMultiplier(dot.Spell.Unit.AttackTables[target.UnitIndex])
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeSnapshotCrit)
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			// Direct damage component
			baseDamage := sim.Roll(505, 621) // Direct damage range
			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)

			if result.Landed() {
				spell.SpellMetrics[target.UnitIndex].Hits++
				spell.DealDamage(sim, result)

				// Apply the DoT
				spell.Dot(target).Apply(sim)
			}
		},
	})
}
