package main

import (
	"fmt"
	"net/http"
	"time"

	"math/rand"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	service      = "DemoSvc"
	histogramVec = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: service,
		Name:      "histogram",
		Help:      "histogram with labels Component, Feature, Status",
	}, []string{"Component", "Feature", "EndPoint", "Status"})
	comp            = "Demo_Identity"
	feature         = "Demo_type"
	endpoint        = "DemoEP"
	reqOk           = histogramVec.WithLabelValues(comp, feature, endpoint, "statusOK")
	reqError        = histogramVec.WithLabelValues(comp, feature, endpoint, "statusError")
	r               = rand.New(rand.NewSource(time.Now().UnixMilli()))
	promHttpHandler = promhttp.Handler()
)

func main() {
	// mode := os.Getenv("PROM_MODE")
	// url := os.Getenv("PROM_GATEWAY")
	mode := "push"
	url := "http://localhost:53394/"
	if mode == "push" {
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
		return
	}
	http.Handle("/metrics", promHttpHandler)
	http.HandleFunc("/hello", hello)
	fmt.Println(http.ListenAndServe(":8088", nil))

}

func hello(res http.ResponseWriter, req *http.Request) {
	switch r.Int() % 2 {
	case 0:
		reqOk.Observe(1)
	default:
		reqError.Observe(1)
	}
	reqOk.Observe(1)
	res.Write([]byte("abcd"))
}
