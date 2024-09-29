package huawei

import dm "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2/model"

func (a *Auth) listZones(domain ...string) (dm.ListPublicZonesResponse, error) {
	client, err := a.client()
	if err != nil {
		return dm.ListPublicZonesResponse{}, err
	}
	req := dm.ListPublicZonesRequest{}
	if len(domain) > 0 {
		req.Name = &domain[0]
	}
	resp, err := client.ListPublicZones(&req)
	if err != nil {
		return dm.ListPublicZonesResponse{}, err
	}
	return *resp, nil
}

func (a *Auth) GetZones() (dm.ShowPublicZoneResponse, error) {
	client, err := a.client()
	if err != nil {
		return dm.ShowPublicZoneResponse{}, err
	}
	req := dm.ShowPublicZoneRequest{
		ZoneId: a.ZoneId,
	}
	resp, err := client.ShowPublicZone(&req)
	if err != nil {
		return dm.ShowPublicZoneResponse{}, err
	}
	return *resp, nil
}
