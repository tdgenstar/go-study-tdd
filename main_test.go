package main

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		name     string
		language string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "it should print hello, world",
			args: args{
				name:     "Teddy",
				language: "",
			},
			want: "Hello, Teddy",
		},
		{
			name: "it should empty name, then print 'World'",
			args: args{
				name:     "",
				language: "",
			},
			want: "Hello, World",
		},
		{
			name: "it should language == spanish and empty name, then Hola, World",
			args: args{
				name:     "",
				language: spanish,
			},
			want: "Hola, World",
		},
		{
			name: "이름은 teddy, 언어는 spanish 라면 'Hola, teddy' 출력",
			args: args{
				name:     "teddy",
				language: spanish,
			},
			want: "Hola, teddy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name, tt.args.language); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
