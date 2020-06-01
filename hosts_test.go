package hosts

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	type args struct {
		rdr io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []Record
		wantErr bool
	}{
		{
			name: "can find records that are valid records but also double commented out",
			args: args{rdr: strings.NewReader(`# this is a comment
##
127.0.0.1 domain.com example.com
# 192.168.5.5 domain.com example.com`)},
			want: []Record{
				{
					Line:        1,
					IsCommented: true,
					IP:          "",
					Hosts:       nil,
				},
				{
					Line:        2,
					IsCommented: true,
					IP:          "",
					Hosts:       nil,
				},
				{
					Line:        3,
					IsCommented: false,
					IP:          "127.0.0.1",
					Hosts:       []string{"domain.com", "example.com"},
				},
				{
					Line:        4,
					IsCommented: true,
					IP:          "192.168.5.5",
					Hosts:       []string{"domain.com", "example.com"},
				},
			},
		},
		{
			name: "can find records that are valid records but also commented out",
			args: args{rdr: strings.NewReader(`# this is a comment
127.0.0.1 domain.com example.com
# 192.168.5.5 domain.com example.com`)},
			want: []Record{
				{
					Line:        1,
					IsCommented: true,
					IP:          "",
					Hosts:       nil,
				},
				{
					Line:        2,
					IsCommented: false,
					IP:          "127.0.0.1",
					Hosts:       []string{"domain.com", "example.com"},
				},
				{
					Line:        3,
					IsCommented: true,
					IP:          "192.168.5.5",
					Hosts:       []string{"domain.com", "example.com"},
				},
			},
		},
		{
			name: "can find records that are valid records",
			args: args{rdr: strings.NewReader(`# this is a comment
127.0.0.1 domain.com example.com`)},
			want: []Record{
				{
					Line:        1,
					IsCommented: true,
					IP:          "",
					Hosts:       nil,
				},
				{
					Line:        2,
					IsCommented: false,
					IP:          "127.0.0.1",
					Hosts:       []string{"domain.com", "example.com"},
				},
			},
		},
		{
			name: "can find records that have comments",
			args: args{rdr: strings.NewReader(`# this is a comment`)},
			want: []Record{
				{
					Line:        1,
					IsCommented: true,
					IP:          "",
					Hosts:       nil,
				},
			},
		},
		{
			name: "empty lines do not return a record",
			args: args{rdr: strings.NewReader(``)},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(tt.args.rdr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = \n%v, \nwant \n%v", got, tt.want)
			}
		})
	}
}
