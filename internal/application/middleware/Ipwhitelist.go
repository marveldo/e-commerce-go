package middleware

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IpwhitelistMiddleware(allowed_CIDRs []string) gin.HandlerFunc {
	var allowed_ips []*net.IPNet
	for _, cidr := range allowed_CIDRs {
		_, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			ip := net.ParseIP(cidr)
			if ip == nil {
				// Invalid input – panic during setup (or log and skip)
				panic("invalid IP or CIDR: " + cidr)
			}
			if ip.To4() != nil {
				_, ipnet, _ = net.ParseCIDR(cidr + "/32")
			} else {
				_, ipnet, _ = net.ParseCIDR(cidr + "/128")
			}

		}
		allowed_ips = append(allowed_ips, ipnet)
	}
	return func(c *gin.Context) {
		ipstr := c.ClientIP()
		ip := net.ParseIP(ipstr)
		if ip == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, 
				gin.H{
					"status" : http.StatusForbidden,
					"error": "invalid client IP",
				})
			return
		}
		for _ , ipnet := range allowed_ips {
			if ipnet.Contains(ip) {
				c.Next()
				return
			}
		}
		 c.AbortWithStatusJSON(http.StatusForbidden, 
			gin.H{
			"status" : http.StatusForbidden,
			"error": "access denied",
		})

	}
}
