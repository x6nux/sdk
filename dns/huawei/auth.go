package huawei

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	dns "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2/region"
)

type Auth struct {
	AccessKey       string `json:"access_token"`
	SecretAccessKey string `json:"secret_access_key"`
	Domain          string `json:"domain"`
	ZoneId          string `json:"zone_id"`
}

func (a *Auth) client() (*dns.DnsClient, error) {
	auth, err := basic.NewCredentialsBuilder().
		WithAk(a.AccessKey).
		WithSk(a.SecretAccessKey).SafeBuild()
	if err != nil {
		return nil, err
	}
	build, err := dns.DnsClientBuilder().
		WithRegion(region.AP_SOUTHEAST_1).
		WithCredential(auth).SafeBuild()
	if err != nil {
		return nil, err
	}
	return dns.NewDnsClient(build), nil
}

func NewAuth(accessKey, secretAccessKey, domain string) (*Auth, error) {
	auth := Auth{
		AccessKey:       accessKey,
		SecretAccessKey: secretAccessKey,
		Domain:          domain,
	}
	zones, err := auth.listZones(domain)
	if err != nil {
		return nil, err
	}
	//判断域名结尾是否为 . 不是则添加
	if auth.Domain[len(auth.Domain)-1:] != "." {
		auth.Domain += "."
	}
	for _, zone := range *zones.Zones {
		if *zone.Name == auth.Domain {
			auth.ZoneId = *zone.Id
			break
		}
	}
	//如果域名不存在，返回错误
	if auth.ZoneId == "" {
		return nil, errors.New("域名不存在")
	}
	return &auth, nil
}
