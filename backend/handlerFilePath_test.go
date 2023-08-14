package backend

import "testing"

func TestFilePathHandler_OpenDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		fp   *FilePathHandler
		args args
	}{
		{
			name: "test-1",
			fp:   &FilePathHandler{},
			args: args{
				path: "../",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fp.OpenDir(tt.args.path)
		})
	}
}
