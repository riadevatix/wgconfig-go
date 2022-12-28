package wgconfig

import "testing"

func Test_trimComment(t *testing.T) {
	tests := []struct {
		name     string
		comment  string
		expected string
	}{
		{"Empty", "", ""},
		{"Clean", "Comment", "Comment"},
		{"With Semicolon", "; A weird one", "A weird one"},
		{"With Hash", "#   I like turtles", "I like turtles"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimComment(tt.comment); got != tt.expected {
				t.Errorf("trimComment() = \"%s\", expected \"%s\"", got, tt.expected)
			}
		})
	}
}
