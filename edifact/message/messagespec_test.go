package message

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	s := NewMessageSpec(
		"testid", "testname",
		"testversion", "testrelease",
		"testcontragency", "testrevision",
		time.Date(2015, time.January, 15, 0, 0, 0, 0, time.UTC),
		"testsource",
	)

	assert.Equal(t, s.String(), "Message testid (testname)", "String() incorrect")
}