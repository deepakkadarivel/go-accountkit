package context

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

const RequestKey = "request"

type Context interface {
	DBConnection() *gorm.DB
}

type AppContext struct {
	dbConnection *gorm.DB
}

func NewContext(dbConnection *gorm.DB) *AppContext {
	return &AppContext{
		dbConnection: dbConnection,
	}
}

func (appContext *AppContext) DBConnection() *gorm.DB {
	return appContext.dbConnection
}

type RequestContext struct {
	appContext *AppContext
}

func NewRequestContext(appContext *AppContext) *RequestContext {
	return &RequestContext{
		appContext: appContext,
	}
}

func ParseRequestContext(ctx context.Context) (Context, error) {
	return ctx.Value(RequestKey).(Context), nil
}

func (r *RequestContext) DBConnection() *gorm.DB {
	return r.appContext.dbConnection
}