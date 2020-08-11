package chart

import "testing"

func TestChartRouter(t *testing.T) {
	chartRouter := Router()
	got := len(chartRouter.Routes()[0].Handlers)
	expected := 2
	if got != expected {
		t.Errorf("handler returned wrong pattern: got %v want %v",
			got, expected)
	}
}
