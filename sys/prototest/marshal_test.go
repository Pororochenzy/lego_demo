package prototest

import (
	"testing"

	jsoniter "github.com/json-iterator/go"

	"google.golang.org/protobuf/proto"
)

func BenchmarkMarsh(b *testing.B) {
	var (
		// ncpu = runtime.NumCPU()
		parallelisms = []int{4, 16, 96}
		dataSizes    = []int{10}
		numKeys      = 1024
		builders     = []TargetBuilder{
			{
				Name: "json",
				Make: func(bench Benchmark) (Target, error) {

					return Target{
						DoMarshal: func(data *JSONData) ([]byte, error) {
							b, err := jsoniter.Marshal(data)
							return b, err
						},
						DoUnMarshal: func(data []byte) error {
							err := jsoniter.Unmarshal(data, &JSONData{})
							return err
						},
					}, nil
				},
			},
			{
				Name: "proto",
				Make: func(bench Benchmark) (Target, error) {
					return Target{
						DoMarshal: func(data *JSONData) ([]byte, error) {
							b, err := proto.Marshal(data)
							return b, err
						},
						DoUnMarshal: func(data []byte) error {
							err := proto.Unmarshal(data, &JSONData{})
							return err
						},
					}, nil
				},
			},
		}
	)

	RunBenchmark(b, compose(parallelisms, dataSizes, numKeys, builders))
}
