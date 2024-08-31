package env

import "testing"

func TestEnv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试读取配置文件",
			args: args{
				key: "PORT",
			},
			want: ":80",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Env(tt.args.key); got != tt.want {
				t.Errorf("Env() = %v, want %v", got, tt.want)
			}
		})
	}
}
