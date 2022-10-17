/*
Copyright © 2022 Asabs18 & davegi
This file is part of the `nantucket` project.
*/

package cmd

import "testing"

func Test_parseCmdLine(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 float64
		want2 float64
		want3 float64
		want4 int
		want5 string
		want6 bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, got4, got5, got6 := parseCmdLine(tt.args.args)
			if got != tt.want {
				t.Errorf("parseCmdLine() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseCmdLine() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("parseCmdLine() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("parseCmdLine() got3 = %v, want %v", got3, tt.want3)
			}
			if got4 != tt.want4 {
				t.Errorf("parseCmdLine() got4 = %v, want %v", got4, tt.want4)
			}
			if got5 != tt.want5 {
				t.Errorf("parseCmdLine() got5 = %v, want %v", got5, tt.want5)
			}
			if got6 != tt.want6 {
				t.Errorf("parseCmdLine() got6 = %v, want %v", got6, tt.want6)
			}
		})
	}
}
