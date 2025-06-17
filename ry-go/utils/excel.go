package utils

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"

	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

// CreateExcel 创建excel
//
// 参数
//
//	-excelName: excel文件名
//
// 返回
//
//	-string: excel绝对路径
//	-error: 下载错误
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func CreateExcel(excelName string) (string, error) {
	f := excelize.NewFile()

	defer func() {
		// 关闭文件
		if err := f.Close(); err != nil {
			zap.L().Sugar().Errorf("excelize excel close error===%+v", err)
		}
	}()

	headers := []string{
		"快递名称", "快递单号", "取件码", "发件人姓名", "发件人手机",
		"发件人地址", "发件人身份证号", "创建人", "创建时间", "修改人", "修改时间",
	}
	// 设置表头
	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 1) // 列从 A 开始，行从 1 开始
		f.SetCellValue("Sheet1", cell, header)
	}

	colWidths := map[string]float64{
		"A": 30, // 快递名称
		"B": 20, // 快递单号
		"C": 20, // 取件码
		"D": 25, // 发件人姓名
		"E": 20, // 发件人手机
		"F": 30, // 发件人地址
		"G": 30, // 发件人身份证号
		"H": 25, // 创建人
		"I": 30, // 创建时间
		"J": 25, // 修改人
		"K": 30, // 修改时间
	}
	// 设置每列宽度
	for col, width := range colWidths {
		if err := f.SetColWidth("Sheet1", col, col, width); err != nil {
			zap.L().Sugar().Errorf("Failed to set column width===%+v", err)
			return "", err
		}
	}

	// 创建样式：设置背景色为浅蓝色
	styleID, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#ADD8E6"}, // 浅蓝色
		},
	})
	if err != nil {
		zap.L().Sugar().Errorf("Failed to set column width===%+v", err)
		return "", err
	}

	// 应用样式到标题行
	startCell, _ := excelize.CoordinatesToCellName(1, 1)          // 起始单元格（A1）
	endCell, _ := excelize.CoordinatesToCellName(len(headers), 1) // 结束单元格（K1）
	if err = f.SetCellStyle("Sheet1", startCell, endCell, styleID); err != nil {
		zap.L().Sugar().Errorf("Failed to set cell style===%+v", err)
		return "", err
	}

	/*f.SetCellValue("Sheet1", "A1", "快递名称")
	f.SetCellValue("Sheet1", "B1", "快递单号")
	f.SetCellValue("Sheet1", "C1", "取件码")
	f.SetCellValue("Sheet1", "D1", "发件人姓名")
	f.SetCellValue("Sheet1", "E1", "发件人手机")
	f.SetCellValue("Sheet1", "F1", "发件人地址")
	f.SetCellValue("Sheet1", "G1", "发件人身份证号")
	f.SetCellValue("Sheet1", "H1", "创建人")
	f.SetCellValue("Sheet1", "I1", "创建时间")
	f.SetCellValue("Sheet1", "J1", "修改人")
	f.SetCellValue("Sheet1", "K1", "修改时间")*/
	// 设置活动的工作表
	f.SetActiveSheet(0)

	currentPath, err := os.Getwd()
	if err != nil {
		zap.L().Sugar().Errorf("获取当前目录错误===%v", err)
		return "", err
	}
	zap.L().Sugar().Infof("获取当前工作目录路径===%s", currentPath)

	// 创建目标目录（如果不存在）
	excelDir := filepath.Join(currentPath, "excel")
	if err = os.MkdirAll(excelDir, os.ModePerm); err != nil {
		zap.L().Sugar().Errorf("Failed to create Excel directory===%+v", err)
		return "", err
	}

	// 保存 Excel 文件
	excelPath := filepath.Join(excelDir, excelName)
	if err = f.SaveAs(excelPath); err != nil {
		zap.L().Sugar().Errorf("保存excel错误===%+v", err)
		return "", err
	}

	zap.L().Sugar().Infof("Excel file saved successfully to===%s", excelPath)
	return excelPath, nil

}

// WriteDataToExcel 数据写入到Excel中
//
//goland:noinspection GoUnhandledErrorResult
func WriteDataToExcel[T any](fileName string, headers []string, list []T) (string, error) {
	f := excelize.NewFile()

	// 关闭文件
	defer func() {
		if err := f.Close(); err != nil {
			zap.L().Sugar().Errorf("excelize excel close error===%+v", err)
		}
	}()

	// 写入表头
	for i, header := range headers {
		cell := indexToColumn(i) + "1"
		f.SetCellValue("Sheet1", cell, header)
		width := len(header) * 2 // 根据表头内容设置宽度，这里简单的以字符数乘以2来设置宽度
		f.SetColWidth("Sheet1", cell, cell, float64(width))
	}

	// 写入数据
	for rowIndex, record := range list {
		rowIndex += 2 // 跳过表头，从第二行开始写入数据
		v := reflect.ValueOf(record)
		if v.Kind() == reflect.Ptr {
			v = v.Elem() // 如果是指针类型，则获取其指向的值
		}
		for colIndex := 0; colIndex < v.NumField(); colIndex++ {
			cell := indexToColumn(colIndex) + strconv.Itoa(rowIndex)
			// 获取字段值并转换为字符串
			var value string
			if v.Type().Field(colIndex).Name == "Id" {
				idPtr := v.Field(colIndex).Interface()
				if idPtr != nil {
					value = cast.ToString(idPtr)
				} else {
					value = "" // 如果指针为 nil，则将值设为空字符串
				}
			} else {
				value = fmt.Sprintf("%v", v.Field(colIndex).Interface())
			}

			f.SetCellValue("Sheet1", cell, value)

			// 写入数据到单元格
			f.SetCellValue("Sheet1", cell, value)

			// 根据写入内容的长度动态设置单元格的宽度加上一些额外的空间
			colWidth := len(value) + 5
			colName, _ := excelize.ColumnNumberToName(colIndex + 1)
			f.SetColWidth("Sheet1", colName, colName, float64(colWidth))
		}
	}

	currentPath, err := os.Getwd()
	if err != nil {
		zap.L().Sugar().Errorf("获取当前目录错误===%+v", err)
		return "", err
	}
	zap.L().Sugar().Infof("获取当前工作目录路径===%s", currentPath)

	filePath := filepath.Join(currentPath, fileName)

	// 确保目录存在
	dir := filepath.Dir(filePath)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		zap.L().Sugar().Errorf("创建目录失败: %+v", err)
		return "", err
	}

	// 保存文件
	if err = f.SaveAs(filePath); err != nil {
		zap.L().Sugar().Errorf("文件保存错误===%+v", err)
		return "", err
	}
	zap.L().Sugar().Infof("获取文件绝对路径===%s", filePath)
	return filePath, nil
}

func WriteDataToExcelBuffer[T any](headers []string, list []T, sheetName string, headerBgColor bool) (*bytes.Buffer, error) {
	if len(list) <= 0 {
		return nil, fmt.Errorf("数据列表为空，无法执行导出操作")
	}

	// 创建列宽记录器 (初始化值为表头宽度)
	f, maxColWidths, err := setHeaderContentAndWidth(sheetName, headers)
	if err != nil {
		zap.L().Sugar().Errorf("excel工作表===%s设置表头内容与记录宽度错误===%+v",sheetName, err)
		return nil, err
	}

	if headerBgColor {
		if err = setHeaderBgColor(f, sheetName, headers); err != nil {
			zap.L().Sugar().Errorf("excel工作表===%s设置表头背景色===%+v",sheetName, err)
			return nil, err
		}
	}

	// 写入数据
	for rowIndex, record := range list {
		v := reflect.ValueOf(record)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		if v.Kind() != reflect.Struct {
			return nil, fmt.Errorf("excel数据类型必须是结构体")
		}

		colIndex := 0
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := v.Type().Field(i)

			// 跳过不可导出的字段
			if !field.CanInterface() {
				continue
			}

			// 跳过标记为忽略的字段，这里是gorm查询忽略的的tag为-
			if tag := fieldType.Tag.Get("gorm"); tag == "-" {
				continue
			}

			// 处理嵌套结构体
			if field.Kind() == reflect.Struct {
				if _, isTime := field.Interface().(time.Time); isTime {
					// 时间类型走正常处理流程
				} else {
					for j := 0; j < field.NumField(); j++ {
						nestedField := field.Field(j)
						nestedType := field.Type().Field(j)

						if !nestedField.CanInterface() {
							continue
						}

						if tag := nestedType.Tag.Get("gorm"); tag == "-" {
							continue
						}

						// 确保不超过表头列数
						if colIndex >= len(headers) {
							break
						}

						colName := indexToColumn(colIndex)
						cell := colName + strconv.Itoa(rowIndex+2)

						value := formatCellValue(nestedField)

						if err = f.SetCellValue(sheetName, cell, value); err != nil {
							zap.L().Sugar().Errorf("excel工作表==%s单元格==%s赋值==%s错误===%+v", sheetName, cell, value, err)
							return nil, err
						}

						// 更新最大列宽
						maxColWidths[colName] = max(maxColWidths[colName], visualWidth(value))
						colIndex++
					}
					continue
				}
			}
			// 确保不超过表头列数
			if colIndex >= len(headers) {
				break
			}

			colName := indexToColumn(colIndex)
			cell := colName + strconv.Itoa(rowIndex+2)

			value := formatCellValue(field)

			if err = f.SetCellValue(sheetName, cell, value); err != nil {
				zap.L().Sugar().Errorf("excel工作表==%s单元格==%s赋值==%s错误===%+v", sheetName, cell, value, err)
				return nil, err
			}

			// 更新最大列宽
			maxColWidths[colName] = max(maxColWidths[colName], visualWidth(value))
			colIndex++
		}
	}

	// 设置自适应列宽
	setColumnAdaptiveWidth(f, sheetName, maxColWidths)

	// 写入到内存缓冲区
	//buf := new(bytes.Buffer)
	// 预分配1MB，可能提升速度并不明显
	var buf = bytes.NewBuffer(make([]byte, 0, 1<<20))
	if err = f.Write(buf); err != nil {
		zap.L().Sugar().Errorf("excel工作表==%s写入缓冲区错误===%+v", sheetName, err)
		return nil, err
	}
	return buf, nil
}

func formatCellValue(field reflect.Value) string {
	var value string
	switch field.Kind() {
	case reflect.Struct:
		if t, ok := field.Interface().(time.Time); ok {
			value = t.Format(time.DateTime)
		} else {
			value = fmt.Sprintf("%v", field.Interface())
		}
	default:
		value = fmt.Sprintf("%v", field.Interface())
	}
	return value
}

func setHeaderBgColor(f *excelize.File, sheetName string, headers []string) error {
	if len(headers) <= 0 {
		return errors.New("excel标题切片不能为空")
	}
	// 创建标题样式 - 灰色背景 + 加粗 + 居中
	headerStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#EBEBEB"}, // 黄色背景
		},
		Font: &excelize.Font{
			Bold: true, // 加粗
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center", // 水平居中
			Vertical:   "center", // 垂直居中
		},
	})
	if err != nil {
		zap.L().Sugar().Errorf("excel工作表==%s创建标题样式错误===%+v", sheetName, err)
		return fmt.Errorf("创建标题样式失败: %v", err)
	}

	for i := range headers {
		cell := indexToColumn(i) + "1"
		if err = f.SetCellStyle(sheetName, cell, cell, headerStyle); err != nil {
			zap.L().Sugar().Errorf("excel工作表==%s设置标题单元格==%s样式错误===%+v", sheetName, cell, err)
			return fmt.Errorf("设置标题单元格样式失败: %v", err)
		}
	}

	return nil
}

// setHeaderContentAndWidth 设置表格标题内容与记录宽度
func setHeaderContentAndWidth(sheetName string, headers []string) (*excelize.File, map[string]int, error) {
	// 确保工作表存在
	if sheetName == "" {
		sheetName = "Sheet1"
	}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			zap.L().Sugar().Errorf("关闭excel文件错误===%+v", err)
		}
	}()

	// 创建或激活指定名称的工作表
	defaultSheet := f.GetSheetName(0)

	// 检查目标工作表是否存在（兼容旧版本）
	targetSheetExists := false
	for i := 0; i < f.SheetCount; i++ {
		if f.GetSheetName(i) == sheetName {
			targetSheetExists = true
			break
		}
	}

	if !targetSheetExists {
		// 如果目标工作表不存在则创建
		sheetIndex, err := f.NewSheet(sheetName)
		if err != nil {
			return nil, nil, err
		}
		zap.L().Sugar().Infof("excel指定的工作表不存在，创建工作表===%s,索引===%d", sheetName, sheetIndex)
		// 创建额外的工作表，macOS预览的时候就能看到激活的工作表名称了
		_, _ = f.NewSheet("sheet2")
		// 删除默认创建的工作表（如果存在）
		if defaultSheet != sheetName && f.SheetCount > 1 {
			if err = f.DeleteSheet(defaultSheet); err != nil {
				zap.L().Sugar().Errorf("excel删除默认工作表==%s错误===%+v", defaultSheet, err)
				return nil, nil, err
			}
		}
		// 设置第一个工作表为激活状态
		f.SetActiveSheet(0)
	}

	maxColWidths := make(map[string]int)

	for i, header := range headers {
		colName := indexToColumn(i)
		cell := colName + "1"
		if err := f.SetCellValue(sheetName, cell, header); err != nil {
			return nil, nil, err
		}
		//f.SetColWidth(sheet, indexToColumn(i), indexToColumn(i), float64(len(header)*2))
		maxColWidths[colName] = max(maxColWidths[colName], visualWidth(header))
	}
	return f, maxColWidths, nil
}

func visualWidth(s string) int {
	width := 0
	for _, r := range s {
		if r > 127 {
			width += 2 // 中文等宽字符
		} else {
			width += 1 // 英文/数字
		}
	}
	return width
}

func setColumnAdaptiveWidth(f *excelize.File, sheetName string, maxColWidths map[string]int) {
	if len(maxColWidths) <= 0 {
		return
	}
	for colName, width := range maxColWidths {
		_ = f.SetColWidth(sheetName, colName, colName, float64(width+5))
	}
}

// indexToColumn 将索引转换为 Excel 列名
func indexToColumn(index int) string {
	result := ""
	for index >= 0 {
		result = string(rune('A'+index%26)) + result
		index = (index-index%26)/26 - 1
	}
	return result
}

// ReadExcelFile 读取Excel文件中的数据
//
//goland:noinspection GoUnhandledErrorResult
func ReadExcelFile(filePath string) ([][]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		zap.L().Sugar().Errorf("打开excel错误:%+v", err)
		return nil, errors.New("读取excel文件失败")
	}

	defer f.Close()

	// 获取所有工作表
	sheetList := f.GetSheetList()
	if len(sheetList) == 0 {
		zap.L().Sugar().Errorln("excel文件中没有工作表")
		return nil, errors.New("excel文件中没有工作表")
	}
	// 获取第一个工作表的名称
	sheetName := sheetList[0]

	// 读取所有行
	rows, err := f.GetRows(sheetName)
	if err != nil {
		zap.L().Sugar().Errorf("读取Excel所有行错误:%+v", err)
		return nil, errors.New("读取Excel文件数据失败")
	}

	// 忽略第一行
	if len(rows) > 1 {
		rows = rows[1:]
	}

	return rows, nil
}
