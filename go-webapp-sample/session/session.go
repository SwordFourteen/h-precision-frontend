package session

import (
	"encoding/json"
	"fmt"
	"gopkg.in/boj/redistore.v1"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/logger"
	"github.com/ybkuroki/go-webapp-sample/model"
)

const (
	// sessionStr represents a string of session key.
	// 定义会话键的常量，表示用于存储会话的键名。
	sessionStr = "GSESSION"
	// Account is the key of account data in the session.
	// 用于存储账户信息的会话键名。
	Account = "Account"
)

// session 结构体封装了存储会话的 store。
type session struct {
	store sessions.Store
}

// Session represents an interface for accessing the session on the application.
// Session 接口定义了会话操作的抽象，包括获取、保存、删除会话以及设置和获取会话中的值。
type Session interface {
	GetStore() sessions.Store
	Get(c echo.Context) *sessions.Session
	Save(c echo.Context) error
	Delete(c echo.Context) error
	SetValue(c echo.Context, key string, value interface{}) error
	GetValue(c echo.Context, key string) string
	SetAccount(c echo.Context, account *model.Account) error
	GetAccount(c echo.Context) *model.Account
}

// NewSession is constructor.
// NewSession 是一个构造函数，根据配置文件创建一个新的会话管理实例。
// 如果 Redis 未启用，使用 CookieStore，否则使用 Redis 作为会话存储。
func NewSession(logger logger.Logger, conf *config.Config) Session {
	// 如果 Redis 未启用，则使用 CookieStore 作为会话存储。
	if !conf.Redis.Enabled {
		logger.GetZapLogger().Infof("use CookieStore for session")
		return &session{sessions.NewCookieStore([]byte("secret"))}
	}

	// 使用 Redis 作为会话存储。
	logger.GetZapLogger().Infof("use redis for session")
	logger.GetZapLogger().Infof("Try redis connection")
	address := fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port) // 构建 Redis 地址
	store, err := redistore.NewRediStore(conf.Redis.ConnectionPoolSize, "tcp", address, "", []byte("secret"))
	if err != nil {
		// 如果 Redis 连接失败，记录错误并退出
		logger.GetZapLogger().Panicf("Failure redis connection, %s", err.Error())
	}
	logger.GetZapLogger().Infof(fmt.Sprintf("Success redis connection, %s", address))
	return &session{store: store}
}

// GetStore returns the session store.
// 返回会话存储对象。
func (s *session) GetStore() sessions.Store {
	return s.store
}

// Get returns a session for the current request.
// 获取当前请求的会话，如果不存在会自动创建新的会话。
func (s *session) Get(c echo.Context) *sessions.Session {
	sess, _ := s.store.Get(c.Request(), sessionStr)
	return sess
}

// Save saves the current session.
// 保存当前会话，将会话数据存储到客户端（例如通过 Cookie 或 Redis）。
func (s *session) Save(c echo.Context) error {
	sess := s.Get(c)
	sess.Options = &sessions.Options{
		Path:     "/",  // 设置会话的作用路径
		HttpOnly: true, // 仅允许通过 HTTP 访问（防止 JavaScript 访问 Cookie）
	}
	return s.saveSession(c, sess)
}

// Delete the current session.
// 删除当前会话，通常通过设置会话的 MaxAge 为 -1 来让会话失效。
func (s *session) Delete(c echo.Context) error {
	sess := s.Get(c)
	sess.Options = &sessions.Options{
		Path:     "/",  // 设置会话的作用路径
		HttpOnly: true, // 仅允许通过 HTTP 访问
		MaxAge:   -1,   // 设置 MaxAge 为 -1 以删除会话
	}
	return s.saveSession(c, sess)
}

// saveSession is a helper method to save the session.
// saveSession 是一个帮助方法，用于保存会话。
func (s *session) saveSession(c echo.Context, sess *sessions.Session) error {
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		// 如果保存会话时发生错误，返回错误信息。
		return fmt.Errorf("error occurred while save session")
	}
	return nil
}

// SetValue sets a key and a value in the session.
// SetValue 将指定的键值对存储到会话中，值会被序列化为 JSON 字符串存储。
func (s *session) SetValue(c echo.Context, key string, value interface{}) error {
	sess := s.Get(c)
	bytes, err := json.Marshal(value)
	if err != nil {
		// 如果序列化为 JSON 时发生错误，返回错误信息。
		return fmt.Errorf("json marshal error while set value in session")
	}
	sess.Values[key] = string(bytes) // 将序列化后的 JSON 字符串存入会话
	return nil
}

// GetValue returns the value of a session by key.
// GetValue 根据指定的键从会话中获取值，如果存在则返回对应的字符串。
func (s *session) GetValue(c echo.Context, key string) string {
	sess := s.Get(c)
	if sess != nil {
		if v, ok := sess.Values[key]; ok {
			data, result := v.(string)
			if result && data != "null" {
				return data // 返回值，如果存在并且不是 "null"
			}
		}
	}
	return "" // 如果没有找到对应的值，则返回空字符串
}

// SetAccount sets the account information into the session.
// SetAccount 将用户账户信息存入会话中。
func (s *session) SetAccount(c echo.Context, account *model.Account) error {
	return s.SetValue(c, Account, account)
}

// GetAccount returns the account information from the session.
// GetAccount 从会话中获取账户信息，如果存在则反序列化为 Account 结构体并返回。
func (s *session) GetAccount(c echo.Context) *model.Account {
	if v := s.GetValue(c, Account); v != "" {
		a := &model.Account{}
		_ = json.Unmarshal([]byte(v), a) // 将 JSON 字符串反序列化为 Account 结构体
		return a
	}
	return nil // 如果未找到账户信息，则返回 nil
}
