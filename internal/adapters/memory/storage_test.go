package memory

import (
	"testing"

	"github.com/VoevodinAnton/metrics/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestStorage_UpdateGauge(t *testing.T) {
	type fields struct {
		GaugeMetrics   map[string]*models.GaugeMetric
		CounterMetrics map[string]*models.CounterMetric
	}
	type args struct {
		Metric *models.GaugeMetric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "gauge positive value",
			fields: fields{
				GaugeMetrics:   make(map[string]*models.GaugeMetric),
				CounterMetrics: make(map[string]*models.CounterMetric),
			},
			args: args{
				Metric: &models.GaugeMetric{
					Name:  "SomeGaugeMetric",
					Type:  0,
					Value: 10.0,
				},
			},
			want: 10.0,
		},
		{
			name: "gauge negative value",
			fields: fields{
				GaugeMetrics:   make(map[string]*models.GaugeMetric),
				CounterMetrics: make(map[string]*models.CounterMetric),
			},
			args: args{
				Metric: &models.GaugeMetric{
					Name:  "SomeGaugeMetric",
					Type:  0,
					Value: -10.0,
				},
			},
			want: -10.0,
		},
		{
			name: "gauge zero value",
			fields: fields{
				GaugeMetrics:   make(map[string]*models.GaugeMetric),
				CounterMetrics: make(map[string]*models.CounterMetric),
			},
			args: args{
				Metric: &models.GaugeMetric{
					Name:  "SomeGaugeMetric",
					Type:  0,
					Value: 0.0,
				},
			},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				gaugeMetrics:   tt.fields.GaugeMetrics,
				counterMetrics: tt.fields.CounterMetrics,
			}
			_ = s.UpdateGauge(tt.args.Metric)

			assert.Equal(t, tt.want, s.gaugeMetrics[tt.args.Metric.Name].Value)
		})
	}
}

func TestStorage_UpdateCounter(t *testing.T) {
	type fields struct {
		GaugeMetrics   map[string]*models.GaugeMetric
		CounterMetrics map[string]*models.CounterMetric
	}
	type args struct {
		Metric *models.CounterMetric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name: "counter positive value",
			fields: fields{
				GaugeMetrics:   make(map[string]*models.GaugeMetric),
				CounterMetrics: make(map[string]*models.CounterMetric),
			},
			args: args{
				Metric: &models.CounterMetric{
					Name:  "SomeCounterMetric",
					Type:  1,
					Value: 10,
				},
			},
			want: 10,
		},
		{
			name: "counter zero value",
			fields: fields{
				GaugeMetrics:   make(map[string]*models.GaugeMetric),
				CounterMetrics: make(map[string]*models.CounterMetric),
			},
			args: args{
				Metric: &models.CounterMetric{
					Name:  "SomeCounterMetric",
					Type:  1,
					Value: 0,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				gaugeMetrics:   tt.fields.GaugeMetrics,
				counterMetrics: tt.fields.CounterMetrics,
			}
			_ = s.UpdateCounter(tt.args.Metric)

			assert.Equal(t, tt.want, s.counterMetrics[tt.args.Metric.Name].Value)
		})
	}
}
