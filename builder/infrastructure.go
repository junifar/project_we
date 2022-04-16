package builder

import "project_we/app/pkg/curl"

var httpRequestorCache curl.IHttpRequestor

func BuildHTTPRequestor() curl.IHttpRequestor {
	if httpRequestorCache != nil {
		return httpRequestorCache
	}

	benchmark := NewBenchmark("initializing http requestor")
	defer benchmark.LogTotalTime()

	httpClient := curl.NewHTTPClient()
	httpRequestor := curl.NewHttpRequestor(httpClient)

	httpRequestorCache = httpRequestor

	return httpRequestorCache
}
