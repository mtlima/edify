package msg

import (
	"bytes"
	"strconv"
)

// A segment that is repeated 1 to n times
type RepeatSegmentGroup struct {
	groups []*SegmentGroup
}

// From Interface RepeatMsgPart
func (g *RepeatSegmentGroup) Count() int {
	return len(g.groups)
}

// From SegmentOrGroup
func (g *RepeatSegmentGroup) Id() string {
	return g.groups[0].Id()
}

func (g *RepeatSegmentGroup) Last() *SegmentGroup {
	return g.groups[len(g.groups)-1]
}

func (g *RepeatSegmentGroup) AddSegmentGroup(group *SegmentGroup) {
	g.groups = append(g.groups, group)
}

func (g *RepeatSegmentGroup) Dump(indent int) string {
	var buf bytes.Buffer
	indentStr := getIndentStr(indent)
	for repeat, group := range g.groups {
		buf.WriteString(
			indentStr + "[" + strconv.FormatInt(int64(repeat), 10) + "] " +
				group.Dump(indent))
	}
	return buf.String()
}

func NewRepeatSegmentGroup(groups ...*SegmentGroup) *RepeatSegmentGroup {
	return &RepeatSegmentGroup{
		groups,
	}
}
