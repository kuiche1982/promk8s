kubens spotmax-maxcloud

kubectl apply -f -<<EOF
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  labels:
    app: prometheus-pushgateway
    release: kube-prometheus-stack
spec:
  endpoints:
    - path: /metrics
      port: http
  namespaceSelector:
    any: true
  selector:
    matchLabels:
      app: prometheus-pushgateway
EOF

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install engineplush prometheus-community/prometheus-pushgateway
engineplush-prometheus-pushgateway.spotmax-maxcloud.svc.cluster.local

echo "sample_metric 3.14" | curl --data-binary @- http://engineplush-prometheus-pushgateway.spotmax-maxcloud.svc.cluster.local:9091/metrics/job/sample_job


Go Sample
completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "db_backup_last_completion_timestamp_seconds",
			Help: "The timestamp of the last successful completion of a DB backup.",
		})
		pusher := push.New(url, "jobname1")
		completionTime.SetToCurrentTime()
		if err := pusher.
			Collector(completionTime).
			Grouping("db", "customers").
			Push(); err != nil {
			fmt.Println("Could not push completion time to Pushgateway:", err)
		}
		metrics2 := prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: service,
			Name:      "histogram",
			Help:      "histogram with labels Component, Feature, Status",
		}, []string{"Component", "Feature", "EndPoint", "Status"})
		reqOk2 := metrics2.WithLabelValues(comp, feature, endpoint, "statusOK")
		reqError2 := metrics2.WithLabelValues(comp, feature, endpoint, "statusError")
		for i := 0; i < 1000; i++ {

			switch r.Int() % 2 {
			case 0:
				reqOk2.Observe(r.NormFloat64()*10 + 100)
			default:
				reqError2.Observe(r.NormFloat64()*10 + 100)
			}
			reqOk.Observe(1)
		}
		if err := pusher.
			Collector(metrics2).
			Push(); err != nil {
			fmt.Println("Could not push completion time to Pushgateway:", err)
		}

{container="pushgateway", db="customers", endpoint="http", exported_job="jobname1", instance="172.30.1.228:9091", job="engineplush-prometheus-pushgateway", namespace="spotmax-maxcloud", pod="engineplush-prometheus-pushgateway-678846f87d-djcx5", service="engineplush-prometheus-pushgateway"}
0
{container="pushgateway", db="customers", endpoint="http", exported_job="sample_metrics3", instance="172.30.1.228:9091", job="engineplush-prometheus-pushgateway", namespace="spotmax-maxcloud", pod="engineplush-prometheus-pushgateway-678846f87d-djcx5", service="engineplush-prometheus-pushgateway"}
0


python sample Node

pip install prometheus-client
https://pypi.org/project/prometheus-client/
https://github.com/prometheus/client_python#exporting-to-a-pushgateway

exporter textfile collector
The textfile collector allows machine-level statistics to be exported out via the Node exporter.

This is useful for monitoring cronjobs, or for writing cronjobs to expose metrics about a machine system that the Node exporter does not support or would not make sense to perform at every scrape (for example, anything involving subprocesses).

from prometheus_client import CollectorRegistry, Gauge, write_to_textfile

registry = CollectorRegistry()
g = Gauge('raid_status', '1 if raid array is okay', registry=registry)
g.set(1)
write_to_textfile('/configured/textfile/path/raid.prom', registry)
A separate registry is used, as the default registry may contain other metrics such as those from the Process Collector.

Exporting to a Pushgateway
The Pushgateway allows ephemeral and batch jobs to expose their metrics to Prometheus.

from prometheus_client import CollectorRegistry, Gauge, push_to_gateway

registry = CollectorRegistry()
g = Gauge('job_last_success_unixtime', 'Last time a batch job successfully finished', registry=registry)
g.set_to_current_time()
push_to_gateway('localhost:9091', job='batchA', registry=registry)