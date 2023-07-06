package logging

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Metadata struct {
	Context string
	Payload interface{}
	Data    Data
	User    User
}

type User struct {
	Id    string
	Name  string
	Email string
}

type TokenResponse struct {
	CompanyID int64    `json:"company_id"`
	Name      string   `json:"name"`
	Username  string   `json:"username"`
	Emails    []Emails `json:"emails"`
	Type      string   `json:"typeDescription"`
}

type Emails struct {
	Email string `json:"email"`
	Main  bool   `json:"is_main"`
}

type Data map[string]map[string]interface{}

func (m Metadata) ToZapMetadata() []interface{} {
	meta := []interface{}{}
	if len(m.Context) > 0 {
		meta = append(meta, "Context")
		meta = append(meta, m.Context)
	}
	if m.Payload != nil {
		meta = append(meta, "Payload")
		meta = append(meta, m.Payload)
	}
	if len(m.Data) > 0 {
		meta = append(meta, "Data")
		meta = append(meta, m.Data)
	}

	return meta
}

func GinContextToMetadata(c *gin.Context) Metadata {
	metadata := Metadata{}

	// Context
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	if raw != "" {
		path = path + "?" + raw
	}
	metadata.Context = path

	// Payload
	if v, ok := c.Get(gin.BodyBytesKey); ok {
		if s, ok := v.([]byte); ok {
			metadata.Payload = string(s)
		}
	}

	// User
	if t, ok := c.Get("user"); ok {
		if token, ok := t.(TokenResponse); ok {
			metadata.User.Id = strconv.FormatInt(token.CompanyID, 10)
			metadata.User.Name = token.Name
			if len(token.Emails) > 0 {
				metadata.User.Email = token.Emails[0].Email
			}
		}
	}

	return metadata
}
