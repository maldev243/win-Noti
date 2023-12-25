package main

import (
	"time"

	"github.com/go-toast/toast"
)

func main() {
	for {
		// Wait for 3 seconds
		time.Sleep(3 * time.Second)

		// Create a new notification
		notification := toast.Notification{
			AppID:   "Example App", // Use an appropriate AppID
			Title:   "Notification Title",
			Message: "Your message here",
		}

		// Show the notification
		err := notification.Push()
		if err != nil {
			panic(err)
		}
	}
}
