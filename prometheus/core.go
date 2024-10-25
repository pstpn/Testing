package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	hist = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "Testim",
		Help: "WTF",
	}, []string{"Oksd", "Fhsgds"})
)

type metricsArr struct {
	Data []float64 `json:"data"`
}

func main() {
	e := echo.New()
	e.Use(echoprometheus.NewMiddleware("testim"))

	e.GET("/metrics", echoprometheus.NewHandler())
	e.POST("/metrics/push/hist", func(c echo.Context) error {
		var req metricsArr

		err := c.Bind(&req)
		if err != nil {
			return err
		}

		for _, i := range req.Data {
			hist.WithLabelValues("Oksd", "GG").Observe(i)
		}

		return c.JSON(http.StatusOK, "OK")
	})

	e.Logger.Fatal(e.Start(":8081"))
}
