package tsdb

import "github.com/grafana/grafana/pkg/components/simplejson"

type Query struct {
	RefId         string
	Query         string
	Model         *simplejson.Json
	Depends       []string
	DataSource    *DataSourceInfo
	Results       []*TimeSeries
	Exclude       bool
	MaxDataPoints int64
	IntervalMs    int64
}

type QuerySlice []*Query

type Request struct {
	TimeRange TimeRange
	Queries   QuerySlice
}

type Response struct {
	BatchTimings []*BatchTiming
	Results      map[string]*QueryResult
}

type DataSourceInfo struct {
	Id                int64
	Name              string
	PluginId          string
	Url               string
	Password          string
	User              string
	Database          string
	BasicAuth         bool
	BasicAuthUser     string
	BasicAuthPassword string
}

type BatchTiming struct {
	TimeElapsed int64
}

type BatchResult struct {
	Error        error
	QueryResults map[string]*QueryResult
	Timings      *BatchTiming
}

type QueryResult struct {
	Error  error
	RefId  string
	Series TimeSeriesSlice
}

type TimeSeries struct {
	Name   string        `json:"name"`
	Points [][2]*float64 `json:"points"`
}

type TimeSeriesSlice []*TimeSeries

func NewTimeSeries(name string, points [][2]*float64) *TimeSeries {
	return &TimeSeries{
		Name:   name,
		Points: points,
	}
}