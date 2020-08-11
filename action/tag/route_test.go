package tag

import "testing"

func TestTagRouter(t *testing.T) {
	tagRouter := Router()
	got := len(tagRouter.Routes()[0].Handlers)
	expected := 2
	if got != expected {
		t.Errorf("handler returned wrong pattern: got %v want %v",
			got, expected)
	}
}
