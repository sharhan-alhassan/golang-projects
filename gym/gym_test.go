

package gym

import "testing"

func TestGym(t *testing.T)  {
	want := "Welcome, Gym members"
	if got := Gym(); got != want {
		t.Errorf("Gym() = %q, want %q", got, want)
	}
}