package processor

var go2Pg = map[string]string{
	"bool":      "boolean",
	"string":    "text",
	"int":       "integer",
	"int64":     "bigint",
	"float64":   "double precision",
	"time.Time": "TIMESTAMP WITH TIME ZONE",
}
