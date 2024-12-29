package mac

import (
	"errors"
	"testing"
)

func TestParseCompactMAC(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    MAC
		wantErr error
	}{
		{
			name:    "valid",
			s:       "e675a31ea872",
			want:    MAC{0x72, 0xa8, 0x1e, 0xa3, 0x75, 0xe6},
			wantErr: nil,
		},
		{
			name:    "invalid",
			s:       "e675a31ea87",
			want:    MAC{},
			wantErr: ErrInvalidFormat,
		},
		{
			name:    "empty",
			s:       "",
			want:    MAC{},
			wantErr: ErrInvalidFormat,
		},
		{
			name:    "short",
			s:       "e675a31ea87",
			want:    MAC{},
			wantErr: ErrInvalidFormat,
		},
		{
			name:    "long",
			s:       "e675a31ea8721",
			want:    MAC{},
			wantErr: ErrInvalidFormat,
		},
		{
			name:    "colon",
			s:       "E6:75:A3:1E:A8:72",
			want:    MAC{},
			wantErr: ErrInvalidFormat,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCompactMAC(tt.s)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Parse() got error %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Parse() got mac = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseColonMAC(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    MAC
		wantErr error
	}{
		{
			name:    "valid",
			s:       "E6:75:A3:1E:A8:72",
			want:    MAC{0x72, 0xa8, 0x1e, 0xa3, 0x75, 0xe6},
			wantErr: nil,
		},
		{
			name:    "invalid",
			s:       "E6:75:A3:1E:A8:7",
			want:    MAC{},
			wantErr: ErrInvalidFormat,
		},
		{
			name:    "empty",
			s:       "",
			want:    MAC{},
			wantErr: ErrInvalidFormat,
		},
		{
			name:    "short",
			s:       "E6:75:A3:1E:A8:7",
			want:    MAC{},
			wantErr: ErrInvalidFormat,
		},
		{
			name:    "long",
			s:       "E6:75:A3:1E:A8:72:1",
			want:    MAC{},
			wantErr: ErrInvalidFormat,
		},
		{
			name:    "compact",
			s:       "e675a31ea872",
			want:    MAC{},
			wantErr: ErrInvalidFormat,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseColonMAC(tt.s)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Parse() got error %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Parse() got mac = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMac_String(t *testing.T) {
	tests := []struct {
		name string
		m    MAC
		want string
	}{
		{
			name: "valid",
			m:    MAC{0x72, 0xa8, 0x1e, 0xa3, 0x75, 0xe6},
			want: "E6:75:A3:1E:A8:72",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("MAC.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
