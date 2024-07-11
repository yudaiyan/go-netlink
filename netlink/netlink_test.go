package netlink

import (
	"testing"
)

func TestGetLocalInterface(t *testing.T) {
	type args struct {
		ifname string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test1", args: args{ifname: "eth0"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, _, err := GetLocalInterface(tt.args.ifname)
			if err != nil {
				t.Errorf("GetLocalInterface() error = %v", err)
				return
			}

		})
	}
}

func TestGetMac(t *testing.T) {
	type args struct {
		ifname string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test1", args: args{ifname: "eth0"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetMac(tt.args.ifname)
			if err != nil {
				t.Errorf("GetMac() error = %v", err)
				return
			}
		})
	}
}
