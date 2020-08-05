package category

import "testing"

func TestMediumRouter(t *testing.T) {
	categoryRouter := Router()
	got := len(categoryRouter.Routes()[0].Handlers)
	expected := 2
	if got != expected {
		t.Errorf("handler returned wrong pattern: got %v want %v",
			got, expected)
	}
}
