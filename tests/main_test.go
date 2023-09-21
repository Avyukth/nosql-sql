package tests

func TestGenerateInsertSQL(t *testing.T) {
	tests := []struct {
		name    string
		oplog   string
		want    string
		wantErr bool
	}{
		{
			name:    "Insert Operation",
			oplog:   "",
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateInsertSQL(); got != tt.want {
				t.Errorf("GenerateInsertSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
