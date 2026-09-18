package workload

type Workload struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID        uint64   `json:"id"`
	WorkUnits uint64   `json:"work_units"`
	DependsOn []uint64 `json:"depends_on"`
}
