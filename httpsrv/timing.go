package httpsrv

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type timingItem struct {
	name     string
	start    time.Time
	duration time.Duration
}

type Timing struct {
	curr     *timingItem
	currAgg  *timingItem
	items    []*timingItem
	aggItems map[string]time.Duration
}

func NewTiming() *Timing {
	return &Timing{
		aggItems: map[string]time.Duration{},
	}
}

func (t *Timing) Start(name string) {
	now := time.Now()
	if t.curr != nil {
		t.curr.duration = now.Sub(t.curr.start)
		t.items = append(t.items, t.curr)
	}
	t.curr = &timingItem{
		name:  name,
		start: now,
	}
}

func (t *Timing) StartAgg(name string) {
	now := time.Now()
	if t.currAgg != nil {
		t.currAgg.duration = now.Sub(t.currAgg.start)
		t.aggItems[t.currAgg.name] += t.currAgg.duration
	}
	t.currAgg = &timingItem{
		name:  name,
		start: now,
	}
}

func (t *Timing) StopCurrAgg() {
	if t.currAgg != nil {
		now := time.Now()
		t.currAgg.duration = now.Sub(t.currAgg.start)
		t.aggItems[t.currAgg.name] += t.currAgg.duration
		t.currAgg = nil
	}
}

func (t *Timing) Write(w http.ResponseWriter) {
	now := time.Now()
	if t.curr != nil {
		t.curr.duration = now.Sub(t.curr.start)
		t.items = append(t.items, t.curr)
		t.curr = nil
	}
	var sl []string
	for _, item := range t.items {
		sl = append(sl, fmt.Sprintf("%s;dur=%d", item.name, item.duration.Milliseconds()))
	}
	t.StopCurrAgg()
	for name, dur := range t.aggItems {
		sl = append(sl, fmt.Sprintf("%s;dur=%d", name, dur.Milliseconds()))
	}
	w.Header().Add("Server-Timing", strings.Join(sl, ", "))
}
