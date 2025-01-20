package name_test

import (
	"fmt"
	"github.com/Jumpaku/schenerate/name"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_registry_Safe(t *testing.T) {
	tests := []struct {
		reserved []string
		modify   func(name string, times int) string
		in       string
		want     string
	}{
		{
			reserved: []string{"a", "b"},
			in:       "x",
			want:     "x",
		},
		{
			reserved: []string{"a", "b"},
			modify: func(name string, times int) string {
				return fmt.Sprintf(`%s_%d`, name, times+1)
			},
			in:   "a",
			want: "a_2",
		},
		{
			reserved: []string{"a", "a_2"},
			modify: func(name string, times int) string {
				return fmt.Sprintf(`%s_%d`, name, times+1)
			},
			in:   "a",
			want: "a_3",
		},
		{
			reserved: []string{"a", "a_2"},
			modify: func(name string, times int) string {
				for i := 0; i < times; i++ {
					name = fmt.Sprintf(`%s_2`, name)
				}
				return name
			},
			in:   "a",
			want: "a_2_2",
		},
	}

	for number, tt := range tests {
		t.Run(fmt.Sprintf(`%02d-%s`, number, tt.want), func(t *testing.T) {
			sut := name.NewRegistry(tt.reserved, tt.modify)
			got := sut.Safe(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}
