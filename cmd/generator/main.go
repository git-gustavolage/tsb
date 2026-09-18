package main

import (
	"flag"
	"fmt"
	"log"

	"tsb/workload"
)

func main() {
	var (
		output                = flag.String("output", "workload.json", "arquivo de saída")
		workloadID            = flag.Uint64("id", 1, "identificador do workload")
		name                  = flag.String("name", "generated-workload", "nome do workload")
		tasks                 = flag.Int("tasks", 100, "quantidade de tasks")
		minWork               = flag.Uint64("min-work", 100, "quantidade mínima de work units por task")
		maxWork               = flag.Uint64("max-work", 1000, "quantidade máxima de work units por task")
		dependencyProbability = flag.Float64("dependency-probability", 0.5, "probabilidade de uma task possuir dependências (0.0 a 1.0)")
		maxDependencies       = flag.Int("max-dependencies", 3, "quantidade máxima de dependências por task")
		seed                  = flag.Int64("seed", 1, "seed utilizada para geração determinística")
	)

	flag.Parse()

	config := workload.GeneratorConfig{
		WorkloadID:            *workloadID,
		Name:                  *name,
		Tasks:                 *tasks,
		MinWorkUnits:          *minWork,
		MaxWorkUnits:          *maxWork,
		DependencyProbability: *dependencyProbability,
		MaxDependencies:       *maxDependencies,
		Seed:                  *seed,
	}

	generator, err := workload.NewGenerator(config)
	if err != nil {
		log.Fatalf("erro ao criar generator: %v", err)
	}

	if err := generator.GenerateFile(*output); err != nil {
		log.Fatalf("erro ao gerar workload: %v", err)
	}

	fmt.Printf("workload generated: %s\n", *output)
}
