package main

import (
	"testing"
)

// Test_Diff_marker_is_moved_to_the_beginning_of_the_line
func Test差分を表す文字が行頭に移動されること(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"追加マーカーが行頭に移動されること": {
			input: "    + foo",
			want:  "+     foo",
		},
		"削除マーカーが行頭に移動されること": {
			input: "    - foo",
			want:  "-     foo",
		},
		"変更マーカーが行頭に移動されること": {
			input: "    ~ foo",
			want:  "~     foo",
		},
		"ハードタブでインデントしている場合でも差分マーカーが移動されること": {
			input: "\t+ foo",
			want:  "+\t foo",
		},
		"空行は変更されないこと": {
			input: "",
			want:  "",
		},
		"差分マーカーで開始しない行は変更されないこと": {
			input: "    foo",
			want:  "    foo",
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
