package file

import (
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/ninjaneers-team/uropa/opa"
)

func Test_configFilesInDir(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "empty directory",
			args:    args{"testdata/emptydir"},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "directory does not exist",
			args:    args{"testdata/does-not-exist"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid directory",
			args: args{"testdata/emptyfiles"},
			want: []string{
				"testdata/emptyfiles/Baz.YamL",
				"testdata/emptyfiles/bar.yaml",
				"testdata/emptyfiles/foo.yml",
				"testdata/emptyfiles/foobar.json",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := configFilesInDir(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("configFilesInDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configFilesInDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getReaders(t *testing.T) {
	type args struct {
		fileOrDir string
	}
	tests := []struct {
		name string
		args args
		want []io.Reader
		// length of returned array
		wantLen int
		wantErr bool
	}{
		{
			name:    "read from standard input",
			args:    args{"-"},
			want:    []io.Reader{os.Stdin},
			wantLen: 1,
			wantErr: false,
		},
		{
			name:    "directory does not exist",
			args:    args{"testdata/does-not-exist"},
			want:    nil,
			wantLen: 0,
			wantErr: true,
		},
		{
			name:    "valid directory",
			args:    args{"testdata/emptyfiles"},
			want:    nil,
			wantLen: 4,
			wantErr: false,
		},
		{
			name:    "valid file",
			args:    args{"testdata/file.yaml"},
			want:    nil,
			wantLen: 1,
			wantErr: false,
		},
		{
			name:    "valid JSON file",
			args:    args{"testdata/file.json"},
			want:    nil,
			wantLen: 1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getReaders(tt.args.fileOrDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("getReaders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantLen != len(got) {
				t.Errorf("getReaders() mismatch in returned length: "+
					"want = %v, got = %v", tt.wantLen, len(got))
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getReaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getContent(t *testing.T) {
	type args struct {
		fileOrDir string
	}
	tests := []struct {
		name    string
		args    args
		want    *Content
		wantErr bool
	}{
		{
			name:    "directory does not exist",
			args:    args{"testdata/does-not-exist"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty directory",
			args:    args{"testdata/emptydir"},
			want:    &Content{},
			wantErr: false,
		},
		{
			name:    "directory with empty files",
			args:    args{"testdata/emptyfiles"},
			want:    &Content{},
			wantErr: false,
		},
		{
			name:    "bad yaml",
			args:    args{"testdata/badyaml"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "bad JSON",
			args:    args{"testdata/badjson"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "single file",
			args: args{"testdata/file.yaml"},
			want: &Content{
				Services: []FPolicy{
					{
						Service: Opa.Service{
							Name: Opa.String("svc2"),
							Host: Opa.String("2.example.com"),
						},
						Routes: []*FRoute{
							{
								Route: Opa.Route{
									Name:  Opa.String("r2"),
									Paths: Opa.StringSlice("/r2"),
								},
							},
						},
					},
				},
				Plugins: []FPlugin{
					{
						Plugin: Opa.Plugin{
							Name: Opa.String("prometheus"),
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "valid directory",
			args: args{"testdata/valid"},
			want: &Content{
				Info: &Info{
					SelectorTags: []string{"tag1"},
				},
				Services: []FPolicy{
					{
						Service: Opa.Service{
							Name: Opa.String("svc2"),
							Host: Opa.String("2.example.com"),
						},
						Routes: []*FRoute{
							{
								Route: Opa.Route{
									Name:  Opa.String("r2"),
									Paths: Opa.StringSlice("/r2"),
								},
							},
						},
					},
					{
						Service: Opa.Service{
							Name: Opa.String("svc1"),
							Host: Opa.String("1.example.com"),
							Tags: Opa.StringSlice("team-svc1"),
						},
						Routes: []*FRoute{
							{
								Route: Opa.Route{
									Name:  Opa.String("r1"),
									Paths: Opa.StringSlice("/r1"),
								},
							},
						},
					},
				},
				Consumers: []FConsumer{
					{
						Consumer: Opa.Consumer{
							Username: Opa.String("foo"),
						},
					},
					{
						Consumer: Opa.Consumer{
							Username: Opa.String("bar"),
						},
					},
					{
						Consumer: Opa.Consumer{
							Username: Opa.String("harry"),
						},
					},
				},
				Plugins: []FPlugin{
					{
						Plugin: Opa.Plugin{
							Name: Opa.String("prometheus"),
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getContent(tt.args.fileOrDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("getContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
