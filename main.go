package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	page := rod.New().MustConnect().MustPage("https://unsplash.com/photos/1HaNBky0jxA")
	page.MustWaitLoad()

	for i := 0; i < 10; i++ {
		page.MustElementR("span", "Download free").MustClick()
		fmt.Println("Success download")

		if i == 0 {
			time.Sleep(2 * time.Second)
			fmt.Println("Try to click got it")
			page.MustElementR("button", "Got it").MustClick()
		}
		time.Sleep(5 * time.Second)
	}

	page.Browser().MustPage("https://unsplash.com/photos/osVeNwhlBes")
	page.MustWaitLoad()

	for i := 0; i < 10; i++ {
		page.MustElementR("span", "Download free").MustClick()
		fmt.Println("Success download")
		time.Sleep(5 * time.Second)
	}
}
