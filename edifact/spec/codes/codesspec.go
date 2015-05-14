package codes

import (
	"fmt"
	"github.com/bbiskup/edify/edifact/util"
	"strings"
)

type CodeSpecMap map[string]*CodeSpec

// EDIFACT Codes as defined e.g. in UNCL.14B
type CodesSpec struct {
	Id          string
	Name        string
	Description string
	codeSpecMap CodeSpecMap
}

type CodesSpecMap map[string]*CodesSpec

func (s *CodesSpec) Contains(code string) bool {
	return s.codeSpecMap[code] != nil
}

func (s *CodesSpec) String() string {
	specsStrs := []string{}
	for _, spec := range s.codeSpecMap {
		specsStrs = append(specsStrs, fmt.Sprintf("\t%s", spec.String()))
	}
	codeSpecsStr := strings.Join(specsStrs, "\n")
	descriptionStr := util.Ellipsis(s.Description, maxDescrDisplayLen)

	return fmt.Sprintf("%s %s %s\n%s", s.Id, s.Name, descriptionStr, codeSpecsStr)
}

func (s *CodesSpec) Len() int {
	if s.codeSpecMap == nil {
		return 0
	} else {
		return len(s.codeSpecMap)
	}
}

func NewCodesSpec(id string, name string, description string, codeSpecs []*CodeSpec) *CodesSpec {
	codeSpecMap := CodeSpecMap{}
	for _, codeSpec := range codeSpecs {
		codeSpecMap[codeSpec.Id] = codeSpec
	}

	return &CodesSpec{
		Id:          id,
		Name:        name,
		Description: description,
		codeSpecMap: codeSpecMap,
	}
}
