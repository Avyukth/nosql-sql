package main

import (
	"testing"
)

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
		{
			name: "Insert Operations",
			args: args{
				oplog: `{
                    "op" : "i",
                    "ns" : "test.student",
                    "o" : {
                        "_id" : "635b79e231d82a8ab1de863b",
                        "name" : "Selena Miller",
                        "roll_no" : 51,
                        "is_graduated" : false,
                        "date_of_birth" : "2000-01-30"
                    }
                }`,
			},
			want:    "INSERT INTO student (_id, date_of_birth, is_graduated, name, roll_no) VALUES ('635b79e231d82a8ab1de863b', '2000-01-30', false, 'Selena Miller', 51);",
			wantErr: false,
		},
	}

	for _, tt := range tests {
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
}

// func TestGenerateSQL(t *testing.T) {
// 	type args struct {
// 		opLog string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    string
// 		wantErr bool
// 	}{
// 		{
// 			name: "Test Invalid Operation",
// 			args: args{
// 				opLog: `{
// 					"op": "x",
// 					"ns": "test.student",
// 					"o": {
// 						"_id": "635b79e231d82a8ab1de863b"
// 					}
// 				}`,
// 			},
// 			want:    "",
// 			wantErr: true, // Expecting an error as the operation is invalid
// 		},
// 		{
// 			name: "Test Missing Namespace",
// 			args: args{
// 				opLog: `{
// 					"op": "i",
// 					"o": {
// 						"_id": "635b79e231d82a8ab1de863b"
// 					}
// 				}`,
// 			},
// 			want:    "",
// 			wantErr: true, // Expecting an error as the namespace is missing
// 		},
// 		{
// 			name: "Test Missing o Field",
// 			args: args{
// 				opLog: `{
// 					"op": "i",
// 					"ns": "test.student"
// 				}`,
// 			},
// 			want:    "",
// 			wantErr: true, // Expecting an error as the o field is missing
// 		},
// 		// TODO: Add more integration test cases as needed.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := GenerateSQL(tt.args.opLog)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GenerateSQL() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("GenerateSQL() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestGenerateUpdateSQL(t *testing.T) {
	type args struct {
		opLog string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Update is_graduated",
			args: args{
				opLog: `{
					"op": "u",
					"ns": "test.student",
					"o": {
						"$v": 2,
						"diff": {
							"u": {
								"is_graduated": true
							}
						}
					},
					"o2": {
						"_id": "635b79e231d82a8ab1de863b"
					}
				}`,
			},
			want:    "UPDATE student SET is_graduated = true WHERE _id = '635b79e231d82a8ab1de863b';",
			wantErr: false,
		},
		{
			name: "Test Missing _id",
			args: args{
				opLog: `{
					"op": "u",
					"ns": "test.student",
					"o": {
						"$v": 2,
						"diff": {
							"u": {
								"is_graduated": true
							}
						}
					},
					"o2": {}
				}`,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test Delete roll_no",
			args: args{
				opLog: `{
            "op": "u",
            "ns": "test.student",
            "o": {
                "$v": 2,
                "diff": {
                    "d": {
                        "roll_no": false
                    }
                }
            },
            "o2": {
                "_id": "635b79e231d82a8ab1de863b"
            }
        }`,
			},
			want:    "UPDATE student SET roll_no = NULL WHERE _id = '635b79e231d82a8ab1de863b';",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateUpdateSQL(tt.args.opLog)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateUpdateSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateUpdateSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
