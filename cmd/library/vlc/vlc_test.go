package vlc

import "testing"

func Test_prepareText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "basic test",
			str:  "This text will be changed, Dear",
			want: "!this text will be changed, !dear",
		},
		{
			name: "special symbols",
			str:  "?!@#$%^&*()_+./,",
			want: "?!@#$%^&*()_+./,",
		},
		{
			name: "upper case",
			str:  "HELLO",
			want: "!h!e!l!l!o",
		},
		{
			name: "numbers",
			str:  "1234567890",
			want: "1234567890",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}
