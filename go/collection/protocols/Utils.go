package protocols

import "github.com/saichler/collect/go/types"

func SetValue(row int, col string, value []byte, tbl *types.Table) {
	if tbl == nil {
		return
	}
	if tbl.Rows == nil {
		tbl.Rows = make(map[uint32]*types.Row)
	}
	rowData, ok := tbl.Rows[uint32(row)]
	if !ok {
		rowData = &types.Row{}
		rowData.Data = make(map[string][]byte)
		tbl.Rows[uint32(row)] = rowData
	}
	rowData.Data[col] = value
	if value != nil && tbl.Columns[col] == "" {
		tbl.Columns[col] = col
	}
}

func Keys(m *types.Map) []string {
	if m == nil || m.Data == nil {
		return []string{}
	}
	result := make([]string, len(m.Data))
	i := 0
	for k, _ := range m.Data {
		result[i] = k
		i++
	}
	return result
}
