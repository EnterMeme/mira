package main

import (
	"github.com/EnterMeme/mira"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))

	r.Redditor("myuser").Compose("mytitle", "mytext")
}
