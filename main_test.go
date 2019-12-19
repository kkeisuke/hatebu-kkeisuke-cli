package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "正常系",
			args: []string{"htb", "-freeword", "golang", "-perPage", "1"},
			want: 0,
		},
		{
			name: "正常系（短縮）",
			args: []string{"htb", "-f", "golang", "-p", "1"},
			want: 0,
		},
		{
			name: "フリーワードなし",
			args: []string{"htb", "-f", "", "-p", "1"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args); got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
