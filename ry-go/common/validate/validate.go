package validate

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	"ry-go/utils"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 使用了国际化就不需要额外处理原本的规则，只需要设置自定义转换信息即可
var validateMessageMap = map[string]string{
	// "email":    "邮箱格式不正确",
	"phone": "无效的手机号格式",
	// "required": "是必填项",
	// "oneof":    "必须是下列值之一：%s",
	// "min":      "最小值是：%s",
	// "max":      "最大值是：%s",
}

// Validate 创建一个全局的validator实例
var Validate *validator.Validate
var Trans ut.Translator

// PhoneValidation 自定义手机号验证
func PhoneValidation(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	// 使用正则表达式验证是否为中国大陆手机号
	re := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return re.MatchString(phone)
}

// 初始化验证器
func init() {
	Validate = validator.New()
	// 注册自定义手机号验证
	Validate.RegisterValidation("phone", PhoneValidation)
	if err := InitTrans("zh"); err != nil {
		log.Printf("init trans failed, err:%v\n", err)
	}
}

// ConvertValidateMessage 转换错误消息
func ConvertValidateMessage(err error) string {
	var msgList []string
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			tag := e.Tag()
			//field := utils.PascalCaseToCamel(e.Field())
			//param := e.Param() //最大最小范围这些限定值
			//log.Println("查看param", param)

			transMsg := e.Translate(Trans)
			log.Println("查看错误信息", transMsg)
			// 因为自定义的国际化之后依旧是英文的，这里判断如果是英文就使用自定义的错误信息
			if utils.IsAllEnglishIgnoreSymbols(transMsg) {
				if msg, exists := validateMessageMap[tag]; exists {
					transMsg = msg
				}
			}
			msgList = append(msgList, transMsg)

			/* if msg, exists := validateMessageMap[tag]; exists {
				if tag == "required" {
					msg = field + msg
				} else if tag == "oneof" || tag == "min" || tag == "max" {
					msg = field + fmt.Sprintf(msg, param)
				}
				msgList = append(msgList, msg)
			} else {
				// 如果没有定义特定的消息，使用默认消息
				msgList = append(msgList, e.Error())
			}  */
		}
		return strings.Join(msgList, ", ")
	}
	return err.Error()
}

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {

	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器

	// 第一个参数是备用（fallback）的语言环境
	// 后面的参数是应该支持的语言环境（支持多个）
	// uni := ut.New(zhT, zhT) 也是可以的
	uni := ut.New(enT, zhT, enT)

	// locale 通常取决于 http 请求头的 'Accept-Language'
	var ok bool
	// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
	Trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
	}

	// 注册翻译器
	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(Validate, Trans)
	case "zh":
		log.Printf("注册了中文翻译器:%s", locale)
		err = zhTranslations.RegisterDefaultTranslations(Validate, Trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(Validate, Trans)
	}
	Validate.RegisterTagNameFunc(CustomTagNameFunc)
	//Validate.RegisterTranslation("required", Trans, RequiredOverrideRegister, RequiredOverrideTranslation)
	return

}

// 我们加上了一个自定义标签，这个标签用于给结构体字段做中文名，它会替代原本的字段名称
func CustomTagNameFunc(field reflect.StructField) string {
	label := field.Tag.Get("label")
	if len(label) == 0 {
		return utils.PascalCaseToCamel(field.Name)
	}
	return label
}

func RequiredOverrideRegister(ut ut.Translator) error { //这个函数的作用是注册翻译模板
	return ut.Add("required", "{}是一个必须填写的字段", true) // {}是占位符 true代表了是否重写已有的模板
}

func RequiredOverrideTranslation(ut ut.Translator, fe validator.FieldError) string { // 这个函数的作用是负责翻译内容
	t, _ := ut.T("required", fe.Field()) //参数可以有多个，取决于注册对应Tag的模板有多少个占位符
	return t
}
