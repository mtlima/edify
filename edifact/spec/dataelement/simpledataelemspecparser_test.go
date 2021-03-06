package dataelement

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSpecParser(t *testing.T) {
	p := NewSimpleDataElemSpecParser(fixtureTextCodesSpecMap())
	_, err := p.ParseSpecFile("../../../testdata/EDED.14B_short")
	assert.Nil(t, err)
}

const specLines = `
     1000  Document name                                           [B]

     Desc: Name of a document.

     Repr: an..35
`

func TestParseSpecLines(t *testing.T) {
	p := NewSimpleDataElemSpecParser(fixtureTextCodesSpecMap())
	spec, err := p.ParseSpec(strings.Split(specLines, "\n"))
	assert.Nil(t, err)
	assert.Equal(t, "1000", spec.Id())
	assert.Equal(t, "Document name", spec.Name())
	assert.Equal(t, "Name of a document.", spec.Descr)
	// log.Printf("res: %s", res)
}

const specLinesLongDescr = `
     1000  Document name                                           [B]

     Desc: Name of a document.
           And some more description.

     Repr: an..35
`

func TestParseSpecLinesLongDescr(t *testing.T) {
	p := NewSimpleDataElemSpecParser(fixtureTextCodesSpecMap())
	spec, err := p.ParseSpec(strings.Split(specLinesLongDescr, "\n"))
	assert.Nil(t, err)
	assert.Equal(t, "Name of a document. And some more description.", spec.Descr)
}

func BenchmarkParseSpecLines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := NewSimpleDataElemSpecParser(fixtureTextCodesSpecMap())
		_, err := p.ParseSpecFile("../../../testdata/EDED.14B")
		assert.Nil(b, err)
		// log.Printf("Parsed %d specs\n", len(specs))
	}
}
