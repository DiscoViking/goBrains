/*
 * Booster structures.
 *
 * This covers the velocity output behaviours that result in a change in location for a creature.
 */

package creature

// Fixed values.
const (
	// Booster types.
	BoosterLinear  = 1
	BoosterAngular = 2

	// Velocity scaling.  Linear is in pixels, but rotation is in radians - so scale the latter down.
	LinPerAng = 10

	// Maximum velocity.
	MaxLinVel = 1.0
	MaxAngVel = MaxLinVel / LinPerAng
)

// Outputs are chargeable.  This means they accept accept charge from nodes in the brain.
func (b *booster) Charge(strength float64) {
	b.charge += strength
}

// Outputs are workers.  This means that the brain will trigger them during processing to perform their actions.
func (b *booster) Work() {
	if b.charge < 0 {
		b.charge = 0
	}

	if b.btype == BoosterLinear {
		b.host.movement.move += b.charge * 0.2 * b.scale
	} else if b.btype == BoosterAngular {
		b.host.movement.rotate += b.charge * 0.2 / LinPerAng * b.scale
	}

	// Cap movement speeds
	if b.host.movement.move > MaxLinVel {
		b.host.movement.move = MaxLinVel
	} else if b.host.movement.move < -MaxLinVel {
		b.host.movement.move = -MaxLinVel
	}
	if b.host.movement.rotate > MaxAngVel {
		b.host.movement.rotate = MaxAngVel
	} else if b.host.movement.rotate < -MaxAngVel {
		b.host.movement.rotate = -MaxAngVel
	}

	// Reset charge now it has been used.
	b.charge = 0
}

// Initialize a new generic booster object.
func (host *Creature) newGenBooster(btype int, scale float64) *booster {

	newBoost := booster{
		outputStruct: outputStruct{
			putStruct: putStruct{host: host},
			charge:    0,
		},
		btype: btype,
		scale: scale,
	}
	b := &newBoost

	// Link the output to the hosts' brain.
	host.brain.AddOutput(b)

	return b
}

// Add a standard set of boosters to a host; one angular and one linear.
func (host *Creature) AddBoosters() (*booster, *booster, *booster, *booster) {
	l1 := host.newGenBooster(BoosterLinear, 1.0)
	l2 := host.newGenBooster(BoosterLinear, -1.0)
	a1 := host.newGenBooster(BoosterAngular, 1.0)
	a2 := host.newGenBooster(BoosterAngular, -1.0)
	return l1, l2, a1, a2
}
