package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
	"github.com/leaanthony/synx"
)

// SpinnerGroup is a group of Spinners
type SpinnerGroup struct {
	sync.Mutex
	sync.WaitGroup
	spinners        []*Spinner
	frames          []string
	currentFrameIdx int
	successSymbol   string
	errorSymbol     string
	running         bool
	drawn           bool
}

// At returns the Spinner at given 0-based index
func (g *SpinnerGroup) At(idx int) *Spinner {
	return g.spinners[idx]
}

// Start the spinners
func (g *SpinnerGroup) Start() {
	g.Lock()
	defer g.Unlock()

	if g.running {
		return
	}
	g.running = true

	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		for g.running {
			select {
			case <-ticker.C:
				g.redraw()
			}
		}
	}()
}

// Stop the spinners
func (g *SpinnerGroup) Stop() {
	g.Lock()
	defer g.Unlock()
	g.running = false
}

// Wait for all spinners to finish
func (g *SpinnerGroup) Wait() {
	g.WaitGroup.Wait()
	g.Stop()
}

func (g *SpinnerGroup) redraw() {
	g.Lock()
	defer g.Unlock()
	if !g.running {
		return
	}
	if g.drawn {
		fmt.Print(cursor.MoveUp(len(g.spinners)))
	}
	for _, spinner := range g.spinners {
		fmt.Print(cursor.ClearEntireLine())
		fmt.Println(spinner.sprint())
	}
	g.currentFrameIdx = (g.currentFrameIdx + 1) % len(g.frames)
	g.drawn = true
}

func (g *SpinnerGroup) currentFrame() string {
	return g.frames[g.currentFrameIdx]
}

// NewSpinnerGroup creates a SpinnerGroup
func NewSpinnerGroup(size int) *SpinnerGroup {
	group := &SpinnerGroup{
		spinners:        make([]*Spinner, size),
		frames:          []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
		currentFrameIdx: 0,
		successSymbol:   "✓",
		errorSymbol:     "✗",
		running:         false,
		drawn:           false,
	}
	for i := 0; i < size; i++ {
		group.spinners[i] = &Spinner{
			message: synx.NewString(fmt.Sprintf("Spinner #%d", i+1)),
			status:  synx.NewInt(runningStatus),
			group:   group,
		}
	}
	group.Add(size)
	return group
}
