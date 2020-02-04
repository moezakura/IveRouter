package main

import "fmt"

var Util = util{}

const (
	B2KB = 1024
	KB   = B2KB
	MB   = B2KB * B2KB
	GB   = B2KB * B2KB * B2KB
)

type util struct {
}

func (util) toDataCast(length float64) string {
	s := ""
	if length > KB {
		if length > MB {
			if length > GB {
				s = fmt.Sprintf("%.4f gb", length/GB)
			} else {
				s = fmt.Sprintf("%.2f mb", length/MB)
			}
		} else {
			s = fmt.Sprintf("%.0f kb", length/KB)
		}
	} else {
		s = fmt.Sprintf("%.0f b", length)
	}
	return s
}
