package benchmark

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type GeneratorConfig struct {
	BenchmarkID uint64

	Chunks int

	MinTasksPerChunk int
	MaxTasksPerChunk int

	MinWorkUnits uint64
	MaxWorkUnits uint64

	MinDelayMS uint64
	MaxDelayMS uint64

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

func (g *Generator) Generate() *Benchmark {
	benchmark := &Benchmark{
		ID:     g.config.BenchmarkID,
		Chunks: make([]Chunk, 0, g.config.Chunks),
	}

	var taskID uint64 = 1

	for chunkID := 0; chunkID < g.config.Chunks; chunkID++ {
		taskCount := g.randomInt(
			g.config.MinTasksPerChunk,
			g.config.MaxTasksPerChunk,
		)

		chunk := Chunk{
			ID: uint64(chunkID),
			DelayMS: g.randomUint64(
				g.config.MinDelayMS,
				g.config.MaxDelayMS,
			),
			Tasks: make([]Task, 0, taskCount),
		}

		for range taskCount {
			task := Task{
				ID: taskID,
				WorkUnits: g.randomUint64(
					g.config.MinWorkUnits,
					g.config.MaxWorkUnits,
				),
			}

			chunk.Tasks = append(chunk.Tasks, task)

			taskID++
		}

		benchmark.Chunks = append(benchmark.Chunks, chunk)
	}

	return benchmark
}

func (g *Generator) GenerateFile(path string) error {
	benchmark := g.Generate()

	data, err := json.MarshalIndent(benchmark, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal benchmark: %w", err)
	}

	data = append(data, '\n')

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write benchmark file: %w", err)
	}

	return nil
}

func validateConfig(config GeneratorConfig) error {
	if config.Chunks <= 0 {
		return fmt.Errorf("chunks must be greater than zero")
	}

	if config.MinTasksPerChunk <= 0 {
		return fmt.Errorf("min tasks per chunk must be greater than zero")
	}

	if config.MaxTasksPerChunk < config.MinTasksPerChunk {
		return fmt.Errorf(
			"max tasks per chunk must be greater than or equal to min tasks per chunk",
		)
	}

	if config.MinWorkUnits == 0 {
		return fmt.Errorf("min work units must be greater than zero")
	}

	if config.MaxWorkUnits < config.MinWorkUnits {
		return fmt.Errorf(
			"max work units must be greater than or equal to min work units",
		)
	}

	if config.MaxDelayMS < config.MinDelayMS {
		return fmt.Errorf(
			"max delay must be greater than or equal to min delay",
		)
	}

	return nil
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
