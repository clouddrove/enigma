package helm

import (
	"testing"
)

func TestInstallHelm(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("InstallHelm() panicked: %v", r)
		}
	}()

	InstallHelm()
}

func TestIsValidEnvVarKey(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want bool
	}{
		{
			name: "empty string",
			key:  "",
			want: false,
		},
		{
			name: "valid key with letters only",
			key:  "PATH",
			want: true,
		},
		{
			name: "valid key with underscore",
			key:  "HELM_PATH",
			want: true,
		},
		{
			name: "valid key with numbers",
			key:  "PATH2",
			want: true,
		},
		{
			name: "invalid key starting with number",
			key:  "2PATH",
			want: false,
		},
		{
			name: "invalid key with special characters",
			key:  "PATH@HOME",
			want: false,
		},
		{
			name: "valid key starting with underscore",
			key:  "_PATH",
			want: true,
		},
		{
			name: "complex valid key",
			key:  "HELM_CHART_PATH_V2",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidEnvVarKey(tt.key); got != tt.want {
				t.Errorf("isValidEnvVarKey(%q) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}
