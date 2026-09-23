package unstable

// floatRunVBMIHost and valid512Host are the host's body selectors, which the
// tests flip to run each body the CPU has.
var floatRunVBMIHost, valid512Host = useFloatRunVBMI, useValid512

// floatRunBodies lists the decimal-array kernel's bodies this host can run:
// the AVX2 one, which refines Eisel-Lemire's product where eiselLemire64 does,
// and, with AVX-512 VBMI, the VBMI one, which declines those numbers instead.
// Both take numbers of up to 19 digits; selecting one turns the kernel on.
func floatRunBodies() []kernelBody {
	if !floatRunHost {
		return nil
	}
	b := []kernelBody{{"avx2", true, func() {
		useFloatRun, useFloatRunLong, useFloatRunVBMI = true, true, false
	}}}
	if floatRunVBMIHost {
		b = append(b, kernelBody{"vbmi", false, func() {
			useFloatRun, useFloatRunLong, useFloatRunVBMI = true, true, true
		}})
	}
	return b
}

// validRunBodies lists the validation walks' bodies this host can run: AVX2,
// and with AVX-512BW the mask-register one. Selecting one turns both walks,
// flat and ring, on.
func validRunBodies() []kernelBody {
	if !validRunHost {
		return nil
	}
	b := []kernelBody{{"avx2", false, func() {
		useValidRun, useValidPoints, useValid512 = true, true, false
	}}}
	if valid512Host {
		b = append(b, kernelBody{"avx512", false, func() {
			useValidRun, useValidPoints, useValid512 = true, true, true
		}})
	}
	return b
}

// restoreKernels puts every number-kernel flag back to the host's.
func restoreKernels() {
	useFloatRun, useFloatRunLong, useFloatRunVBMI = floatRunHost, floatRunLongHost, floatRunVBMIHost
	useValidRun, useValidPoints, useValid512 = validRunHost, validPointsHost, valid512Host
}
