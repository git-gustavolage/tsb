package benchmark

type Benchmark struct {
	ID     uint64  `json:"id"`
	Chunks []Chunk `json:"chunks"`
}

type Chunk struct {
	ID      uint64 `json:"id"`
	DelayMS uint64 `json:"delay_ms"`
	Tasks   []Task `json:"tasks"`
}

type Task struct {
	ID        uint64 `json:"id"`
	WorkUnits uint64 `json:"work_units"`
}
