package statusz

import "fmt"

type Var interface {
	MarshalMetrics() ([]*Metric, error)
}

type VarMetric interface {
	MarshalMetric() (*Metric, error)
}

var _ Var = (VarGroup)(nil)

type VarGroup []interface{}

func (vg VarGroup) MarshalMetrics() ([]*Metric, error) {
	metrics := make([]*Metric, 0, len(vg))
	for _, v := range vg {
		switch t := v.(type) {
		case VarMetric:
			metric, err := t.MarshalMetric()
			if err != nil {
				return nil, err
			}
			metrics = append(metrics, metric)
		case Var:
			metric, err := t.MarshalMetrics()
			if err != nil {
				return nil, err
			}
			metrics = append(metrics, metric...)
		default:
			panic(fmt.Errorf("unknown var type %T", v))
		}

	}
	return metrics, nil
}
