package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

// GetOS 获取操作系统
func GetOS(c *gin.Context) string {
	userAgentStr := c.Request.Header.Get("User-Agent")
	if userAgentStr == "" {
		return "Unknown"
	}

	ua := user_agent.New(userAgentStr)
	return normalizeOS(ua.OS())
}

// GetBrowser 获取浏览器信息
func GetBrowser(c *gin.Context) string {
	userAgentStr := c.Request.Header.Get("User-Agent")
	if userAgentStr == "" {
		return "Unknown"
	}

	ua := user_agent.New(userAgentStr)
	browser, _ := ua.Browser()
	return normalizeBrowser(browser)
}

// normalizeOS 标准化操作系统名称
func normalizeOS(os string) string {
	if os == "" {
		return "Unknown"
	}

	os = strings.ToLower(os)
	switch {
	case strings.Contains(os, "windows"):
		return "Windows"
	case strings.Contains(os, "mac"):
		return "macOS"
	case strings.Contains(os, "linux"):
		return "Linux"
	case strings.Contains(os, "android"):
		return "Android"
	case strings.Contains(os, "ios") || strings.Contains(os, "iphone"):
		return "iOS"
	case strings.Contains(os, "ipad"):
		return "iPadOS"
	default:
		return os
	}
}

// normalizeBrowser 标准化浏览器名称
func normalizeBrowser(browser string) string {
	if browser == "" {
		return "Unknown"
	}

	browser = strings.ToLower(browser)
	switch {
	case strings.Contains(browser, "chrome"):
		return "Chrome"
	case strings.Contains(browser, "firefox"):
		return "Firefox"
	case strings.Contains(browser, "safari"):
		return "Safari"
	case strings.Contains(browser, "edge"):
		return "Edge"
	case strings.Contains(browser, "opera"):
		return "Opera"
	case strings.Contains(browser, "ie") || strings.Contains(browser, "internet explorer"):
		return "Internet Explorer"
	case strings.Contains(browser, "postman"):
		return "Postman"
	case strings.Contains(browser, "curl"):
		return "cURL"
	default:
		return browser
	}
}
