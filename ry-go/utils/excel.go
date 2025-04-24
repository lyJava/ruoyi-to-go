package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"

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
