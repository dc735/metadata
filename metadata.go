package metadata

import (
	"fmt"
	"strings"
	"time"
)

const DateTimeFormat = "2006-01-02T15:04:05Z"

func ParseTime(text string) (time.Time, error) {
	return time.Parse(DateTimeFormat, string(text))
}

func MustParseTime(text string) time.Time {
	t, _ := ParseTime(text)
	return t
}

func DateTime(t time.Time) string {
	return t.Format(DateTimeFormat)
}

func DateTimePtr(t *time.Time) string {
	return t.Format(DateTimeFormat)
}

type Keys []string

func (k Keys) Len() int           { return len(k) }
func (k Keys) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k Keys) Less(i, j int) bool { return k[i] < k[j] }
func (k Keys) get(i int) string   { return k[i] }

func Lines(notes string) []string {
	return strings.Split(strings.TrimSpace(strings.Replace(notes, "\r\n", "\n", -1)), "\n")
}

func LatLon(latlon *float32) string {
	return fmt.Sprintf("%.4f", *latlon)
}

// simple debugging helper function
func SimpleDiff(s1, s2 string) string {

	l1 := strings.Split(
		strings.TrimSpace(strings.Replace(s1, "\t", "  ", -1)), "\n",
	)
	l2 := strings.Split(
		strings.TrimSpace(strings.Replace(s2, "\t", "  ", -1)), "\n",
	)

	var n, w1, w2 int
	for i := 0; i < len(l1) && i < len(l2); i++ {
		if l1[i] == l2[i] {
			continue
		}
		if len(l1[i]) > w1 {
			w1 = len(l1[i])
		}
		if len(l2[i]) > w2 {
			w2 = len(l2[i])
		}
		if l := len(fmt.Sprintf("%d", i+1)); l > n {
			n = l
		}
	}
	for i := len(l2); i < len(l1); i++ {
		if len(l1[i]) > w1 {
			w1 = len(l1[i])
		}
		if l := len(fmt.Sprintf("%d", i+1)); l > n {
			n = l
		}
	}
	for i := len(l1); i < len(l2); i++ {
		if len(l2[i]) > w2 {
			w2 = len(l2[i])
		}
		if l := len(fmt.Sprintf("%d", i+1)); l > n {
			n = l
		}
	}

	var s []string
	for i := 0; i < len(l1) && i < len(l2); i++ {
		if l1[i] == l2[i] {
			continue
		}
		s = append(s, fmt.Sprintf(fmt.Sprintf("\t[%%%dd]!!! %%-%ds ! %%-%ds !!!", n, w1, w2), i+1, l1[i], l2[i]))
	}
	for i := len(l2); i < len(l1); i++ {
		s = append(s, fmt.Sprintf(fmt.Sprintf("\t[%%%dd]+++ %%-%ds + %%-%ds +++", n, w1, w2), i+1, l1[i], ""))
	}
	for i := len(l1); i < len(l2); i++ {
		s = append(s, fmt.Sprintf(fmt.Sprintf("\t[%%%dd]--- %%-%ds - %%-%ds ---", n, w1, w2), i+1, "", l2[i]))
	}

	return strings.Join(s, "\n")
}
