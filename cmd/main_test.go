package main

import (
	"testing"
)

type args struct {
	opLog string
}

type testCase struct {
	name    string
	args    args
	want    string
	wantErr bool
}

func runTests(t *testing.T, tests []testCase, testFunc func(string) (string, error)) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testFunc(tt.args.opLog)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateSQL(t *testing.T) {
	t.Parallel()
	t.Run("Insert", insertStatement)

	t.Run("Update", updateStatement)

	t.Run("Delete", deleteStatement)
}

func insertStatement(t *testing.T) {
	tests := []testCase{
		{
			name: "Insert Operations",
			args: args{
				opLog: `{
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
	runTests(t, tests, GenerateInsertSQL)
}

func updateStatement(t *testing.T) {

	tests := []testCase{
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
	runTests(t, tests, GenerateUpdateSQL)
}

func deleteStatement(t *testing.T) {
	tests := []testCase{
		{
			name: "Test Delete Operation",
			args: args{
				opLog: `{
            "op": "d",
            "ns": "test.student",
            "o": {
                "_id": "635b79e231d82a8ab1de863b"
            }
        }`,
			},
			want:    "DELETE FROM student WHERE _id = '635b79e231d82a8ab1de863b';",
			wantErr: false,
		},
	}
	runTests(t, tests, GenerateDeleteSQL)

}
