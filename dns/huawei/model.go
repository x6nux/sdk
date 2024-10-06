package huawei

import "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2/model"

const (
	RecordTypeA     = "A"     // ipv4
	RecordTypeAAAA  = "AAAA"  // ipv6
	RecordTypeCNAME = "CNAME" // 域名
	RecordTypeMX    = "MX"    // 邮件
	RecordTypeTXT   = "TXT"   // 文本
	RecordTypeSRV   = "SRV"   // 服务
	RecordTypeNS    = "NS"    // 名称服务器
	RecordTypeCAA   = "CAA"   // 证书
)

type Record struct {

	// 域名，后缀需以zone name结束且为FQDN（即以“.”号结束的完整主机名）。
	Name string `json:"name"`

	// 可选配置，对域名的描述。  长度不超过255个字符。  默认值为空。
	Description string `json:"description,omitempty"`

	// Record Set的类型。  公网域名场景的记录类型: A、AAAA、MX、CNAME、TXT、NS、SRV、CAA。  内网域名场景的记录类型: A、AAAA、MX、CNAME、TXT、SRV。
	Type string `json:"type"`

	// 资源状态。
	Status string `json:"status,omitempty"`

	// 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。 如果您的服务地址经常更换，建议TTL值设置相对小些，反之，建议设置相对大些。
	Ttl int32 `json:"ttl,omitempty"`

	// 解析记录的值。不同类型解析记录对应的值的规则不同。
	Records []string `json:"records"`

	// 资源标签。
	Tags []Tag `json:"tags,omitempty"`
}

type Tag struct {

	// 键。最大长度36个unicode字符。 key不能为空。不能包含“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”，且首尾字符不能为空格。
	Key string `json:"key"`

	// 值。每个值最大长度43个unicode字符，可以为空字符串。 不能包含“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”，且首尾字符不能为空格。
	Value *string `json:"value,omitempty"`
}

func (r *Record) createRecord() *model.CreateRecordSetRequestBody {
	des := r.Description
	status := r.Status
	ttl := r.Ttl
	var tags []model.Tag
	for _, i := range r.Tags {
		tags = append(tags, model.Tag{
			Key:   i.Key,
			Value: i.Value,
		})
	}
	return &model.CreateRecordSetRequestBody{
		Name:        r.Name,
		Description: &des,
		Type:        r.Type,
		Status:      &status,
		Ttl:         &ttl,
		Records:     r.Records,
		Tags:        &tags,
	}
}

func (r *Record) updateRecord() *model.UpdateRecordSetReq {
	return &model.UpdateRecordSetReq{
		Name:        &r.Name,
		Description: &r.Description,
		Type:        &r.Type,
		Ttl:         &r.Ttl,
		Records:     &r.Records,
	}
}
