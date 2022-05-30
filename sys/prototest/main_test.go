package prototest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/valyala/fastrand"
)

type Benchmark struct {
	Parallelism   int
	Data          *JSONData
	TargetBuilder TargetBuilder
}

type Target struct {
	DoMarshal   func(data *JSONData) ([]byte, error)
	DoUnMarshal func(data []byte) error
}

type TargetBuilder struct {
	Name string
	Make func(bench Benchmark) (Target, error)
}

func gen(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(byte(fastrand.Uint32n(26) + 'a'))
	}
	return sb.String()
}

func genData(n int) *JSONData {
	childrens := []*Children{}
	for i := 0; i < n; i++ {
		children := &Children{
			ID:       int32(fastrand.Uint32()),
			NodeName: "fastrand.Uint32() + 'c'",
		}
		childrens = append(childrens, children)
	}

	root := &Root{
		ID:       int32(fastrand.Uint32()),
		NodeName: "fastrand.Uint32() + 'n'",
		Children: childrens,
	}
	return &JSONData{
		ID:   int32(fastrand.Uint32()),
		Name: "fastrand.Uint32() + 'a'",
		Root: root,
	}
}

func gens(s, n int) (gs []string) {
	gs = make([]string, n)
	for i := range gs {
		gs[i] = gen(s)
	}
	return gs
}

func compose(parallelisms, dataSize []int, numKeys int, builders []TargetBuilder) []Benchmark {
	benchmarks := make([]Benchmark, 0, len(parallelisms)*len(dataSize)*len(builders))

	for _, p := range parallelisms {
		for _, k := range dataSize {
			d := genData(k)
			for _, builder := range builders {
				benchmarks = append(benchmarks, Benchmark{
					Parallelism:   p,
					Data:          d,
					TargetBuilder: builder,
				})
			}

		}
	}

	return benchmarks
}

func RunBenchmark(b *testing.B, benchmarks []Benchmark) {
	for _, bench := range benchmarks {
		bench := bench
		b.Run(fmt.Sprintf("%s-parallelism(%d)- ", bench.TargetBuilder.Name, bench.Parallelism), func(b *testing.B) {
			target, err := bench.TargetBuilder.Make(bench)
			if err != nil {
				b.Fatalf("%s setup fail: %v", bench.TargetBuilder.Name, err)
			}
			if bench.Parallelism == 0 {
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					b1, err := target.DoMarshal(bench.Data)
					if err != nil {
						b.Errorf("%s error during benchmark: %v", bench.TargetBuilder.Name, err)
					}

					if err := target.DoUnMarshal(b1); err != nil {
						b.Errorf("%s error during benchmark: %v", bench.TargetBuilder.Name, err)
					}
				}

			} else {
				b.SetParallelism(bench.Parallelism)
				b.ResetTimer()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						b2, err := target.DoMarshal(bench.Data)
						if err != nil {
							b.Errorf("%s error during benchmark: %v", bench.TargetBuilder.Name, err)
						}

						if err := target.DoUnMarshal(b2); err != nil {
							b.Errorf("%s error during benchmark: %v", bench.TargetBuilder.Name, err)
						}
					}
				})
			}
			b.StopTimer()
		})
	}
}
