package server

import (
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/SSripilaipong/muto/core/base"
)

func newRequestObject(r *http.Request, requestClass base.Class) (base.Object, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	reqStruct := buildRequestStructure(r, string(bodyBytes))
	return base.NewOneLayerObject(requestClass, reqStruct), nil
}

func buildRequestStructure(r *http.Request, body string) base.Structure {
	rawPath := r.URL.RawPath
	query := r.URL.Query()
	headers := r.Header

	queryList := buildKeyValueList(query)
	queryMap := buildKeyValueMap(query)
	headersList := buildKeyValueList(headers)
	headersMap := buildKeyValueMap(headers)
	pathSegments := buildPathSegments(r.URL.Path)

	return base.NewStructureFromRecords([]base.StructureRecord{
		base.NewStructureRecord(base.NewTag("method"), base.NewString(r.Method)),
		base.NewStructureRecord(base.NewTag("path"), base.NewString(r.URL.Path)),
		base.NewStructureRecord(base.NewTag("raw-path"), base.NewString(rawPath)),
		base.NewStructureRecord(base.NewTag("query"), queryList),
		base.NewStructureRecord(base.NewTag("query-map"), queryMap),
		base.NewStructureRecord(base.NewTag("headers"), headersList),
		base.NewStructureRecord(base.NewTag("headers-map"), headersMap),
		base.NewStructureRecord(base.NewTag("body"), base.NewString(body)),
		base.NewStructureRecord(base.NewTag("remote-addr"), base.NewString(r.RemoteAddr)),
		base.NewStructureRecord(base.NewTag("path-segments"), pathSegments),
		base.NewStructureRecord(base.NewTag("path-params"), base.NewStructureFromRecords(nil)),
	})
}

func buildKeyValueList(values map[string][]string) base.Object {
	keys := sortedKeys(values)
	var nodes []base.Node
	for _, key := range keys {
		for _, value := range values[key] {
			nodes = append(nodes, base.NewStructureFromRecords([]base.StructureRecord{
				base.NewStructureRecord(base.NewTag("key"), base.NewString(key)),
				base.NewStructureRecord(base.NewTag("value"), base.NewString(value)),
			}))
		}
	}
	return base.NewConventionalList(nodes...)
}

func buildKeyValueMap(values map[string][]string) base.Structure {
	keys := sortedKeys(values)
	var records []base.StructureRecord
	for _, key := range keys {
		value := ""
		if len(values[key]) > 0 {
			value = values[key][0]
		}
		records = append(records, base.NewStructureRecord(base.NewString(key), base.NewString(value)))
	}
	return base.NewStructureFromRecords(records)
}

func buildPathSegments(path string) base.Object {
	trimmed := strings.Trim(path, "/")
	if trimmed == "" {
		return base.NewConventionalList()
	}
	parts := strings.Split(trimmed, "/")
	nodes := make([]base.Node, 0, len(parts))
	for _, part := range parts {
		nodes = append(nodes, base.NewString(part))
	}
	return base.NewConventionalList(nodes...)
}

func sortedKeys(values map[string][]string) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
