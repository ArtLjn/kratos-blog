/*
@Time : 2024/5/19 下午8:23
@Author : ljn
@File : filter
@Software: GoLand
*/
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Filter struct {
	Domain    []string
	Open      bool
	WhiteList []string
	Key       string
}

var F *Filter

func NewFilter(domain []string, open bool, whiteList []string, key string) *Filter {
	return &Filter{
		Domain:    domain,
		Open:      open,
		WhiteList: whiteList,
		Key:       key,
	}
}

type FilterOption interface {
	Apply() gin.HandlerFunc
}

type FilterOptions struct {
	filters []FilterOption
}

// NewFilterOptions 创建一个新的filterOptions实例
func NewFilterOptions(opts ...Option) *FilterOptions {
	f := &FilterOptions{}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type Option func(o *FilterOptions)

// WithOriginFilter 添加OriginFilter到filterOptions
func WithOriginFilter() Option {
	return func(options *FilterOptions) {
		options.filters = append(options.filters, originFilter{})
	}
}

// originFilter 实现FilterOption接口
type originFilter struct{}

func (f originFilter) Apply() gin.HandlerFunc {
	return func(c *gin.Context) {
		referer := c.Request.Header.Get("Referer")
		if F.Open && len(F.Domain) > 0 {
			for _, v := range F.Domain {
				if strings.HasPrefix(referer, v) {
					c.Header("Access-Control-Allow-Origin", referer)
					c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
					c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
					c.Header("Access-Control-Allow-Credentials", "true")
					if c.Request.Method == "OPTIONS" {
						c.AbortWithStatus(http.StatusNoContent)
						return
					}
					c.Next()
					break
				} else {
					c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
					c.Abort()
					return
				}
			}
		}
	}
}

type ipWhiteFilter struct{}

func WithIpWhiteFilter() Option {
	return func(options *FilterOptions) {
		options.filters = append(options.filters, ipWhiteFilter{})
	}
}

func (f ipWhiteFilter) Apply() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		for _, whiteIp := range F.WhiteList {
			if ip == whiteIp {
				c.Next()
				return
			} else {
				c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
				c.Abort()
				return
			}
		}
	}
}
