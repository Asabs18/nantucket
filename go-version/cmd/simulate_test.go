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
		want  int     // POP=
		want1 float64 // VPROB=
		want2 float64 // TPROB=
		want3 float64 // DPROB=
		want4 int     // DAYS=
		want5 string  // FILE=
		want6 bool    // MULTI=
	}{
		{
			name: "Test parseCmdLine #1",
			args: args{
				args: []string{"pop=1500", "vprob=60", "days=72", "multi=false"},
			},
			want:  1500,
			want1: 60,
			want2: 0.0,
			want4: 72,
			want6: false,
		},
		{
			name: "Test parseCmdLine #2",
			args: args{
				args: []string{"vprob=6", "dprob=4", "days=700", "file=filename.txt", "multi=true"},
			},
			want1: 6,
			want3: 4,
			want4: 700,
			want5: "filename.txt",
			want6: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1, _, _, got4, _, got6 := parseCmdLine(tt.args.args)
			// if got != tt.want {
			// 	t.Errorf("parseCmdLine() got = %v, want %v", got, tt.want)
			// }
			if got1 != tt.want1 {
				t.Errorf("parseCmdLine() got1 = %v, want %v", got1, tt.want1)
			}
			// if got2 != tt.want2 {
			// 	t.Errorf("parseCmdLine() got2 = %v, want %v", got2, tt.want2)
			// }
			// if got3 != tt.want3 {
			// 	t.Errorf("parseCmdLine() got3 = %v, want %v", got3, tt.want3)
			// }
			if got4 != tt.want4 {
				t.Errorf("parseCmdLine() got4 = %v, want %v", got4, tt.want4)
			}
			// if got5 != tt.want5 {
			// 	t.Errorf("parseCmdLine() got5 = %v, want %v", got5, tt.want5)
			// }
			if got6 != tt.want6 {
				t.Errorf("parseCmdLine() got6 = %v, want %v", got6, tt.want6)
			}
		})
	}
}
