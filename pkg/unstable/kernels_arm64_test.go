package unstable

// floatRunBodies lists the decimal-array kernel's bodies this host can run:
// the one NEON body, which refines Eisel-Lemire's product where eiselLemire64
// does — on a core with the dot-product instruction its fold uses, and none
// on one without (a Cortex-A72: Graviton1, a Raspberry Pi 4), where selecting
// it would be an illegal instruction.
func floatRunBodies() []kernelBody {
	if !floatRunHost {
		return nil
	}
	return []kernelBody{{"neon", true, func() { useFloatRun, useFloatRunLong = true, true }}}
}

// validRunBodies lists the validation walks' bodies this host can run: the
// NEON walks, which share the conversion kernel's gate.
func validRunBodies() []kernelBody {
	if !validRunHost {
		return nil
	}
	return []kernelBody{{"neon", false, func() { useValidRun, useValidPoints = true, true }}}
}

// restoreKernels puts every number-kernel flag back to the host's.
func restoreKernels() {
	useFloatRun, useFloatRunLong = floatRunHost, floatRunLongHost
	useValidRun, useValidPoints = validRunHost, validPointsHost
}
