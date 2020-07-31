package sql2struct

import (
	"fmt"
	"os"
	"text/template"

	"github.com/goTool/internal/word"
)

const structTpl = `type {{.TableName | ToCamelCase}} struct { 
	{{range .Columns}} {{ $length := len .Comment }} {{ if gt $length 0 }}
	// {{ .Comment }} {{ else }} // {{ .Name }} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{ .Name | ToCamelCase }} {{ .Type }}  {{.Tag}} {{ else }}{{ .Name }} {{ end }} {{ end }}
}	
	
func (model {{ .TableName | ToCamelCase }}) TableName() string {
	return "{{ .TableName }}"
}`

// StructTemplate 定义模板结构体
type StructTemplate struct {
	structTpl string
}

// StructColumn 定义列类型结构体
type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

// StructTemplateDB 定义DB模板
type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

// NewStructTemplate 初始化一个模板
func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

// Assemblycolumns 对查询COLUMNS表所组装得到的tbColumns进行进一步的分解和转换
func (t *StructTemplate) Assemblycolumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     fmt.Sprintf("`json:"+"%s"+"`", column.ColumnName),
			Comment: column.ColumnComment,
		})
	}

	return tplColumns
}

// Generate 声明并组装模版对象，使用Execute方法进行渲染
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToLowerCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}

	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}

	return nil
}
