package telemetry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTags(t *testing.T) {

	//
	// Setup
	//

	tags := Tags{}

	tags.Add("tag-1", "1")
	tags.Add("tag-2", "2")
	tags.Add("tag-3", "3")
	tags.Add("tag-4", "4")

	//
	// Verify
	//

	// tags
	assert.Equal(t, len(tags.getValues()), 8)
}
