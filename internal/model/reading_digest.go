package model

import "fmt"

type ReadingDigest struct {
	ID      string            `json:"id"`
	Title   string            `json:"title"`
	Sections map[string]string `json:"sections"`
	Ready   bool              `json:"ready"`
}

func PopulateReadingDigest(digest *ReadingDigest, title string) {
	digest.Title = title
	digest.Sections = map[string]string{"overview": "pending"}
	if title == "broken" {
		panic("summary provider returned malformed content")
	}
	digest.Sections["overview"] = "ready"
	digest.Ready = true
}

func (d *ReadingDigest) Summary() string {
	if !d.Ready {
		panic(fmt.Sprintf("digest %s is not ready", d.ID))
	}
	return d.Sections["overview"]
}
