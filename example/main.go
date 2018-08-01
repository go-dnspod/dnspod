package main

import (
	"encoding/json"
	"fmt"
	"go-dnspod/dnspod"
)

// type Record struct {
// 	ID            int64   `json:"id,string"`
// 	TTL           int     `json:"ttl,string"`
// 	Value         string  `json:"value"`
// 	Enabled       Enabled `json:"enabled,string"`
// 	UpdatedOn     Time    `json:"updated_on"`
// 	Name          string  `json:"name"`
// 	Line          string  `json:"line"`
// 	LineID        string  `json:"line_id"`
// 	Type          string  `json:"type"`
// 	Weight        int     `json:"weight"`
// 	MonitorStatus string  `json:"monitor_status"`
// 	Remark        string  `json:"remark"`
// 	UseAQB        Yes     `json:"use_aqb"`
// 	MX            int     `json:"mx,string"`
// }

func main() {
	testJson := `{"id":"325137273","ttl":"600","value":"47.52.74.252","enabled":"1","status":"enabled","updated_on":"2017-10-04 12:10:24","name":"*","line":"默认","line_id":"0","type":"A","weight":null,"monitor_status":"","remark":"","use_aqb":"no","mx":"0"}`
	var r dnspod.Record
	fmt.Println(json.Unmarshal([]byte(testJson), &r), r)

}
