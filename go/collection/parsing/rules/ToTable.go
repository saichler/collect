package rules

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/serializer/go/serialize/object"
	"github.com/saichler/shared/go/share/interfaces"
	"strings"
)

type ToTable struct{}

func (this *ToTable) Name() string {
	return "ToTable"
}

func (this *ToTable) ParamNames() []string {
	return []string{"columns"}
}

func (this *ToTable) Parse(resources interfaces.IResources, workSpace map[string]interface{}, params map[string]*types.Parameter, any interface{}) error {
	input, ok := workSpace[Input].(string)
	if !ok {
		return nil
	}
	colmns, err := getIntInput(workSpace, Columns)
	if err != nil {
		return err
	}
	lines := strings.Split(input, "\n")
	table := &types.Table{}
	table.Rows = make(map[int32]*types.Row)
	for i, line := range lines {
		if table.Columns == nil {
			table.Columns = getColumns(line)
			if len(table.Columns) != colmns {
				return resources.Logger().Error("Number of columns mismatch, expected:", colmns, ", actual:", len(table.Columns))
			}
			continue
		}
		row := &types.Row{}
		row.Data = getValues(line, table.Columns)
		table.Rows[int32(i)] = row
	}
	workSpace[Output] = table
	return nil
}

func getValues(line string, columns map[int32]string) map[int32][]byte {
	line = strings.TrimSpace(line)
	result := make(map[int32][]byte, 0)
	begin := 0
	for i := 0; i < len(columns); i++ {
		col := columns[int32(i)]
		if begin+len(col) > len(line) {
			result[int32(i)] = []byte{}
		} else {
			value := strings.TrimSpace(line[begin : begin+len(col)])
			obj := object.NewEncode([]byte{}, 0)
			obj.Add(value)
			result[int32(i)] = obj.Data()
			begin = begin + len(col)
		}
	}
	return result
}

func getColumns(line string) map[int32]string {
	columns := make(map[int32]string, 0)
	begin := 0
	open := false
	var count int32 = 0
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' && !open {
			open = true
		} else if open && line[i] != ' ' {
			open = false
			col := line[begin:i]
			columns[count] = col
			count++
			begin = i
		}
	}
	if begin < len(line)-1 {
		col := line[begin:len(line)]
		columns[count] = col
	}
	return columns
}
