package dataelement

import (
	"log"
	"strings"
	"testing"
)

func TestSpecParser(t *testing.T) {
	p := NewSimpleDataElementSpecParser()
	_, err := p.ParseSpecFile("../../testdata/EDED.14B_short")
	if err != nil {
		t.Fatalf("Parse error: %s", err)
	}
}

const specLines = `
     1000  Document name                                           [B]

     Desc: Name of a document.

     Repr: an..35
`

func TestParseSpecLines(t *testing.T) {
	p := NewSimpleDataElementSpecParser()
	res, err := p.ParseSpec(strings.Split(specLines, "\n"))
	if err != nil {
		t.Fatalf("Failed to parse specLines", err)
	}
	log.Printf("res: %s", res)
}

func BenchmarkParseSpecLines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := NewSimpleDataElementSpecParser()
		specs, err := p.ParseSpecFile("../../testdata/EDED.14B")
		if err != nil {
			log.Printf("Parse error: %s\n", err)
			return
		}
		log.Printf("Parsed %d specs\n", len(specs))
	}
}