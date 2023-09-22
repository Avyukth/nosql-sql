package main

import "testing"

func TestGenerateInsertSQL(t *testing.T) {
	type args struct {
		oplog string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
r _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateInsertSQL(tt.args.oplog)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateInsertSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateInsertSQL() = %v, want %v", got, tt.want)
			}
		})
	}
