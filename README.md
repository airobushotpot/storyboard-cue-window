# Storyboard Cue Window

Group timed storyboard markers into fixed-length clips. It returns the
occupied clip indices and the cues that belong to each clip. This is useful
when dividing a long storyboard into shorter image-to-video jobs.

## Usage

```go
import (
    "time"

    cuewindow "github.com/airobushotpot/storyboard-cue-window"
)

segments, err := cuewindow.Group([]cuewindow.Cue{
    {Offset: 0, Label: "Wide shot"},
    {Offset: 5 * time.Second, Label: "Close-up"},
}, 5*time.Second)
```

Supply cues in chronological order. A cue exactly on a segment boundary
belongs to the following segment. Cue offsets remain relative to the start
of the entire storyboard. The module handles timing and leaves media
processing to the caller. Go 1.22 or newer. MIT license.

Related workflow at [Image to Video AI](https://imagetovideoai.tools/).
