package register

import (
	"testing"
)

// test errors returned by both parseMap and parseConditional for invalid register names
func Test_invalid(t *testing.T) {
	for _, regName := range []string{"0", "01", "$", "$a", "$v2", "$a4", "$tf", "$s8", "$k2", "$r1"} {
		t.Run(regName, func(t *testing.T) {
			if _, err := parseMap(regName); err == nil {
				t.Errorf("parseMap() = %v, expected error", regName)
			}
			if _, err := parseConditional(regName); err == nil {
				t.Errorf("parseConditional() = %v, expected error", regName)
			}
		})
	}
}

// itterate through all registers in the register map, and verify that the map and conditional techniques, as well as the public Parse function, return the same results
func Test_compareMapVsConditional(t *testing.T) {
	for k := range registers {
		regName := "$" + k
		t.Run(regName, func(t *testing.T) {
			pParse, errParse := Parse(regName)
			pCond, errCond := parseConditional(regName)
			pMap, errMap := parseMap(regName)

			if pCond != pMap {
				t.Errorf("parseMap() = %v, parseConditional() = %v", pMap, pCond)
				return
			}
			if pParse != pMap {
				t.Errorf("Parse() = %v, paresMap() = %v", pParse, pMap)
				return
			}
			if errParse != nil {
				t.Errorf("parseConditional() error = %v", errCond)
			}
			if errCond != nil {
				t.Errorf("parseConditional() error = %v", errCond)
			}
			if errMap != nil {
				t.Errorf("parseMap() error = %v", errMap)
			}
		})
	}
}

func Benchmark_parseMap(b *testing.B) {
	// make array of registers
	numRegs := len(registers)
	allNames := make([]string, numRegs)
	i := 0
	for k := range registers {
		allNames[i] = k
		i++
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		parseMap(allNames[n%numRegs])
	}
}

func Benchmark_parseConditional(b *testing.B) {
	// make array of registers
	numRegs := len(registers)
	allNames := make([]string, numRegs)
	i := 0
	for k := range registers {
		allNames[i] = k
		i++
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		parseConditional(allNames[n%numRegs])
	}
}
