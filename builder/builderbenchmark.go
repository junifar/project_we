package builder

import (
	"github.com/prometheus/common/log"
	"strconv"
	"time"
)

type BenchmarkTimer interface {
	Stop() (timeTaken int64)
	LogTotalTime()
	LogError(err error) (errorMsg string)
}

type Benchmark struct {
	StartTime   time.Time
	Description string
	EndTime     time.Time
}

func NewBenchmark(description string) BenchmarkTimer {
	log.Info("start " + description)

	return &Benchmark{
		StartTime:   time.Now(),
		Description: description,
	}
}

func (benchmark *Benchmark) Stop() (secondsTaken int64) {
	benchmark.EndTime = time.Now()
	return benchmark.EndTime.Sub(benchmark.StartTime).Milliseconds()
}

func (benchmark *Benchmark) LogTotalTime() {
	timeTaken := benchmark.Stop()
	timeTakenStr := strconv.FormatInt(timeTaken, 10) + " ms"
	log.Info("done " + benchmark.Description + ". time: " + timeTakenStr)
}

func (benchmark *Benchmark) LogError(err error) (errorMsg string) {
	errorMsg = "error " + benchmark.Description + " : " + err.Error()
	log.Error(errorMsg)
	return
}
