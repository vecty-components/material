package component

import (
	"testing"

	"github.com/gopherjs/gopherjs/js"
)

func TestNew(t *testing.T) {
	type args struct {
		n Type
	}
	type testcase struct {
		name       string
		args       args
		wantStatus string
		wantErr    bool
	}
	tests := make([]testcase, TypeCount, TypeCount)
	for i, ctype := range ComponentTypes() {
		newComponent := &component{nil, ctype, Stopped}
		tests[i] = testcase{
			name:       ctype.String(),
			args:       args{ctype},
			wantStatus: newComponent.String(),
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMdcComponent, err := New(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMdcComponent.String() != tt.wantStatus {
				t.Errorf("New() = %s, want %s", gotMdcComponent, tt.wantStatus)
				return
			}
			mdcO := gotMdcComponent.GetObject()
			if mdcO == nil || mdcO == js.Undefined {
				t.Errorf("New() object is: %v", mdcO)
				return
			}
		})
	}
}
