package postgres

const (
	getCounterMetricQuery = `SELECT name, sum(value) FROM counter_metrics WHERE name = $1
		GROUP BY name;`
	getGaugeMetricQuery = `SELECT name, sum(value) FROM gauge_metrics
		GROUP BY name HAVING name = $1;`
	insertGaugeMetricQuery   = `INSERT INTO gauge_metrics (name, value) VALUES ($1, $2);`
	insertCounterMetricQuery = `INSERT INTO counter_metrics (name, value) VALUES ($1, $2);`
	getGaugeMetricsQuery     = `SELECT name, value FROM gauge_metrics gm1 WHERE updated_at  = (
		SELECT MAX(updated_at)
		FROM gauge_metrics gm2
		WHERE gm2.name = gm1.name
	);`
	getCounterMetricsQuery = `SELECT name, value FROM counter_metrics cm1 WHERE updated_at  = (
		SELECT MAX(updated_at)
		FROM counter_metrics cm2
		WHERE cm2.name = cm1.name
	);`

	insertGaugeMetricQueryName   = "insertGaugeMetricQuery"
	insertCounterMetricQueryName = "insertCounterMetricQuery"
)
