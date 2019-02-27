package statusz

import "runtime"

func init() {
	registerSystemVar()
}

type systemVar struct{}

func (systemVar) Marshal() ([]*Metric, error) {
	ms := &runtime.MemStats{}
	runtime.ReadMemStats(ms)
	metrics := []*Metric{
		{Name: "allocated_mem", Value: &Metric_U64{U64: ms.Alloc}},
		{Name: "total_allocated_mem", Value: &Metric_U64{U64: ms.TotalAlloc}},
		{Name: "virtual_mem", Value: &Metric_U64{U64: ms.Sys}},
	}
	return metrics, nil
}

func registerSystemVar() {
	Register("system", systemVar{})
}
