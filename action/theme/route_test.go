package theme

import "testing"

func TestMediumRouter(t *testing.T) {
	themeRouter := Router()
	got := len(themeRouter.Routes()[0].Handlers)
	expected := 2
	if got != expected {
		t.Errorf("handler returned wrong pattern: got %v want %v",
			got, expected)
	}
}
