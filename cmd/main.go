package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func main() {

}

type ColumnValue struct {
	Column string
	Value  string
}
type OpLogEntry struct {
	Op string                 `json:"op"`
	Ns string                 `json:"ns"`
	O  map[string]interface{} `json:"o"`
	O2 map[string]interface{} `json:"o2"`
}

func GenerateSQL(opLog string) (string, error) {
	var opLogObject OpLogEntry
	err := json.Unmarshal([]byte(opLog), &opLogObject)
	if err != nil {
		return "", err
	}

	switch opLogObject.Op {
	case "i":
		return GenerateInsertSQL(opLog)
	case "u":
		return GenerateUpdateSQL(opLog)
	default:
		return "", nil
	}
}

func GenerateInsertSQL(opLog string) (string, error) {
	var opLogObject OpLogEntry
	err := json.Unmarshal([]byte(opLog), &opLogObject)
	if err != nil {
		return "", err
	}
	switch opLogObject.Op {
	case "i":
		columns, values, err := getColumnAndValue(opLogObject.O)
		if err != nil {
			return "", err
		}
		// Extracting table name from Ns
		tableName := strings.Split(opLogObject.Ns, ".")[1]

		sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, strings.Join(columns, ", "), strings.Join(values, ", "))
		return sql, nil

	default:
		return "", nil
	}

}

func getColumnAndValue(operation map[string]interface{}) ([]string, []string, error) {
	var columnValues []ColumnValue
	for key, value := range operation {
		var formattedValue string
		switch v := value.(type) {
		case string:
			formattedValue = fmt.Sprintf("'%v'", v)
		default:
			formattedValue = fmt.Sprintf("%v", v)
		}
		columnValues = append(columnValues, ColumnValue{Column: key, Value: formattedValue})
	}

	// Sort columnValues slice based on the Column field
	sort.Slice(columnValues, func(i, j int) bool {
		return columnValues[i].Column < columnValues[j].Column
	})

	// Extract sorted columns and values
	columns := make([]string, len(columnValues))
	values := make([]string, len(columnValues))
	for i, cv := range columnValues {
		columns[i] = cv.Column
		values[i] = cv.Value
	}

	return columns, values, nil
}

func GenerateUpdateSQL(opLog string) (string, error) {
	var opLogObject OpLogEntry
	err := json.Unmarshal([]byte(opLog), &opLogObject)
	if err != nil {
		return "", err
	}

	// Extracting table name from Ns
	tableName := strings.Split(opLogObject.Ns, ".")[1]

	// Extracting the _id from o2 field
	id, ok := opLogObject.O2["_id"]
	if !ok {
		return "", fmt.Errorf("missing _id in o2 field")
	}

	// Extracting the diff field from o field
	diff, ok := opLogObject.O["diff"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("missing or invalid diff in o field")
	}

	// Constructing the SQL update statement
	var updates []string
	for op, fields := range diff {
		switch op {
		case "u": // Handle update operation
			for column, value := range fields.(map[string]interface{}) {
				updates = append(updates, fmt.Sprintf("%s = %v", column, value))
			}
		case "d": // Handle delete operation
			for column := range fields.(map[string]interface{}) {
				updates = append(updates, fmt.Sprintf("%s = NULL", column))
			}
		}
	}

	sql := fmt.Sprintf("UPDATE %s SET %s WHERE _id = '%v';", tableName, strings.Join(updates, ", "), id)
	return sql, nil
}
