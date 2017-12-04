package gody

import (
	"os"
	"testing"
	"time"
)

var testData = FormatTarget{
	ddbresult: []map[string]interface{}{
		{"order_id": 1234},
		{"user_name": "hoge"},
		{"user_id": 2345},
		{"purchased_detail": map[string]interface{}{
			"item_id": 5678,
			"name":    "testItem1",
			"price":   1296,
		},
		},
		{"purchased_at": "2017-01-01 17:25:56"},
		{"created_at": time.Now()},
	},
}

func Test_toJson(t *testing.T){
	toJson(testData, os.Stdout)
}

func Test_deprecatedToJson(t *testing.T){
	deprecatedToJson(testData, os.Stdout)
}

func Benchmark_toJson(b *testing.B) {
	out, _ := os.Open(os.DevNull)
	for i := 0; i < b.N; i++ {
		toJson(testData, out)
	}
}

func Benchmark_deprecatedToJson(b *testing.B) {
	out, _ := os.Open(os.DevNull)
	for i := 0; i < b.N; i++ {
		deprecatedToJson(testData, out)
	}
}
