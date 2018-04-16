package utility

import (
	"testing"
	"time"
)

func TestIsTime(t *testing.T) {
	//t.Parallel()

	pDate := time.Date(2018, time.March, 20, 9, 5, 20, 80, time.Local)
	eDate := time.Date(2018, time.March, 10, 12, 5, 20, 30, time.Local)
	tDate := time.Now()

	t.Run("pDate", func(t *testing.T) {
		if IsTime(pDate, tDate) {
			t.Errorf("passing date failed... There is a bug.")
		}
	})

	t.Run("eDate", func(t *testing.T) {
		if IsTime(eDate, tDate) {
			t.Errorf("Ege check 13 days.  Failed error!")
		}
	})
}
