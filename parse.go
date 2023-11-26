package link

import "io"

type Link struct {
	Href string
	Text string
}

// parse html doc and return slice of links sturct

func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
