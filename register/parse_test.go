package register

import (
	"testing"
)

func Test_Parse_compareMapVsConditional(t *testing.T) {
	for k := range registers {
		regName := "$" + k
		t.Run(regName, func(t *testing.T) {
			pCond, errCond := parseConditional(regName)
			pMap, errMap := parseMap(regName)

			if pCond != pMap {
				t.Errorf("parseMap() = %v, parseConditional() = %v", pCond, pMap)
				return
			}
			if errCond != errMap {
				t.Errorf("parseMap() = %v, parseConditional() = %v", errCond, errMap)
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
