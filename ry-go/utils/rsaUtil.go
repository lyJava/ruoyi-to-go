package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"os"
)

// previewKey 安全展示密钥前缀（防止日志泄露敏感信息）
func previewKey(data []byte, length int) string {
	if len(data) < length {
		return string(data)
	}
	return string(data[:length]) + "..."
}

// LoadPrivateKey 获取私钥内容（自动处理不同环境的路径）
func LoadPrivateKey(privatePem string) ([]byte, error) {
	currentPath, err := GetCurrentPath()
	if err != nil {
		return nil, err
	}

	relPath := filepath.Join(currentPath, "resources", privatePem)
	return os.ReadFile(relPath)
}

// GenerateRSAKey 生成RSA密钥对并写入指定文件
//
// 参数:
//   - bits: RSA密钥位数（必须 >= 2048）
//   - publicKeyPath: 公钥文件路径
//   - privateKeyPath: 私钥文件路径
//
// 返回:
//   - error: 包含上下文信息的错误对象
//
// 安全说明:
//   - 公钥权限设置为 0444（只读）
//   - 私钥权限设置为 0400（仅所有者可读）
func GenerateRSAKey(bits int, publicKeyPath, privateKeyPath string) error {

	if publicKeyPath == "" || privateKeyPath == "" {
		err := errors.New("必须指定公钥和私钥文件路径")
		return err
	}

	currentPath, err := GetCurrentPath()
	if err != nil {
		return fmt.Errorf("获取当前目录错误===%w", err)
	}

	// 定义错误收集器
	var errs []string

	// 检查公钥文件
	filePathPublic := filepath.Join(currentPath, publicKeyPath)
	if exists, err1 := checkFileExists(filePathPublic); err != nil {
		return fmt.Errorf("检查公钥文件失败: %w", err1)
	} else if exists {
		errs = append(errs, fmt.Sprintf("公钥文件已存在: %s", filePathPublic))
	}

	// 检查私钥文件
	filePathPrivate := filepath.Join(currentPath, privateKeyPath)
	if exists, err2 := checkFileExists(filePathPrivate); err2 != nil {
		return fmt.Errorf("检查私钥文件失败: %w", err2)
	} else if exists {
		errs = append(errs, fmt.Sprintf("私钥文件已存在: %s", filePathPrivate))
	}

	// 统一处理所有存在性错误
	if len(errs) > 0 {
		// 使用 errors.Join 合并多个错误
		return fmt.Errorf("%s", strings.Join(errs, ";"))
	}

	// 创建内存缓冲区
	publicKeyBuf := &bytes.Buffer{}
	privateKeyBuf := &bytes.Buffer{}

	if err = GenerateRSAKeyPair(bits, publicKeyBuf, privateKeyBuf); err != nil {
		return err
	}

	if err = os.WriteFile(publicKeyPath, publicKeyBuf.Bytes(), 0644); err != nil {
		return err
	}

	if err = os.WriteFile(privateKeyPath, privateKeyBuf.Bytes(), 0600); err != nil {
		return err
	}
	return nil
}

// GenerateRSAKeyPair 生成RSA公私钥对并写入指定Writer（PKCS#1私钥 + PKIX公钥）
// 参数:
//   - bits: RSA密钥位数（建议 >= 2048）
//   - publicKeyWriter: 公钥写入目标（通常为文件或内存Buffer）
//   - privateKeyWriter: 私钥写入目标
//
// 返回:
//   - error: 成功返回nil，失败返回错误并记录日志
func GenerateRSAKeyPair(bits int, publicKeyWriter, privateKeyWriter io.Writer) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509PrivateKey,
	}

	if err = pem.Encode(privateKeyWriter, privateBlock); err != nil {
		return err
	}

	publicKey := &privateKey.PublicKey
	x509PublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509PublicKey,
	}

	err = pem.Encode(publicKeyWriter, publicBlock)
	if err != nil {
		return err
	}
	return nil
}

// EncryptOAEP 加密
func EncryptOAEP(publicKey, password []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		err := fmt.Errorf("failed to parse certificate PEM")
		return nil, err
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey := pub.(*rsa.PublicKey)

	h := sha256.New()
	return rsa.EncryptOAEP(h, rand.Reader, rsaPublicKey, password, nil)
}

// DecryptOAEP 使用RSA-OAEP算法解密数据，支持PKCS#1和PKCS#8格式私钥
//
// 参数:
//   - privateKeyPEM: PEM编码的RSA私钥（支持PKCS#1/PKCS#8格式）
//   - cipherData: 待解密的密文数据
//
// 返回:
//   - []byte: 解密后的明文数据
//   - error: 包含上下文信息的错误对象
//
// 注意事项:
//   - 使用SHA-256作为哈希函数，与OAEP默认实现一致
//   - PKCS#8私钥必须是RSA类型，其他类型会返回错误
//   - 密文长度必须等于密钥长度（单位：字节）
func DecryptOAEP(privateKeyPEM, cipherData []byte) ([]byte, error) {
	// 1. 解码PEM块
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		return nil, errors.New("invalid PEM encoding: no PEM blocks found")
	}

	// 2. 解析私钥（支持PKCS#1和PKCS#8格式）
	var (
		privateKey *rsa.PrivateKey
		err        error
	)

	switch block.Type {
	case "RSA PRIVATE KEY": // PKCS#1
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	case "PRIVATE KEY": // PKCS#8
		key, parseErr := x509.ParsePKCS8PrivateKey(block.Bytes)
		if parseErr != nil {
			err = parseErr
			break
		}
		var ok bool
		if privateKey, ok = key.(*rsa.PrivateKey); !ok {
			return nil, errors.New("PKCS#8 private key is not RSA type")
		}
	default:
		return nil, fmt.Errorf("unsupported PEM type: %q", block.Type)
	}

	if err != nil {
		return nil, fmt.Errorf("private key parse failed: %w", err)
	}

	// 3. 验证密文长度
	keySize := privateKey.Size()
	if len(cipherData) != keySize {
		return nil, fmt.Errorf("ciphertext length mismatch: need %d bytes, got %d", keySize, len(cipherData))
	}

	// 4. 创建哈希实例（使用SHA256）
	hash := sha256.New()

	// 5. 执行OAEP解密
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, cipherData, nil)
	if err != nil {
		return nil, fmt.Errorf("OAEP decryption failed: %w", err)
	}

	return plaintext, nil
}

// RsaDecrypt 使用RSA私钥解密数据，支持 PKCS#1 和 PKCS#8 格式私钥
//
// 参数:
//   - privateKeyPEM: PEM编码的私钥数据
//   - cipherText: 待解密的密文
//
// 返回:
//   - []byte: 解密后的明文
//   - error: 错误信息，包含具体的失败原因
//
// 注意事项:
//   - PKCS#8 私钥必须是RSA类型，其他类型（如ECDSA）会返回错误
//   - 建议使用 OAEP 解密方式（需调用方自行实现），PKCS1v15 存在已知风险
func RsaDecrypt(privateKeyPEM []byte, cipherText []byte) ([]byte, error) {
	// 解码PEM块
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		zap.L().Error("无效的PEM格式私钥")
		return nil, errors.New("invalid PEM encoding")
	}

	var (
		privateKey *rsa.PrivateKey
		err        error
	)

	// 根据PEM类型选择解析方式
	switch block.Type {
	// PKCS#1
	case "RSA PRIVATE KEY":
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	// PKCS#8
	case "PRIVATE KEY":
		key, parseErr := x509.ParsePKCS8PrivateKey(block.Bytes)
		if parseErr != nil {
			err = parseErr
			break
		}
		var ok bool
		if privateKey, ok = key.(*rsa.PrivateKey); !ok {
			zap.L().Error("PKCS#8私钥不是RSA类型", zap.String("key_type", fmt.Sprintf("%T", key)))
			return nil, errors.New("non-RSA PKCS#8 private key")
		}
	default:
		zap.L().Error("PKCS#8私钥不是RSA类型", zap.String("pem_type", block.Type))
		return nil, fmt.Errorf("unsupported PEM type: %s", block.Type)
	}

	if err != nil {
		zap.L().Error("私钥解析失败", zap.String("key_type", fmt.Sprintf("%T", block.Type)))
		return nil, fmt.Errorf("private key parse failed: %w", err)
	}

	// 执行解密（注意：PKCS1v15有已知风险，建议使用OAEP）
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		zap.L().Sugar().Error("解密失败，可能原因：密钥不匹配/密文长度错误", zap.Error(err),
			zap.Int("cipher_len", len(cipherText)), zap.Int("key_size", privateKey.Size()))
		return nil, fmt.Errorf("decryption failed: %w", err)
	}

	return plaintext, nil
}

// GetKeyByteByPath 从文件中读取密钥
//
// 参数
//
//	pemPath (string): 密钥路径
func GetKeyByteByPath(pemPath string) ([]byte, error) {
	fileByte, err := os.ReadFile(pemPath)
	if err != nil {
		zap.L().Sugar().Errorf("read key byte from %s failed %+v", pemPath, err)
		return nil, errors.New("读取密钥文件错误")
	}
	zap.L().Sugar().Infof("\"read key byte from %s success", pemPath)
	return fileByte, nil
}

// HashPassword 对明文密码进行安全哈希处理
//
// 参数:
//
//	-plainPassword: 需要哈希的明文密码（建议长度 8-72 字符）
//	-cost: bcrypt 算法成本（建议值 10-14，默认使用 bcrypt.DefaultCost）
//
// 返回:
//
//	-string: 成功返回 bcrypt 哈希字符串，失败返回空字符串
//	-error: 错误详细信息，调用方必须处理
//
// 安全说明:
//   - bcrypt 自动处理 salt 生成
//   - 明文密码超过 72 字节将被截断
//   - 建议在前端进行密码复杂度校验
func HashPassword(plainPassword string, cost ...int) (string, error) {
	// 输入验证
	if plainPassword == "" {
		return "", errors.New("密码不能为空")
	}

	// 设置动态成本（允许调用方覆盖默认成本）
	bcryptCost := bcrypt.DefaultCost
	if len(cost) > 0 {
		if cost[0] < bcrypt.MinCost || cost[0] > bcrypt.MaxCost {
			return "", fmt.Errorf("invalid bcrypt cost: %d (min %d, max %d)",
				cost[0], bcrypt.MinCost, bcrypt.MaxCost)
		}
		bcryptCost = cost[0]
	}

	// 转换为字节只需一次内存分配
	passwordBytes := []byte(plainPassword)

	// 生成安全哈希
	hashedBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcryptCost)
	if err != nil {
		return "", fmt.Errorf("password hashing failed: %w", err)
	}

	return string(hashedBytes), nil
}

// ComparePassword 安全比较已哈希密码与明文密码是否匹配
//
// 参数
//
//	-hashedPassword: 数据库存储的bcrypt哈希值
//	-plainPassword: 用户提供的明文密码
//
// 返回值:
//
//	-bool: 是否匹配
//
// 安全提示:
//
//	此函数应始终返回bool类型，避免泄露计时信息
func ComparePassword(hashedPassword string, plainPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		return false
	}
	return true
}
