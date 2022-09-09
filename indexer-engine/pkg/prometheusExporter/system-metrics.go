package prometheusexporter

import (
	"fmt"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func SystemMetrices(e *echo.Echo) {
	fmt.Println("[Prometheus] Exposing whole server metrics in endpoint `/metrics`")
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
}
