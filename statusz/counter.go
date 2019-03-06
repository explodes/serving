package statusz

var _ VarMetric = (*Counter)(nil)

type Counter struct {
	name  string
	count uint64
}

func NewCounter(name string) *Counter {
	return &Counter{
		name:  name,
		count: 0,
	}
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) MarshalMetric() (*Metric, error) {
	return &Metric{Name: c.name, Value: &Metric_U64{U64: c.count}}, nil
}
