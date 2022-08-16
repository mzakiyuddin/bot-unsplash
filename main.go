package main

import (
	"time"

	"github.com/go-rod/rod"
)

func main() {
	page := rod.New().MustConnect().MustPage("https://unsplash.com/photos/1HaNBky0jxA")
	page.MustWaitLoad()
	page.MustElementR("span", "Download free").MustClick()
	time.Sleep(5 * time.Second)
	page.MustElementR("span", "Download free").MustClick()
	time.Sleep(5 * time.Second)
	page.MustElementR("span", "Download free").MustClick()
	time.Sleep(5 * time.Second)
	page.MustElementR("span", "Download free").MustClick()
	time.Sleep(5 * time.Second)
	page.MustElementR("span", "Download free").MustClick()
	time.Sleep(5 * time.Second)
	page.MustElementR("span", "Download free").MustClick()
	time.Sleep(5 * time.Second)
	page.MustElementR("span", "Download free").MustClick()
	time.Sleep(5 * time.Second)
}
