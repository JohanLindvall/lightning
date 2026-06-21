package unstable

import "testing"

func TestDaysFromCivilCached(t *testing.T) {
	monthLen := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for y := 1850; y <= 2400; y++ {
		leap := y%4 == 0 && (y%100 != 0 || y%400 == 0)
		for m := 1; m <= 12; m++ {
			dmax := monthLen[m-1]
			if m == 2 && leap {
				dmax = 29
			}
			for d := 1; d <= dmax; d++ {
				if got, want := daysFromCivilCached(y, m, d), daysFromCivil(y, m, d); got != want {
					t.Fatalf("%04d-%02d-%02d: cached=%d, algorithm=%d", y, m, d, got, want)
				}
			}
		}
	}
}
