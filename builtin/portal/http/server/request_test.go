package server

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
)

func TestBuildRequestStructure(t *testing.T) {
	req := httptest.NewRequest("POST", "http://example.com/foo/bar?x=1&x=2&y=3", nil)
	req.URL.RawPath = "/foo/bar"
	req.RemoteAddr = "1.2.3.4:5678"
	req.Header.Add("X-Test", "a")
	req.Header.Add("X-Test", "b")
	req.Header.Set("Content-Type", "text/plain")

	got := buildRequestStructure(req, "payload")

	expected := base.NewStructureFromRecords([]base.StructureRecord{
		base.NewStructureRecord(base.NewTag("method"), base.NewString("POST")),
		base.NewStructureRecord(base.NewTag("path"), base.NewString("/foo/bar")),
		base.NewStructureRecord(base.NewTag("raw-path"), base.NewString("/foo/bar")),
		base.NewStructureRecord(base.NewTag("query"), base.NewConventionalList(
			base.NewStructureFromRecords([]base.StructureRecord{
				base.NewStructureRecord(base.NewTag("key"), base.NewString("x")),
				base.NewStructureRecord(base.NewTag("value"), base.NewString("1")),
			}),
			base.NewStructureFromRecords([]base.StructureRecord{
				base.NewStructureRecord(base.NewTag("key"), base.NewString("x")),
				base.NewStructureRecord(base.NewTag("value"), base.NewString("2")),
			}),
			base.NewStructureFromRecords([]base.StructureRecord{
				base.NewStructureRecord(base.NewTag("key"), base.NewString("y")),
				base.NewStructureRecord(base.NewTag("value"), base.NewString("3")),
			}),
		)),
		base.NewStructureRecord(base.NewTag("query-map"), base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(base.NewString("x"), base.NewString("1")),
			base.NewStructureRecord(base.NewString("y"), base.NewString("3")),
		})),
		base.NewStructureRecord(base.NewTag("headers"), base.NewConventionalList(
			base.NewStructureFromRecords([]base.StructureRecord{
				base.NewStructureRecord(base.NewTag("key"), base.NewString("Content-Type")),
				base.NewStructureRecord(base.NewTag("value"), base.NewString("text/plain")),
			}),
			base.NewStructureFromRecords([]base.StructureRecord{
				base.NewStructureRecord(base.NewTag("key"), base.NewString("X-Test")),
				base.NewStructureRecord(base.NewTag("value"), base.NewString("a")),
			}),
			base.NewStructureFromRecords([]base.StructureRecord{
				base.NewStructureRecord(base.NewTag("key"), base.NewString("X-Test")),
				base.NewStructureRecord(base.NewTag("value"), base.NewString("b")),
			}),
		)),
		base.NewStructureRecord(base.NewTag("headers-map"), base.NewStructureFromRecords([]base.StructureRecord{
			base.NewStructureRecord(base.NewString("Content-Type"), base.NewString("text/plain")),
			base.NewStructureRecord(base.NewString("X-Test"), base.NewString("a")),
		})),
		base.NewStructureRecord(base.NewTag("body"), base.NewString("payload")),
		base.NewStructureRecord(base.NewTag("remote-addr"), base.NewString("1.2.3.4:5678")),
		base.NewStructureRecord(base.NewTag("path-segments"), base.NewConventionalList(
			base.NewString("foo"),
			base.NewString("bar"),
		)),
		base.NewStructureRecord(base.NewTag("path-params"), base.NewStructureFromRecords(nil)),
	})

	assert.Equal(t, expected, got)
}
