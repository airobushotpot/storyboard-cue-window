// Package cuewindow groups storyboard cues into fixed-duration segments.
package cuewindow

import (
	"fmt"
	"time"
)

// Cue marks an event at an offset from the beginning of the storyboard.
type Cue struct {
	Offset time.Duration
	Label  string
}

// Segment contains input cues in one half-open timeline interval.
type Segment struct {
	Index int64
	Cues  []Cue
}

// Group returns the occupied segments in ascending order. Input cues must
// already be in chronological order. Exact boundary cues enter the next
// segment, and the original cue offsets are preserved.
func Group(cues []Cue, segmentLength time.Duration) ([]Segment, error) {
	if segmentLength <= 0 {
		return nil, fmt.Errorf("segment length must be positive")
	}

	var segments []Segment
	var previous time.Duration

	for index, cue := range cues {
		if cue.Offset < previous {
			return nil, fmt.Errorf("cue %d has a negative or out-of-order offset", index)
		}
		previous = cue.Offset

		segmentIndex := int64(cue.Offset / segmentLength)
		last := len(segments) - 1
		if last < 0 || segments[last].Index != segmentIndex {
			segments = append(segments, Segment{Index: segmentIndex})
			last++
		}
		segments[last].Cues = append(segments[last].Cues, cue)
	}
	return segments, nil
}
