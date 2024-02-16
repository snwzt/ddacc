package helper

import (
	"fmt"
	"net/http"
	"time"

	"snwzt/ddacc/data"
)

func ReadWithTimeout(resp *http.Response, buf []byte, timeout time.Duration) (int, error) {
	resultChan := make(chan data.ReadData)

	go func() {
		n, err := resp.Body.Read(buf)
		resultChan <- data.ReadData{
			N:   n,
			Err: err,
		}
		close(resultChan)
	}()

	select {
	case result := <-resultChan:
		return result.N, result.Err
	case <-time.After(timeout):
		return 0, fmt.Errorf("reading response body timed out")
	}
}
