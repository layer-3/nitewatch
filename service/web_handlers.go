package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type healthStatus string

const (
	statusHealthy  healthStatus = "healthy"
	statusDegraded healthStatus = "degraded"
)

type healthResponse struct {
	Status       healthStatus           `json:"status"`
	Service      string                 `json:"service"`
	Timestamp    string                 `json:"timestamp"`
	Dependencies map[string]depCheck    `json:"dependencies,omitempty"`
}

type depCheck struct {
	Status healthStatus `json:"status"`
	Error  string       `json:"error,omitempty"`
}

func newHealthResponse() healthResponse {
	return healthResponse{
		Status:       statusHealthy,
		Service:      "nitewatch",
		Timestamp:    time.Now().UTC().Format(time.RFC3339),
		Dependencies: make(map[string]depCheck),
	}
}

func (r *healthResponse) addDependency(name string, check depCheck) {
	r.Dependencies[name] = check
	if check.Status == statusDegraded && r.Status == statusHealthy {
		r.Status = statusDegraded
	}
}

func (r *healthResponse) httpStatusCode() int {
	return http.StatusOK
}

func attachWebHandlers(svc *Service) {
	svc.web.Engine.GET("/health", getHealth(svc))
	svc.web.Engine.GET("/health/live", livenessHandler())
	svc.web.Engine.GET("/health/ready", readinessHandler(svc))
}

func getHealth(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := "healthy"
		if !svc.IsWorkerReady() {
			status = "degraded"
		}
		c.JSON(http.StatusOK, gin.H{
			"status": status,
			"worker": svc.IsWorkerReady(),
		})
	}
}

func livenessHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := newHealthResponse()
		c.JSON(http.StatusOK, response)
	}
}

func readinessHandler(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		response := newHealthResponse()

		if svc.IsWorkerReady() {
			response.addDependency("worker", depCheck{
				Status: statusHealthy,
			})
		} else {
			response.addDependency("worker", depCheck{
				Status: statusDegraded,
				Error:  "worker not ready",
			})
		}

		c.JSON(response.httpStatusCode(), response)
	}
}
