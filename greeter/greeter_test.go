package greeter

import "testing"

func TestGreet(t *testing.T) {
	type args struct {
		name string
		hour int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test good morning and valid name",
			args: args{name: "Andrew", hour: 7},
			want: "Good morning, Andrew!",
		},
		{
			name: "test good morning and invalid name",
			args: args{name: "andrew", hour: 7},
			want: "Good morning, Andrew!",
		},
		{
			name: "test hello and valid name",
			args: args{name: "Andrew", hour: 13},
			want: "Hello, Andrew!",
		},
		{
			name: "test hello and invalid name",
			args: args{name: "   andrew   ", hour: 13},
			want: "Hello, Andrew!",
		},
		{
			name: "test good evening and valid name",
			args: args{name: "Andrew", hour: 19},
			want: "Good evening, Andrew!",
		},
		{
			name: "test good evening and invalid name",
			args: args{name: "   andrew   ", hour: 19},
			want: "Good evening, Andrew!",
		},
		{
			name: "test good night and valid name",
			args: args{name: "Andrew", hour: 23},
			want: "Good night, Andrew!",
		},
		{
			name: "test good night and invalid name",
			args: args{name: "   andrew   ", hour: 23},
			want: "Good night, Andrew!",
		},
		{
			name: "test invalid hour and valid name",
			args: args{name: "Andrew", hour: 25},
			want: "Hour out of range",
		},
		{
			name: "test hour is negative integer and valid name",
			args: args{name: "Andrew", hour: -1},
			want: "Hour out of range",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.args.name, tt.args.hour); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
