package main

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/barchart"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/donut"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/mum4k/termdash/widgets/textinput"
)

// redrawInterval is how often termdash redraws the screen.
const redrawInterval = 250 * time.Millisecond

// widgets holds the widgets used by this demo.
type widgets struct {
	segDist  *segmentdisplay.SegmentDisplay
	input    *textinput.TextInput
	rollT    *text.Text
	spGreen  *sparkline.SparkLine
	spRed    *sparkline.SparkLine
	gauge    *gauge.Gauge
	heartLC  *linechart.LineChart
	barChart *barchart.BarChart
	donut    *donut.Donut
	leftB    *button.Button
	rightB   *button.Button
	sineLC   *linechart.LineChart

	buttons *layoutButtons
}

// newWidgets creates all widgets used by this demo.
func newWidgets(ctx context.Context, c *container.Container) (*widgets, error) {
	updateText := make(chan string)
	sd, err := newSegmentDisplay(ctx, updateText)
	if err != nil {
		return nil, err
	}

	input, err := newTextInput(updateText)
	if err != nil {
		return nil, err
	}

	rollT, err := newRollText(ctx)
	if err != nil {
		return nil, err
	}
	spGreen, spRed, err := newSparkLines(ctx)
	if err != nil {
		return nil, err
	}
	g, err := newGauge(ctx)
	if err != nil {
		return nil, err
	}

	heartLC, err := newHeartbeat(ctx)
	if err != nil {
		return nil, err
	}

	bc, err := newBarChart(ctx)
	if err != nil {
		return nil, err
	}

	don, err := newDonut(ctx)
	if err != nil {
		return nil, err
	}

	leftB, rightB, sineLC, err := newSines(ctx)
	if err != nil {
		return nil, err
	}
	return &widgets{
		segDist:  sd,
		input:    input,
		rollT:    rollT,
		spGreen:  spGreen,
		spRed:    spRed,
		gauge:    g,
		heartLC:  heartLC,
		barChart: bc,
		donut:    don,
		leftB:    leftB,
		rightB:   rightB,
		sineLC:   sineLC,
	}, nil
}

// layoutType represents the possible layouts the buttons switch between.
type layoutType int

const (
	// layoutAll displays all the widgets.
	layoutAll layoutType = iota
	// layoutText focuses onto the text widget.
	layoutText
)

// rootID is the ID assigned to the root container.
const rootID = "root"

func main() {
	t, err := termbox.New(termbox.ColorMode(terminalapi.ColorMode256))
	if err != nil {
		panic(err)
	}
	defer t.Close()

	c, err := container.New(t, container.ID(rootID))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	w, err := newWidgets(ctx, c)
	if err != nil {
		panic(err)
	}
	lb, err := newLayoutButtons(c, w)
	if err != nil {
		panic(err)
	}
	w.buttons = lb

	gridOpts, err := gridLayout(w, layoutAll) // equivalent to contLayout(w)
	if err != nil {
		panic(err)
	}

	if err := c.Update(rootID, gridOpts...); err != nil {
		panic(err)
	}

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == keyboard.KeyEsc || k.Key == keyboard.KeyCtrlC {
			cancel()
		}
	}
	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(redrawInterval)); err != nil {
		panic(err)
	}
}

// periodic executes the provided closure periodically every interval.
// Exits when the context expires.
func periodic(ctx context.Context, interval time.Duration, fn func() error) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := fn(); err != nil {
				panic(err)
			}
		case <-ctx.Done():
			return
		}
	}
}

// newTextInput creates a new TextInput field that changes the text on the
// SegmentDisplay.
func newTextInput(updateText chan<- string) (*textinput.TextInput, error) {
	input, err := textinput.New(
		textinput.Label("Change text to: ", cell.FgColor(cell.ColorBlue)),
		textinput.MaxWidthCells(20),
		textinput.PlaceHolder("enter any text"),
		textinput.OnSubmit(func(text string) error {
			updateText <- text
			return nil
		}),
		textinput.ClearOnSubmit(),
	)
	if err != nil {
		return nil, err
	}
	return input, err
}

// newSegmentDisplay creates a new SegmentDisplay that initially shows the
// Termdash name. Shows any text that is sent over the channel.
func newSegmentDisplay(ctx context.Context, updateText <-chan string) (*segmentdisplay.SegmentDisplay, error) {
	sd, err := segmentdisplay.New()
	if err != nil {
		return nil, err
	}

	colors := []cell.Color{
		cell.ColorBlue,
		cell.ColorRed,
		cell.ColorYellow,
		cell.ColorBlue,
		cell.ColorGreen,
		cell.ColorRed,
		cell.ColorGreen,
		cell.ColorRed,
	}

	text := "Termdash"
	step := 0

	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				state := textState(text, sd.Capacity(), step)
				var chunks []*segmentdisplay.TextChunk
				for i := 0; i < sd.Capacity(); i++ {
					if i >= len(state) {
						break
					}

					color := colors[i%len(colors)]
					chunks = append(chunks, segmentdisplay.NewChunk(
						string(state[i]),
						segmentdisplay.WriteCellOpts(cell.FgColor(color)),
					))
				}
				if len(chunks) == 0 {
					continue
				}
				if err := sd.Write(chunks); err != nil {
					panic(err)
				}
				step++

			case t := <-updateText:
				text = t
				sd.Reset()
				step = 0

			case <-ctx.Done():
				return
			}
		}
	}()
	return sd, nil
}