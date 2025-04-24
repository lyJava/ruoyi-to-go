package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"math"
	"net"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mssola/user_agent"
	"github.com/rs/xid"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

// ToJsonFormat 便于阅读的json
//
// 参数
//
//	-data: 泛型
//
// 返回
//
//	-string: 非紧凑型json
func ToJsonFormat(data any) string {
	if data == "" {
		return ""
	}

	// 检查是否是空字符串
	if s, ok := data.(string); ok && s == "" {
		return ""
	}

	marshal, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		zap.L().Sugar().Errorf("数据格式化json错误===%+v", err)
		return ""
	}

	str := string(marshal)
	// 检查是否是空对象
	if str == "{}" || str == "null" {
		return ""
	}

	return str
}

// ToJson 紧凑型json
//
// 参数
//
//	-data: 泛型
//
// 返回
//
//	-string: 紧凑型json
func ToJson(data any) string {
	if data == "" {
		return ""
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		zap.L().Sugar().Errorf("转换json错误===%+v", err)
		return ""
	}
	return string(marshal)
}

// ForcedToJsonFormat 强制json格式化输出
//
// 参数
//
//	-b: 字节
//
// 返回
//
//	-string: json字符串
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func ForcedToJsonFormat(b []byte) string {
	// 1. 检查输入是否为空
	if len(b) == 0 {
		zap.L().Sugar().Debug("Empty input provided")
		return ""
	}

	// 2. 检查是否是有效的JSON（可选）
	if !json.Valid(b) {
		zap.L().Sugar().Warnf("Invalid JSON input: %q", string(b))
		return string(b) // 返回原始字符串或处理错误的其他方式
	}
	var result any
	// 将JSON字符串解码到result变量中
	err := json.Unmarshal(b, &result)
	if err != nil {
		zap.L().Sugar().Errorf("Error occurred during unmarshaling. Error: %+v, Input: %q", err, string(b))
		return string(b)
	}

	// 格式化输出JSON字符串
	formattedJson, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		zap.L().Sugar().Errorf("Error occurred during marshaling. Error:%+v", err)
		return string(b)
	}
	return string(formattedJson)
}

// ConvertJsonStrToSlice 将json字符串转换为切片
//
// 参数
//
//	-jsonStr: json字符串
//
// 返回
//
//	-[]any: 泛型切片
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func ConvertJsonStrToSlice(jsonStr string) []any {
	if jsonStr == "" {
		return nil
	}
	// Step 1: 将 JSON 字符串解析为 `any`
	var rawData any
	err := json.Unmarshal([]byte(jsonStr), &rawData)
	if err != nil {
		zap.L().Sugar().Errorf("Error occurred json Unmarshal marshaling. Error:%+v", err)
		return nil
	}

	// Step 2: 使用 cast.ToSlice 将 `interface{}` 转换为切片
	return cast.ToSlice(rawData)
}

// GetCurrentAbPathByCaller 获取当前执行文件路径
//
// 返回
//
//	-string: 当前执行文件路径
//
// 返回
//
//	-string: 执行文件路径
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// Contains 切片是否包含某个元素
//
// 参数
//
//	-slice: 泛型切片
//	-e: 泛型
//
// 返回
//
//	-bool: 是否包含
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func Contains(slice []any, e any) bool {
	/*
		// 这里的s为any的时候需要使用反射判断
		sv := reflect.ValueOf(s)
		if sv.Kind() == reflect.Slice {
			for i := 0; i < sv.Len(); i++ {
				if sv.Index(i).Interface() == e {
					return true
				}
			}
		}
		return false*/
	if len(slice) > 1000 {
		// 如果切片长度超过 1000，使用map方式
		return ContainsMap(slice, e)
	} else {
		for _, v := range slice {
			if reflect.DeepEqual(v, e) {
				return true
			}
		}
		return false
	}
}

func StringSliceToAnySlice(slice []string) []any {
	result := make([]any, 0, len(slice))
	for _, str := range slice {
		result = append(result, str)
	}
	return result
}

// MapContains Map是否包含某个元素
//
// 参数
//
//	-slice: 泛型切片
//	-e: 泛型
//
// 返回
//
//	-bool: 是否包含
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func MapContains[T any](mapData map[string]T, e any) bool {
	for _, item := range mapData {
		if reflect.DeepEqual(item, e) {
			return true
		}
	}
	return false
}

// MapContainsKey 判断 map[string]any 是否包含指定的 key
func MapContainsKey[T any](mapData map[string]T, key string) bool {
	_, exists := mapData[key]
	return exists
}

// CamelToSnakeCase 将驼峰转下划线方式
func CamelToSnakeCase(s string) string {
	var buffer bytes.Buffer
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// PascalCaseToCamel 将首字母大写首字母大写的驼峰命名转换首字母小写的驼峰命名
func PascalCaseToCamel(s string) string {
	if s == "" || len(s) == 0 {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// CamelCaseToSnakeCase 驼峰转换首字母小写下划线
//
// 参数
//
//	-camelCase: 驼峰字符串
//
// 返回
//
//	-string: 首字母小写下划线
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func CamelCaseToSnakeCase(camelCase string) string {
	var matchFirstCap = regexp.MustCompile("([A-Z])")
	snake := matchFirstCap.ReplaceAllString(camelCase, "_$1")
	snake = strings.ToLower(snake)
	return snake
}

// SnakeToCamel 将带有下划线的字符串转换为驼峰并且首字母小写
// 参数
//
//	-s: 下划线字符串
//
// 返回
//
//	-string: 小写开头驼峰字符串
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func SnakeToCamel(s string) string {
	var buffer bytes.Buffer
	upperNext := false

	for i, r := range s {
		if r == '_' {
			upperNext = true
		} else {
			if upperNext {
				buffer.WriteRune(unicode.ToUpper(r))
				upperNext = false
			} else {
				if i == 0 {
					buffer.WriteRune(unicode.ToLower(r))
				} else {
					buffer.WriteRune(r)
				}
			}
		}
	}

	return buffer.String()
}

// ContainsMap 函数用于判断切片是否包含指定元素，使用 map 的方式
// 参数
//
//	-list: 泛型切片
//	-e: 泛型对象
//
// 返回
//
//	-bool: 是/否
func ContainsMap(list []any, e any) bool {
	m := make(map[any]bool)
	for _, v := range list {
		m[v] = true
	}
	return m[e]
}

// IsAllEnglish 是否都是英文
//
// 参数
//   - s: 字符串
//
// 返回
//
//	-bool: 是/否
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func IsAllEnglish(s string) bool {
	reg, err := regexp.Compile("^[A-Za-z]+$")
	if err != nil {
		panic(err) // 处理正则表达式编译错误
	}
	return reg.MatchString(s)
}

// IsAllEnglishIgnoreSymbols 是否都是英文(忽略符号及其他的)
//
// 参数
//   - s: 字符串
//
// 返回
//
//	-bool: 是/否
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func IsAllEnglishIgnoreSymbols(s string) bool {
	// Remove all symbols and numbers
	re := regexp.MustCompile("[^A-Za-z]+")
	// Replace all non-letter characters with an empty string
	cleanedString := re.ReplaceAllString(s, "")

	// Check if the cleaned string contains only English letters
	return len(cleanedString) > 0 && regexp.MustCompile("^[A-Za-z]+$").MatchString(cleanedString)
}

// FormatFileSize 格式化转换文件大小
// 参数
//
//	-size: 文件大小(字节)
//
// 返回
//
//	-string: 转换计算后的大小
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func FormatFileSize(size int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/KB)
	default:
		return fmt.Sprintf("%d B", size)
	}
}

// BuildArgsWithBrackets 将泛型切片类型的参数列表转换为一个用括号括起来的、以逗号分隔的字符串
//
// 参数
//
//	-list: 泛型切片
//
// 返回
//
//	-string: 用括号括起来的、以逗号分隔的字符串
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func BuildArgsWithBrackets(list []any) string {
	if len(list) == 0 {
		return ""
	}
	builder := strings.Builder{}
	// 预分配容量，提升性能
	builder.Grow(len(list))
	builder.WriteString("(")

	for i, arg := range list {
		builder.WriteString(fmt.Sprintf("%v", arg))
		if i < len(list)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(")")
	return builder.String()
}

// GeneratePlaceholder 生成占位符字符串，例如: $1, $2, $3或? ,? ,?
// 参数
//
//	-n: 占位符生成数量
//	-format: 占位符格式化字符串，比如：$%d或者?
//
// 返回
//
//	-string: 占位符字符串$1, $2, $3或者? ,? ,?
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func GeneratePlaceholder[T int | int64](count T, format string) string {
	if count <= 0 {
		return ""
	}
	//var placeholders []string
	list := make([]string, 0, count)
	for i := 1; i <= int(count); i++ {
		list = append(list, fmt.Sprintf(format, i))
	}
	return strings.Join(list, ", ")
}

// GeneratePlaceholderWithQuestionMark 生成占位符字符串，例如: ?, ?, ?
// 参数
//
//	n (int): 生成的数量
//
// 返回
//
//	-string: 占位符字符串
//
//goland:noinspection GoUnhandledErrorResult,GoUnusedExportedFunction
func GeneratePlaceholderWithQuestionMark[T int | int64](n T) string {
	if n <= 0 {
		return ""
	}
	// 预分配切片容量，避免多次内存分配
	placeholders := make([]string, 0, n)
	for i := 1; i <= int(n); i++ {
		placeholders = append(placeholders, "?")
	}
	return strings.Join(placeholders, ", ")
}

// SetPageDefault 设置分页参数默认值
func SetPageDefault(page, size int64) (int64, int64) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	return page, size
}

// CalcTotalPage 计算总页数
func CalcTotalPage(total, size int64) int64 {
	return int64(int(math.Ceil(float64(total) / float64(size))))
}

// GetValForContext 从上下文获取值
//
// 参数:
//
//   - ctx:上下文
//   - key:键
//
// 返回:
//   - string:获取的值
func GetValForContext(ctx context.Context, key string) string {
	val, ok := ctx.Value(key).(string)
	if !ok {
		zap.L().Sugar().Error("requestId not found or invalid type in context")
		return ""
	}
	return val
}

// SetValForContext 上下文设置值
//
// 参数:
//
//   - ctx:上下文
//   - key：键
//   - value:值
//
// 返回:
//   - context：context.Context  上下文
func SetValForContext(ctx context.Context, key, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

// NormalUUID 返回常规UUID
func NormalUUID() string {
	return uuid.NewString()
}

// ShortUUID 返回短UUID
func ShortUUID() string {
	return xid.New().String()
}

type MyStruct struct {
	// Name is the name of the person
	Name string `json:"name,string,omitempty" excel:"yes"` // 名称
	// Age is the age of the person
	Age     int    `json:"age,omitempty" excel:"yes"`     //年龄
	Email   string `json:"email,omitempty" excel:"yes"`   // 邮箱
	Address string `json:"address,omitempty" excel:"yes"` // 地址
	SysCode string `json:"sysCode,omitempty" excel:"yes"` // 系统编码
}

// FieldInfo 存储字段的注释和标签
type FieldInfo struct {
	Comment string
	Tags    map[string]string
}

// GetStructComment2 返回字段注释、标签以及字段顺序
func GetStructComment2(packagePath, structName string) (map[string]FieldInfo, []string) {
	// packagePath 如果是 main 的话，直接设置为 . 即可
	// 创建一个新的文件集
	fset := token.NewFileSet()

	// 解析当前目录下的 Go 文件
	pkgs, err := parser.ParseDir(fset, packagePath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// 存储字段信息
	fieldInfoMap := make(map[string]FieldInfo)
	// 存储字段顺序
	var fieldOrder []string

	// 遍历解析出的包
	for _, pkg := range pkgs {
		// 使用 go/doc 包生成文档
		docPkg := doc.New(pkg, "./", doc.AllDecls)

		// 遍历包中的类型
		for _, typeSpec := range docPkg.Types {
			// 找到我们感兴趣的结构体
			if typeSpec.Name == structName {
				// 遍历结构体的字段
				for _, field := range typeSpec.Decl.Specs {
					if typeDef, ok := field.(*ast.TypeSpec); ok {
						if structType, ok := typeDef.Type.(*ast.StructType); ok {
							// 遍历结构体的字段
							for _, field := range structType.Fields.List {
								// 获取字段名
								if len(field.Names) > 0 {
									fieldName := field.Names[0].Name
									// 记录字段顺序
									fieldOrder = append(fieldOrder, fieldName)

									// 获取字段注释（行内注释）
									var comment string
									if field.Comment != nil {
										var comments []string
										for _, comment := range field.Comment.List {
											cleanedComment := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))
											comments = append(comments, cleanedComment)
										}
										comment = strings.Join(comments, " ")
									}

									// 获取字段标签
									tags := make(map[string]string)
									if field.Tag != nil {
										tag := strings.Trim(field.Tag.Value, "`")
										structTag := reflect.StructTag(tag)

										// 提取 json 标签
										if jsonTag := structTag.Get("json"); jsonTag != "" {
											tags["json"] = jsonTag
										}

										// 提取 excel 标签
										if excelTag := structTag.Get("excel"); excelTag != "" {
											tags["excel"] = excelTag
										}
									}

									// 存储字段信息
									fieldInfoMap[fieldName] = FieldInfo{
										Comment: comment,
										Tags:    tags,
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return fieldInfoMap, fieldOrder
}

func GetStructComment(packagePath, structName string) (map[string]string, map[string]map[string]string, []string) {
	// packagePath 如果是main的话，直接设置为.即可
	// 创建一个新的文件集
	fileSet := token.NewFileSet()

	// 解析当前目录下的 Go 文件
	pkgMap, err := parser.ParseDir(fileSet, packagePath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	structMap := make(map[string]string)
	structTagMap := make(map[string]map[string]string)
	// 存储字段顺序
	var fieldOrder []string

	// 遍历解析出的包
	for _, pkg := range pkgMap {
		// 使用 go/doc 包生成文档
		docPkg := doc.New(pkg, "./", doc.AllDecls)

		// 遍历包中的类型
		for _, typeSpec := range docPkg.Types {
			// 找到我们感兴趣的结构体
			// if typeSpec.Name == "MyStruct" {
			if typeSpec.Name == structName {
				// 遍历结构体的字段
				for _, field := range typeSpec.Decl.Specs {
					if typeDef, ok := field.(*ast.TypeSpec); ok {
						if structType, ok := typeDef.Type.(*ast.StructType); ok {
							if !ok {
								continue // 如果不是结构体，跳过
							}

							// 遍历结构体的字段
							for _, field := range structType.Fields.List {
								// 获取字段名
								if len(field.Names) > 0 {
									fieldName := field.Names[0].Name
									// 记录字段顺序
									fieldOrder = append(fieldOrder, fieldName)
									// 获取字段注释（field.Doc获取 紧邻字段的注释（字段上方的注释）才会被解析为字段注释）
									// 注释在字段右侧（行内注释），需要使用 field.Comment
									/*if field.Doc != nil {
										var comments []string
										for _, comment := range field.Doc.List {
											comments = append(comments, strings.TrimSpace(comment.Text))
										}
										// 将注释拼接成一个字符串
										structMap[fieldName] = strings.Join(comments, " ")
									} else {
										// 如果没有注释，设置为空字符串
										structMap[fieldName] = ""
									}*/
									if field.Comment != nil {
										var comments []string
										for _, comment := range field.Comment.List {
											cleanedComment := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))
											comments = append(comments, cleanedComment)
										}
										// 将注释拼接成一个字符串
										structMap[fieldName] = strings.Join(comments, " ")
									} else {
										// 如果没有注释，设置为空字符串
										structMap[fieldName] = ""
									}

									// 获取字段标签
									if field.Tag != nil {
										// 初始化 structTagMap[fieldName]
										if structTagMap[fieldName] == nil {
											structTagMap[fieldName] = make(map[string]string)
										}

										// 初始化 structTagMap[fieldName]["tags"]
										if structTagMap[fieldName] == nil {
											structTagMap[fieldName] = make(map[string]string)
										}

										tag := strings.Trim(field.Tag.Value, "`")
										structTag := reflect.StructTag(tag)

										// 提取 json 标签
										if jsonTag := structTag.Get("json"); jsonTag != "" {
											structTagMap[fieldName]["json"] = jsonTag
										}

										// 提取 excel 标签
										if excelTag := structTag.Get("excel"); excelTag != "" {
											structTagMap[fieldName]["excel"] = excelTag
											var comments []string
											for _, comment := range field.Comment.List {
												cleanedComment := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))
												comments = append(comments, cleanedComment)
											}
											// 将注释拼接成一个字符串
											structTagMap[fieldName]["comment"] = strings.Join(comments, " ")
										}
									}

								}

							}
						}
					}
				}
			}
		}
	}

	return structMap, structTagMap, fieldOrder
}

func BuildExcelHeadersAndColumns(tageMap map[string]map[string]string, fieldOrder []string) ([]string, []string) {
	var columns []string
	var excelHeaders []string
	for _, fieldName := range fieldOrder {
		itemMap := tageMap[fieldName]
		columns = append(columns, CamelToSnakeCase(fieldName))
		for key, value := range itemMap {
			if key == "excel" && value == "yes" {
				excelHeaders = append(excelHeaders, itemMap["comment"])
			}
		}
	}
	return columns, excelHeaders
}

// GetTimeDiff 计算两个日期相差的天数，小时，分钟
func GetTimeDiff(start, end time.Time) (days, hours, minutes int) {
	// 确保两个时间在同一时区
	start = start.UTC()
	end = end.UTC()

	// 计算总差值
	diff := end.Sub(start)

	// 计算天数
	totalHours := diff.Hours()
	days = int(totalHours / 24)

	// 计算剩余小时
	remainingHours := totalHours - float64(days)*24
	hours = int(remainingHours)

	// 计算剩余分钟
	remainingMinutes := diff.Minutes() - float64(days)*24*60 - float64(hours)*60
	minutes = int(remainingMinutes)

	return
}

// GetTimeDiffWithSign 负时间差
func GetTimeDiffWithSign(start, end time.Time) (days, hours, minutes int, isNegative bool) {
	if start.After(end) {
		days, hours, minutes = GetTimeDiff(end, start)
		return days, hours, minutes, true
	}
	days, hours, minutes = GetTimeDiff(start, end)
	return days, hours, minutes, false
}

// GetUserAgent 获取登陆用户浏览器及IP信息
func GetUserAgent(c echo.Context) map[string]any {
	// 获取登陆用户浏览器及IP信息
	userAgent := c.Request().Header.Get("User-Agent")
	if userAgent != "" {
		uaParse := user_agent.New(userAgent)
		uaParse.Parse(userAgent)
		browser, version := uaParse.Browser()
		osInfo := uaParse.OSInfo()

		userAgentMap := map[string]any{
			"browser": browser + "-" + version,
			"ip":      GetClientIP(c),
			"mobile":  uaParse.Mobile(),
			"os":      osInfo.Name + "-" + osInfo.Version,
		}
		return userAgentMap
	}

	return nil
}

// GetClientIP 获取客户端IP
func GetClientIP(c echo.Context) string {
	// 这里不需要区分大小写
	headers := []string{
		"X-Forwarded-For",
		"Proxy-Client-IP",
		"WL-Proxy-Client-IP",
		"X-Real-IP",
	}

	// 按顺序检查每个 Header
	for _, header := range headers {
		value := c.Request().Header.Get(header)
		if value == "" {
			continue
		}

		// 处理多个 IP 的情况（兼容逗号和空格分隔）
		ips := strings.FieldsFunc(value, func(r rune) bool {
			return r == ',' || r == ' '
		})

		// 寻找第一个有效 IP
		for _, ip := range ips {
			ip = strings.TrimSpace(ip)
			if IsValidIP(ip) {
				return SanitizeIP(ip)
			}
		}
	}

	// 回退到 RemoteAddr
	remoteAddr := c.Request().RemoteAddr
	if ip, _, err := net.SplitHostPort(remoteAddr); err == nil {
		return SanitizeIP(ip)
	}
	return SanitizeIP(remoteAddr)
}

// IsValidIP IP 有效性验证
func IsValidIP(ip string) bool {
	parsed := net.ParseIP(ip)
	return parsed != nil &&
		!strings.EqualFold(ip, "unknown") &&
		!parsed.IsUnspecified()
}

// SanitizeIP IP 清洗和标准化
func SanitizeIP(ip string) string {
	// 处理 IPv6 本地地址
	if ip == "::1" {
		return "127.0.0.1"
	}

	// 去除端口号（应对未分割的情况）
	if host, _, err := net.SplitHostPort(ip); err == nil {
		ip = host
	}

	// 转为标准字符串格式
	if parsed := net.ParseIP(ip); parsed != nil {
		return parsed.String()
	}

	// 简单清理非法字符（仿 EscapeUtil.clean）
	return strings.Map(func(r rune) rune {
		if r == ':' || r == '.' || r >= '0' && r <= '9' {
			return r
		}
		if r >= 'a' && r <= 'f' || r >= 'A' && r <= 'F' {
			return r
		}
		return -1
	}, ip)
}

// OsInfo 获取系统信息
func OsInfo() (*host.InfoStat, []cpu.InfoStat, *mem.VirtualMemoryStat, []disk.PartitionStat) {
	// 系统基本信息
	hostInfo, _ := host.Info()
	fmt.Printf(`
=== 系统概览 ===
主机名:       %s
操作系统:     %s %s
内核版本:     %s
启动时间:     %s
运行时长:     %.0d 小时
`,
		hostInfo.Hostname,
		hostInfo.Platform, hostInfo.PlatformVersion,
		hostInfo.KernelVersion,
		time.Unix(int64(hostInfo.BootTime), 0).Format(time.DateTime),
		hostInfo.Uptime/3600,
	)

	// CPU 信息
	cpuInfo, _ := cpu.Info()
	cpuPercent, _ := cpu.Percent(time.Second, true)
	physicalCores, _ := cpu.Counts(false)
	logicalCores, _ := cpu.Counts(true)
	fmt.Printf(`
=== CPU 信息 ===
物理核心:     %d 核
逻辑核心:     %d 线程
型号:        %s
当前使用率:   %.1f%%
`,
		physicalCores,
		logicalCores,
		cpuInfo[0].ModelName,
		cpuPercent[0],
	)

	// 内存信息
	memStat, _ := mem.VirtualMemory()
	fmt.Printf(`
=== 内存信息 ===
总内存:      %.2f GB
已用内存:    %.2f GB
使用率:      %.1f%%
`,
		float64(memStat.Total)/1024/1024/1024,
		float64(memStat.Used)/1024/1024/1024,
		memStat.UsedPercent,
	)

	// 磁盘信息
	partitions, _ := disk.Partitions(false)
	for _, p := range partitions {
		usage, _ := disk.Usage(p.Mountpoint)
		fmt.Printf(`
磁盘 %s (%s)
挂载点: %s
总空间: %.2f GB
剩余空间: %.2f GB
`,
			p.Device, p.Fstype,
			p.Mountpoint,
			float64(usage.Total)/1024/1024/1024,
			float64(usage.Free)/1024/1024/1024,
		)
	}

	// 系统负载（Linux/Mac）
	if runtime.GOOS != "windows" {
		avg, _ := load.Avg()
		fmt.Printf(`
=== 系统负载 ===
1分钟负载:   %.2f
5分钟负载:   %.2f
15分钟负载:  %.2f
`,
			avg.Load1, avg.Load5, avg.Load15)
	}

	return hostInfo, cpuInfo, memStat, partitions
}

// TimeForHuman 格式化日期
//
// 参数
//
//	timeValue (int64): 10位时间戳
func TimeForHuman(t time.Time) string {

	now := time.Now().Local() // 获取本地时区当前时间
	diffTime := now.Unix() - t.Unix()

	if diffTime < 0 {
		return t.Format("2006-01-02")
	}

	const (
		second = 1
		minute = 60 * second
		hour   = 60 * minute
		day    = 24 * hour
		week   = 7 * day
		month  = 30 * day
		year   = 365 * day
	)

	switch {
	case diffTime < 5:
		return "刚刚"
	case diffTime < minute:
		return fmt.Sprintf("%d秒前", diffTime)
	case diffTime < hour:
		return fmt.Sprintf("%d分钟前", diffTime/minute)
	case diffTime < day:
		return fmt.Sprintf("%d小时前", diffTime/hour)
	case diffTime < week:
		return fmt.Sprintf("%d天前", diffTime/day)
	case diffTime < month:
		weeks := diffTime / week
		days := (diffTime % week) / day
		if days > 0 {
			return fmt.Sprintf("%d周%d天前", weeks, days)
		}
		return fmt.Sprintf("%d周前", weeks)
	case diffTime < year:
		return fmt.Sprintf("%d月前", diffTime/month)
	default:
		return fmt.Sprintf("%d年前", diffTime/year)
	}
}

func ToAnySlice[T any](list []T) []any {
	if len(list) <= 0 {
		return nil
	}
	result := make([]any, len(list))
	for i, v := range list {
		result[i] = v
	}
	return result
}

func ToInt64Slice[T any](list []T) []int64 {
	if len(list) <= 0 {
		return nil
	}
	result := make([]int64, len(list))
	for i, v := range list {
		result[i] = cast.ToInt64(v)
	}
	return result
}

// Deduplicate 切片去重
func Deduplicate[T any, K comparable](list []T, keyFunc func(T) K) []T {
	if len(list) <= 0 {
		return nil
	}
	seen := make(map[K]struct{})
	result := make([]T, 0, len(list))
	for _, item := range list {
		key := keyFunc(item)
		if _, exists := seen[key]; !exists {
			seen[key] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// DeduplicateStrings 字符串切片去重
func DeduplicateStrings(list []string) []string {
	return Deduplicate(list, func(s string) string { return s })
}

type PageResult[T any] struct {
	Records int `json:"records"`
	Content []T `json:"content"`
}

func BuildPageBySilence[T any](list []T, pageNumber, pageSize int64) PageResult[T] {
	total := len(list)

	if pageSize <= 0 || pageNumber <= 0 || len(list) == 0 {
		return PageResult[T]{Records: total}
	}

	start := (pageNumber - 1) * pageSize
	if start >= int64(len(list)) {
		return PageResult[T]{Records: total, Content: []T{}}
	}

	// 计算结束索引（不超过数据长度）
	end := start + pageSize
	if end > int64(len(list)) {
		end = int64(len(list))
	}

	return PageResult[T]{
		Records: total,
		Content: list[start:end],
	}

}
