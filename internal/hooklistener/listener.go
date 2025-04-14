package hooklistener

import (
	hook "github.com/robotn/gohook"
)

type EventHandler func(ev hook.Event)

func ListenForKeys(handler EventHandler) {
	evs := hook.Start()
	go func() {
		for ev := range evs {
			if ev.Kind == hook.KeyDown {
				handler(ev)
			}
		}
	}()
}
