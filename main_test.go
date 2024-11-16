package main

import (
	"testing"
)

func Test差分を表す文字が行頭に移動されること(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"追加マーカーが行頭に移動されること": {
			input: "    + foo",
			want:  "+     foo",
		},
	}

	for summary, tt := range tests {
		t.Run(summary, func(t *testing.T) {
			got := unindentDiffMarker(tt.input)
			if got != tt.want {
				t.Errorf("got: %q, want: %q", got, tt.want)
			}
		})
	}
}
