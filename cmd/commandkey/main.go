package main

import (
	"fmt"
	"os"
	"time"

	"github.com/justbarnt/commandkey/internal/actions"
	"github.com/justbarnt/commandkey/internal/hooklistener"
	hook "github.com/robotn/gohook"
)

var (
	leaderActive   bool
	lastLeaderTime time.Time
)

func main() {
	fmt.Println("Listening for global key events (Ctrl+Q to quit)...")

	hooklistener.ListenForKeys(func(ev hook.Event) {
		fmt.Printf("Key pressed: %v (Rawcode: %d)\n", ev.Keychar, ev.Rawcode)

		if ev.Rawcode == 32 && ev.Mask&hooklistener.ModCtrl != 0 {
			leaderActive = true
			lastLeaderTime = time.Now()
			fmt.Println("Leader key activated")
		} else if leaderActive && time.Since(lastLeaderTime) < 2*time.Second {
			if ev.Keychar == 'h' || ev.Keychar == 'H' {
				fmt.Println("Action key 'H' detected, launching app...")
				actions.LaunchEditor()
				leaderActive = false
			}
		} else if ev.Keychar == 'q' && ev.Mask&hooklistener.ModCtrl != 0 {
			fmt.Println("Exiting...")
			hook.End()
			os.Exit(0)
		}
	})

	select {}
}
