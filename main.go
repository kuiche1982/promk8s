package main

import (
	"fmt"
	"net/http"
	"time"

	"math/rand"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
