package flect

import "testing"

func TestGetType(t *testing.T) {
	type Object struct{}
	type Variable string
	type unexported string

	type args struct {
		myvar interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Object",
			args: args{myvar: Object{}},
			want: "Object",
		},
		{
			name: "Variable",
			args: args{myvar: Variable("")},
			want: "Variable",
		},
		{
			name: "unexported",
			args: args{myvar: unexported("")},
			want: "unexported",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetType(tt.args.myvar); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}
