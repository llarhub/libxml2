package libxml2_test

import (
	"testing"
	"unsafe"

	"github.com/goplus/lib/c"
	"github.com/llarhub/libxml2"
)

// Adapted from https://github.com/goplus/llcppg/blob/v0.8.0/_llcppgtest/libxml2/demo/parsexml/parsexml.go.
func TestParseXML(t *testing.T) {
	libxml2.InitParser()
	defer libxml2.CleanupParser()

	xml := "<?xml version='1.0'?><root><person><name>Alice</name><age>25</age></person></root>"
	doc := libxml2.ReadMemory((*c.Char)(unsafe.Pointer(unsafe.StringData(xml))), c.Int(len(xml)), nil, nil, 0)
	if doc == nil {
		t.Fatal("failed to parse XML")
	}
	defer doc.Free()

	root := doc.RootElement()
	if root == nil {
		t.Fatal("missing root element")
	}
	if got := c.GoString((*c.Char)(unsafe.Pointer(root.Name))); got != "root" {
		t.Fatalf("root element = %q, want %q", got, "root")
	}
	t.Log("root element: root")
}
