package main

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

/*
Benchmark vs profiling in general:

Benchmark(basic analysis) - Used to check the overall time and memory taken for set of code or function.

Profiling(Detailed anlysis) - Can check the different profiling like cpu/mem/concurrency in line by line and more detailed technical runtime environment parameters

*/

//*testing.B - benchmark - Should follow the same function name and file name format for this benchmark.

// here below three different method/logic functions for string concatenation to compare the performance by benchmark.

//go command for running all benchmark functions
//go test -bench . -benchmem

//running all the benchmark functions for 10 times for average more reliable results, and stores result in txt file.
//go test -bench . -benchmem -count 10 > 10_runs_bench.txt

/*                                         Average time exe for each loop
BenchmarkStringBuilderConcatenation-8             485919              2472 ns/op            3320 B/op          9 allocs/op
BenchmarkStringConcatenation-8                      6592            162327 ns/op          530277 B/op        999 allocs/op
BenchmarkFmtSprintfConcatenation-8                  4260            284130 ns/op          546557 B/op       1998 allocs/op

Above values column explanation:
485919  - No of times outer benchmark loop executed - for i := 0; i < b.N; i++ {}
2472 ns/op - Average Time taken for each benchmark loop execution
3320 B/op, 9 allocs/op - Memory byte allocation for each benchmark loop execution

EX:BenchmarkFmtSprintfConcatenation-8  //this "8" denotes number of system threads used to run this benchmark
*/

func BenchmarkStringBuilderConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for j := 0; j < 1000; j++ {
			sb.WriteString("h")
		}
		_ = sb.String()
	}
}

func BenchmarkStringConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < 1000; j++ {
			s += "h"
		}
	}
}

func BenchmarkFmtSprintfConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < 1000; j++ {
			s = fmt.Sprintf("%s%s", s, "h")
		}
	}
}

func BenchmarkGoRoutineSample(b *testing.B) {

	for i := 0; i < b.N; i++ {
		wg := &sync.WaitGroup{}
		for i := 0; i < 10000; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 10000000; i++ {
					_ = 10
				}
				wg.Done()
			}()
		}

		wg.Wait()

	}

}

/*
Profiling with go test benchmark:

go test -bench='.' -cpuprofile='cpu.prof' -memprofile='mem.prof'  //This command will create the ".prof" files for cpu profiling and memory profiling.

go tool pprof cpu.prof //then using this command we can read the .prof file and run commands into this to analyse this


Alternate more popular ways of profiling packages:

runtime/pprof
net/pprof

*/
