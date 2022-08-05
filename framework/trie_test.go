package framework

import (
	"reflect"
	"testing"
)

func Test_node_matchNode(t *testing.T) {
	type fields struct {
		isLast  bool
		segment string
		handler ControllerHandler
		childs  []*node
	}
	type args struct {
		uri string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
	}{
		{
			"",
			fields{},
			args{uri: "/test1/test2"},
			&node{},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				isLast:  tt.fields.isLast,
				segment: tt.fields.segment,
				handler: tt.fields.handler,
				childs:  tt.fields.childs,
			}
			if got := n.matchNode(tt.args.uri); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matchNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
