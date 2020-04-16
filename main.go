package main

import (
	"fmt"

	"github.com/leonelquinteros/gotext"
)

func main() {
	enLocale := gotext.NewLocale("locale", "en_US")
	enLocale.AddDomain("messages")

	arLocale := gotext.NewLocale("locale", "ar_SA")
	arLocale.AddDomain("messages")

	po := new(gotext.Po)
	po.ParseFile("locale/ar_SA/LC_MESSAGES/messages.po")

	fmt.Printf("English With Just a Get: %s\n", enLocale.Get("This is a sample string"))
	fmt.Printf("Arabic With Just a Get: %s\n", arLocale.Get("This is a sample string"))
	fmt.Printf("Arabic Plural To 0: %s\n", arLocale.GetN("This is a sample string", "This is a sample string", 0))
	fmt.Printf("Arabic Using po.Get: %s\n", po.Get("This is a sample string"))

}
