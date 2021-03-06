package segment

import (
	"fmt"
)

type SegSpecMap map[string]*SegSpec

// Regular implementation of SegSpecProvider for production
type SegSpecProviderImpl struct {
	segSpecs SegSpecMap
}

func NewSegSpecProviderImpl(segSpecs SegSpecMap) *SegSpecProviderImpl {
	return &SegSpecProviderImpl{segSpecs}
}

func (p *SegSpecProviderImpl) Ids() []string {
	result := make([]string, 0, len(p.segSpecs))
	for id, _ := range p.segSpecs {
		result = append(result, id)
	}
	return result
}

func (p *SegSpecProviderImpl) Get(id string) *SegSpec {
	result := p.segSpecs[id]
	if result == nil {
		// e.g. UNH, UNT are not defined in UNCE specs, because they
		// are not part of the release cycle. Instead, they are defined
		// in part 1 of ISO9735 (file testdata/r1241.txt)
		//log.Printf("Missing segment spec:1 '%s'", id)

		if IsUnValidatedSegment(id) {
			return NewSegSpec(
				id, fmt.Sprintf("missing-%s", id),
				"dummy_function", nil)
		} else {
			return nil
		}
	} else {
		return result
	}
}

func (p *SegSpecProviderImpl) Len() int {
	return len(p.segSpecs)
}
