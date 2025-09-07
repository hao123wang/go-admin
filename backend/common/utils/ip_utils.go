package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

// IPInfo 地理位置信息
type IPInfo struct {
	Code int    `json:"code"`
	Addr string `json:"addr"`
	IP   string `json:"ip"`
}

// GetLocalIP 获取本机IP地址
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		ip := ipNet.IP
		// 跳过回环地址和非全局单播地址
		if ip.IsLoopback() || !ip.IsGlobalUnicast() {
			continue
		}

		// 优先返回IPv4地址
		if ip4 := ip.To4(); ip4 != nil {
			return ip4.String(), nil
		}
	}

	return "", fmt.Errorf("未找到有效IP地址")
}

// GetRealAddressByIP 获取IP地址的真实地理位置
func GetRealAddressByIP(ipStr string) string {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "无效IP地址"
	}

	if isLocalIP(ip) {
		return "内网地址"
	}

	if isLANIP(ip) {
		return "局域网"
	}

	return getIPLocation(ipStr)
}

// isLocalIP 判断是否为本地IP
func isLocalIP(ip net.IP) bool {
	return ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast()
}

// isLANIP 判断是否为局域网IP
func isLANIP(ip net.IP) bool {
	if ip4 := ip.To4(); ip4 != nil {
		// IPv4 局域网段
		switch {
		case ip4[0] == 10:
			return true
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return true
		case ip4[0] == 192 && ip4[1] == 168:
			return true
		case ip4[0] == 169 && ip4[1] == 254: // APIPA
			return true
		}
	} else if ip6 := ip.To16(); ip6 != nil {
		// IPv6 局域网段
		switch {
		case ip[0] == 0xfe && ip[1] == 0x80: // Link-local
			return true
		case ip[0] == 0xfc || ip[0] == 0xfd: // Unique Local Address
			return true
		}
	}
	return false
}

// getIPLocation 获取IP地理位置
func getIPLocation(ip string) string {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	url := fmt.Sprintf("https://whois.pconline.com.cn/ipJson.jsp?json=true&ip=%s", ip)

	resp, err := client.Get(url)
	if err != nil {
		return "网络请求失败"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "服务不可用"
	}

	var ipInfo IPInfo
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&ipInfo); err != nil {
		return "解析失败"
	}

	if ipInfo.Code == 0 && ipInfo.Addr != "" {
		return ipInfo.Addr
	}

	return "未知地址"
}
