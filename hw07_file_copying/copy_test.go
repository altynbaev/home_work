package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/udhos/equalfile"
)

func TestCopy(t *testing.T) {
	type args struct {
		fromPath string
		toPath   string
		offset   int64
		limit    int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		refFile string
	}{
		{
			name: "offset0_limit0",
			args: args{
				fromPath: "testdata/input.txt",
				toPath:   "/tmp/test_offset0_limit0.txt",
				offset:   0,
				limit:    0,
			},
			wantErr: false,
			refFile: "testdata/out_offset0_limit0.txt",
		},
		{
			name: "offset0_limit10",
			args: args{
				fromPath: "testdata/input.txt",
				toPath:   "/tmp/test_offset0_limit10.txt",
				offset:   0,
				limit:    10,
			},
			wantErr: false,
			refFile: "testdata/out_offset0_limit10.txt",
		},
		{
			name: "offset0_limit1000",
			args: args{
				fromPath: "testdata/input.txt",
				toPath:   "/tmp/test_offset0_limit1000.txt",
				offset:   0,
				limit:    1000,
			},
			wantErr: false,
			refFile: "testdata/out_offset0_limit1000.txt",
		},
		{
			name: "offset100_limit1000",
			args: args{
				fromPath: "testdata/input.txt",
				toPath:   "/tmp/test_offset100_limit1000.txt",
				offset:   100,
				limit:    1000,
			},
			wantErr: false,
			refFile: "testdata/out_offset100_limit1000.txt",
		},
		{
			name: "offset0_limit10000",
			args: args{
				fromPath: "testdata/input.txt",
				toPath:   "/tmp/test_offset0_limit10000.txt",
				offset:   0,
				limit:    10000,
			},
			wantErr: false,
			refFile: "testdata/out_offset0_limit10000.txt",
		},
		{
			name: "offset6000_limit1000",
			args: args{
				fromPath: "testdata/input.txt",
				toPath:   "/tmp/test_offset6000_limit1000.txt",
				offset:   6000,
				limit:    1000,
			},
			wantErr: false,
			refFile: "testdata/out_offset6000_limit1000.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copy(tt.args.fromPath, tt.args.toPath, tt.args.offset, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
			cmp := equalfile.New(nil, equalfile.Options{}) // compare using single mode
			equal, err := cmp.CompareFile(tt.args.toPath, tt.refFile)
			require.NoError(t, err)
			assert.True(t, equal)
		})
	}
}
