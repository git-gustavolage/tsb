package main

import (
	"flag"
	"fmt"
	"log"

	"tsb/benchmark"
)

func main() {
	var (
		output      = flag.String("output", "benchmark.json", "arquivo de saída")
		benchmarkID = flag.Uint64("id", 1, "identificador do benchmark")
		chunks      = flag.Int("chunks", 10, "quantidade de chunks")
		minTasks    = flag.Int("min-tasks", 10, "quantidade mínima de tasks por chunk")
		maxTasks    = flag.Int("max-tasks", 10, "quantidade máxima de tasks por chunk")
		minWork     = flag.Uint64("min-work", 100, "quantidade mínima de work units")
		maxWork     = flag.Uint64("max-work", 1000, "quantidade máxima de work units")
		minDelay    = flag.Uint64("min-delay", 1000, "delay mínimo entre chunks em ms")
		maxDelay    = flag.Uint64("max-delay", 1000, "delay máximo entre chunks em ms")
		seed        = flag.Int64("seed", 1, "seed utilizado para geração determinística")
	)

	flag.Parse()

	config := benchmark.GeneratorConfig{
		BenchmarkID:      *benchmarkID,
		Chunks:           *chunks,
		MinTasksPerChunk: *minTasks,
		MaxTasksPerChunk: *maxTasks,
		MinWorkUnits:     *minWork,
		MaxWorkUnits:     *maxWork,
		MinDelayMS:       *minDelay,
		MaxDelayMS:       *maxDelay,
		Seed:             *seed,
	}

	generator, err := benchmark.NewGenerator(config)
	if err != nil {
		log.Fatalf("erro ao criar generator: %v", err)
	}

	if err := generator.GenerateFile(*output); err != nil {
		log.Fatalf("erro ao gerar benchmark: %v", err)
	}

	fmt.Printf("benchmark generated: %s\n", *output)
}
