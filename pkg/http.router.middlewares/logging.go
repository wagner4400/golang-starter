package middlewares

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
	"time"
)

// LoggingRequest holds the information from the request
type LoggingRequest struct {
	Status   int           `json:"status"`
	Method   string        `json:"method"`
	Path     string        `json:"path"`
	Latency  time.Duration `json:"latency"`
	Request  string        `json:"request"`
	Response string        `json:"response"`
}

// Logging is a middleware to log all the requests and responses
func Logging(skippedPaths ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()

		// Copy request body
		requestBC := requestBodyCopier{body: new(bytes.Buffer), ReadCloser: c.Request.Body}
		c.Request.Body = &requestBC

		// Copy response body
		responseBC := responseBodyCopier{body: new(bytes.Buffer), ResponseWriter: c.Writer}
		c.Writer = &responseBC

		c.Next()

		latency := time.Since(start)

		requestURI := c.Request.RequestURI
		for _, path := range skippedPaths {
			if strings.HasPrefix(requestURI, path) {
				return
			}
		}

		requestBytes, err := io.ReadAll(requestBC.body)
		if err != nil {
			fmt.Printf("Error extracting request body to logging: %s \r\n", err)
		}

		responseBytes, err := io.ReadAll(responseBC.body)
		if err != nil {
			fmt.Printf("Error extracting request body to logging: %s \r\n", err)
		}

		log(c, LoggingRequest{
			Status:   c.Writer.Status(),
			Method:   c.Request.Method,
			Path:     c.Request.RequestURI,
			Latency:  latency,
			Request:  string(requestBytes),
			Response: string(responseBytes),
		})
	}
}

const logFormat = "status &d | method %s | path %s | latency &s | request %s | response %s"

func log(c *gin.Context, log LoggingRequest) {

	//retrievedLogger := c.Value("logger_key").(*logger.Log)
	//
	//retrievedLogger.Info(fmt.Sprintf(logFormat,
	//	log.Status,
	//	log.Method,
	//	log.Path,
	//	log.Latency,
	//	log.Request,
	//	log.Response),
	//)
}

type responseBodyCopier struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseBodyCopier) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type requestBodyCopier struct {
	io.ReadCloser
	body *bytes.Buffer
}

func (w *requestBodyCopier) Read(b []byte) (int, error) {
	read, err := w.ReadCloser.Read(b)
	if read > 0 {
		w.body.Write(b[:read])
	}
	return read, err
}
