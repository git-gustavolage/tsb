package commands

import (
	"flag"
	"fmt"
	"log"

	"tsb/workload"
)

func runWorkload(args []string) {
	if len(args) == 0 {
		printWorkloadUsage()
		return
	}

	switch args[0] {
	case "generate":
		runWorkloadGenerate(args[1:])

	case "help", "--help", "-h":
		printWorkloadUsage()

	default:
		fmt.Printf("unknown workload command: %s\n", args[0])
		printWorkloadUsage()
	}
}

func runWorkloadGenerate(args []string) {
	flags := flag.NewFlagSet("workload generate", flag.ExitOnError)

	output := flags.String("output", "workload.json", "arquivo de saída")
	workloadID := flags.Uint64("id", 1, "identificador do workload")
	name := flags.String("name", "generated-workload", "nome do workload")
	tasks := flags.Int("tasks", 100, "quantidade de tasks")
	minWork := flags.Uint64("min-work", 100, "quantidade mínima de work units")
	maxWork := flags.Uint64("max-work", 1000, "quantidade máxima de work units")
	dependencyProbability := flags.Float64("dependency-probability", 0.5, "probabilidade de uma task possuir dependências")
	maxDependencies := flags.Int("max-dependencies", 3, "quantidade máxima de dependências")
	seed := flags.Int64("seed", 1, "seed utilizada para geração determinística")

	flags.Parse(args)

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

func printWorkloadUsage() {
	fmt.Println(`Usage: tsb workload <command> [options]

Commands:
  generate    Generate a workload file
  help        Show this help message

Examples:
  tsb workload generate
  tsb workload generate --tasks 1000
  tsb workload generate --output workload.json --seed 42

Use:
  tsb workload generate --help`)
}
