package timer

import (
	"testing"
	"time"
)

func TestGetCalculateTime(t *testing.T) {
	now := time.Now()
	addTime := "2h"
	target := now.Add(time.Hour * 2)
	if ans, err := GetCalculateTime(now, addTime); err != nil || ans != target {
		t.Errorf("got %q, want %q", ans, target)
	}
}
