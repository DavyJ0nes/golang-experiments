package cnfg

import (
	"os"
	"reflect"
	"testing"
)

func Test_readConfigFile(t *testing.T) {
	testFileOK, err := os.Open("../config.yml")
	if err != nil {
		t.Fatalf("Unknown Error: %v", err.Error())
	}

	type args struct {
		file *os.File
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []configPair
		wantErr    bool
	}{
		{
			"Simple",
			args{file: testFileOK},
			[]configPair{
				{"baseFilePath", "/tmp"},
				{"machineName", "freddie"},
				{"daysToKeep", "30"},
				{"cpuMultiplier", "100"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := readConfigFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("got error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("got = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestGet(t *testing.T) {
	// Initialise
	MustLoad("../config.yml")

	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Empty Input",
			args{""},
			"",
		},
		{
			"Unknown Key",
			args{"Unknown"},
			"",
		},
		{
			"Known Key",
			args{"machineName"},
			"freddie",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.key); got != tt.want {
				t.Errorf("got = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}
