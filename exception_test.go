package exception

import (
	"errors"
	"reflect"
	"testing"
)

func TestIsError(t *testing.T) {
	type args struct {
		e interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"nil",args{nil}, false},
		{"error", args{errors.New("test")}, true},
		{"exception_failure", args{exception{code:-1,message: "test",}}, true},
		{"exception_success", args{exception{code:0,message: "success"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsError(tt.args.e); got != tt.want {
				t.Errorf("HasException() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want Exception
	}{
		{"error", args{errors.New("test-error")}, exception{code:-1,message: "test-error"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTryGenerateException(t *testing.T) {
	type args struct {
		mainException Exception
		err           error
	}
	tests := []struct {
		name string
		args args
		want Exception
	}{
		{"error", args{Failure, errors.New("err")}, Failure.SubError(errors.New("err"))},
		{"error", args{Failure, nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateWhenError(tt.args.mainException, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateWhenError() = %v, want %v", got, tt.want)
			}
		})
	}
}