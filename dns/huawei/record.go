package huawei

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2/model"
	"strings"
)

func (a *Auth) CreateRecord(r Record) (model.CreateRecordSetResponse, error) {
	request := &model.CreateRecordSetRequest{
		ZoneId: a.ZoneId,
	}
	client, err := a.client()
	if err != nil {
		return model.CreateRecordSetResponse{}, err
	}
	request.Body = r.createRecord()
	if !strings.HasSuffix(request.Body.Name, a.Domain) {
		// name 必须以域名结尾
		//判断当前name是否使用 . 作为结尾
		if request.Body.Name[len(request.Body.Name)-1:] != "." {
			request.Body.Name += "."
		}
		request.Body.Name += a.Domain
	}
	response, err := client.CreateRecordSet(request)
	if err != nil {
		return model.CreateRecordSetResponse{}, err
	}
	return *response, err
}

func (a *Auth) RecordList(Type ...string) (model.ListRecordSetsByZoneResponse, error) {
	client, err := a.client()
	if err != nil {
		return model.ListRecordSetsByZoneResponse{}, err
	}
	request := &model.ListRecordSetsByZoneRequest{}
	request.ZoneId = a.ZoneId
	if len(Type) == 1 {
		request.Type = &Type[0]
	}
	response, err := client.ListRecordSetsByZone(request)
	if err != nil {
		return model.ListRecordSetsByZoneResponse{}, err
	}
	return *response, nil
}

func (a *Auth) DeleteRecord(id string) (model.DeleteRecordSetResponse, error) {
	client, err := a.client()
	if err != nil {
		return model.DeleteRecordSetResponse{}, err
	}
	request := &model.DeleteRecordSetRequest{
		ZoneId:      a.ZoneId,
		RecordsetId: id,
	}
	response, err := client.DeleteRecordSet(request)
	if err != nil {
		return model.DeleteRecordSetResponse{}, err
	}
	return *response, nil
}

func (a *Auth) FindRecord(id string) (model.ShowRecordSetResponse, error) {
	client, err := a.client()
	if err != nil {
		return model.ShowRecordSetResponse{}, err
	}
	request := &model.ShowRecordSetRequest{
		ZoneId:      a.ZoneId,
		RecordsetId: id,
	}
	response, err := client.ShowRecordSet(request)
	if err != nil {
		return model.ShowRecordSetResponse{}, err
	}
	return *response, nil
}

func (a *Auth) UpdateRecord(id string, r Record) (model.UpdateRecordSetResponse, error) {
	client, err := a.client()
	if err != nil {
		return model.UpdateRecordSetResponse{}, err
	}
	request := &model.UpdateRecordSetRequest{
		ZoneId:      a.ZoneId,
		RecordsetId: id,
	}
	request.Body = r.updateRecord()
	name := r.Name
	if !strings.HasSuffix(name, a.Domain) {
		// name 必须以域名结尾
		//判断当前name是否使用 . 作为结尾
		if name[len(name)-1:] != "." {
			name += "."
		}
		name += a.Domain
	}
	request.Body.Name = &name
	response, err := client.UpdateRecordSet(request)
	if err != nil {
		return model.UpdateRecordSetResponse{}, err
	}
	return *response, nil
}
