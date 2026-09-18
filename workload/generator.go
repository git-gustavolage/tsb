package workload

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type GeneratorConfig struct {
	WorkloadID uint64
	Name       string

	Tasks int

	MinWorkUnits uint64
	MaxWorkUnits uint64

	DependencyProbability float64
	MaxDependencies       int

	Seed int64
}

type Generator struct {
	config GeneratorConfig
	rng    *rand.Rand
}

func NewGenerator(config GeneratorConfig) (*Generator, error) {
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return &Generator{
		config: config,
		rng:    rand.New(rand.NewSource(config.Seed)),
	}, nil
}

func (g *Generator) Generate() *Workload {
	workload := &Workload{
		ID:    g.config.WorkloadID,
		Name:  g.config.Name,
		Tasks: make([]Task, 0, g.config.Tasks),
	}

	for i := 0; i < g.config.Tasks; i++ {
		task := Task{
			ID:        uint64(i + 1),
			WorkUnits: g.randomUint64(g.config.MinWorkUnits, g.config.MaxWorkUnits),
		}

		if i > 0 && g.shouldHaveDependencies() {
			task.DependsOn = g.generateDependencies(i, g.config.MaxDependencies)
		}

		workload.Tasks = append(workload.Tasks, task)
	}

	return workload
}

func (g *Generator) GenerateFile(path string) error {
	workload := g.Generate()

	data, err := json.MarshalIndent(workload, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal workload: %w", err)
	}

	data = append(data, '\n')

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write workload file: %w", err)
	}

	return nil
}

func validateConfig(config GeneratorConfig) error {
	if config.Tasks <= 0 {
		return fmt.Errorf("tasks must be greater than zero")
	}

	if config.MinWorkUnits == 0 {
		return fmt.Errorf("min work units must be greater than zero")
	}

	if config.MaxWorkUnits < config.MinWorkUnits {
		return fmt.Errorf("max work units must be greater than or equal to min work units")
	}

	if config.DependencyProbability < 0 || config.DependencyProbability > 1 {
		return fmt.Errorf("dependency probability must be between 0 and 1")
	}

	if config.MaxDependencies < 0 {
		return fmt.Errorf("max dependencies must be greater than or equal to zero")
	}

	return nil
}

func (g *Generator) shouldHaveDependencies() bool {
	return g.rng.Float64() < g.config.DependencyProbability
}

func (g *Generator) generateDependencies(taskIndex int, maxDependencies int) []uint64 {
	available := taskIndex

	if maxDependencies > available {
		maxDependencies = available
	}

	if maxDependencies == 0 {
		return nil
	}

	dependencyCount := g.randomInt(1, maxDependencies)

	ids := make([]uint64, available)

	for i := range available {
		ids[i] = uint64(i + 1)
	}

	g.rng.Shuffle(len(ids), func(i, j int) {
		ids[i], ids[j] = ids[j], ids[i]
	})

	dependencies := ids[:dependencyCount]

	return dependencies
}

func (g *Generator) randomInt(min, max int) int {
	if min == max {
		return min
	}

	return min + g.rng.Intn(max-min+1)
}

func (g *Generator) randomUint64(min, max uint64) uint64 {
	if min == max {
		return min
	}

	return min + uint64(g.rng.Int63n(int64(max-min+1)))
}
