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
