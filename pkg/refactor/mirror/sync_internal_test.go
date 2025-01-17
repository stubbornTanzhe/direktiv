package mirror

import (
	"reflect"
	"testing"
)

func Test_splitPathToDirectories(t *testing.T) {
	tests := []struct {
		name string
		dir  string
		want []string
	}{
		{name: "case", dir: "/a/b/c/d", want: []string{"/a", "/a/b", "/a/b/c", "/a/b/c/d"}},
		{name: "case", dir: "/a", want: []string{"/a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitPathToDirectories(tt.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitPathToDirectories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDirektivVars(t *testing.T) {
	tests := []struct {
		name  string
		paths []string
		want  [][]string
		want1 [][]string
	}{
		{
			name: "valid_case",
			paths: []string{
				"/var.delta/data.txt",
				"/var.ALPHA.json",
				"/a/v/var.something",
				"/a/v/var.something2",
				"/a/c/workflow",
				"/a/c/workflow.yaml.x1",
				"/banana/page-1.yaml",
				"/banana/page-1.yaml.page.html",
			},
			want: [][]string{
				{"/var", "ALPHA.json"},
				{"/a/v/var", "something"},
				{"/a/v/var", "something2"},
			},
			want1: [][]string{
				{"/a/c/workflow.yaml", "x1"},
				{"/banana/page-1.yaml", "page.html"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseDirektivVars(tt.paths)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDirektivVars() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ParseDirektivVars() got1 = %v, want1 %v", got1, tt.want1)
			}
		})
	}
}
