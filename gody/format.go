package gody

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"sort"
	"unicode/utf8"
)

type FormatTarget struct {
	ddbresult []map[string]interface{}
	format    string
	header    bool
	fields    []string
	cmd       *cobra.Command
}

func Format(target FormatTarget) {
	switch target.format {
	case "ssv":
		toXsv(target, " ")
	case "csv":
		toXsv(target, ",")
	case "tsv":
		toXsv(target, "\t")
	case "json":
		toJson(target)
	}
}

func toXsv(target FormatTarget, delimiter string) {
	w := csv.NewWriter(target.cmd.OutOrStdout())
	delm, _ := utf8.DecodeRuneInString(delimiter)
	w.Comma = delm

	head := make([]string, 0, len(target.ddbresult))
	if len(target.fields) > 0 {
		head = target.fields
	} else {
		// https://qiita.com/hi-nakamura/items/5671eae147ffa68c4466
		// headをユニークなsliceにする
		encountered := map[string]bool{}
		for _, v := range target.ddbresult {
			for k := range v {
				if !encountered[k] {
					encountered[k] = true
					head = append(head, k)
				}
			}
		}
		// headerをsortしておおよそ同じ順に表示されるようにする
		sort.Strings(head)
	}

	if target.header {
		w.Write(head)
		w.Flush()
	}

	var (
		body [][]string
	)
	body_unit := make([]string, len(head))
	for _, v := range target.ddbresult {
		for _, h := range head {
			// 存在しないキーの場合は、値を"_"にする
			_, ok := v[h]
			if ok {
				body_unit = append(body_unit, fmt.Sprint(v[h]))
			} else {
				body_unit = append(body_unit, "_")
			}
		}
		body = append(body, body_unit)
		body_unit = make([]string, len(head))
	}

	for _, b := range body {
		w.Write(b)
		w.Flush()
	}
}

func toJson(target FormatTarget) {
	var jsonString []byte
	if len(target.fields) > 0 {
		m := make(map[string]interface{}, len(target.fields))
		marr := []map[string]interface{}{}
		for _, v := range target.ddbresult {
			for _, f := range target.fields {
				_, ok := v[f]
				if ok {
					m[f] = v[f]
				}
			}
			marr = append(marr, m)
			m = map[string]interface{}{}
		}
		jsonString, _ = json.Marshal(marr)
	} else {
		jsonString, _ = json.Marshal(target.ddbresult)
	}
	target.cmd.Println(string(jsonString))
}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
