package controllers

import (
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/context"
	"micro_test/models"
	"micro_test/proto"
)

//转换详情
func TransPddSessions(info *models.PddSessions) *proto.PddSessions {
	out := new(proto.PddSessions)
	out.Id = int32(info.ID)
	out.TaokeID = int32(info.TaokeID)
	out.ScreenName = info.ScreenName
	out.OpenId = info.OpenId
	out.Token = info.Token
	out.RefreshToken = info.RefreshToken
	out.IsDefault = int32(info.IsDefault)
	out.ExpiredAt = TransTime(info.ExpiredAt)
	out.RefreshExpiredAt = TransTime(info.RefreshExpiredAt)
	out.CreatedAt = TransTime(info.CreatedAt)
	out.UpdatedAt = TransTime(info.UpdatedAt)
	return out
}

//转换详情
func UnTransPddSessions(info *proto.PddSessions) *models.PddSessions {
	out := new(models.PddSessions)
	out.ID = uint(info.Id)
	out.TaokeID = uint(info.TaokeID)
	out.ScreenName = info.ScreenName
	out.OpenId = info.OpenId
	out.Token = info.Token
	out.RefreshToken = info.RefreshToken
	out.IsDefault = uint8(info.IsDefault)
	out.ExpiredAt = UnTransTime(info.ExpiredAt)
	out.RefreshExpiredAt = UnTransTime(info.RefreshExpiredAt)
	out.CreatedAt = UnTransTime(info.CreatedAt)
	out.UpdatedAt = UnTransTime(info.UpdatedAt)
	return out
}

//添加
func (s *Server) AddPddSessions(ctx context.Context, in *proto.RequestAddPddSessions, out *proto.ResponseAddPddSessions) error {
	logs.Info("Server.AddPddSessions")
	info := UnTransPddSessions(in.PddSessions)
	var err error
	info, err = models.CreatePddSessions(info)
	if err != nil {
		return  err
	}
	out.PddSessions = TransPddSessions(info)
	return nil
}

//更新
func (s *Server) UpdatePddSessions(ctx context.Context, in *proto.RequestUpdatePddSessions, out *proto.ResponseUpdatePddSessions) error {
	logs.Info("Server.UpdatePddSessions")
	info := UnTransPddSessions(in.PddSessions)
	var err error
	info, err = models.UpdatePddSessions(info)
	if err != nil {
		return err
	}
	out.PddSessions = TransPddSessions(info)
	return nil
}

//删除
func (s *Server) DeletePddSessions(ctx context.Context, in *proto.RequestDeletePddSessions, out *proto.ResponseDeletePddSessions) error {
	logs.Info("Server.DeletePddSessions")
	err := models.DeletePddSessions(uint(in.Id))
	if err != nil {
		return  err
	}
	return nil
}

//详情
func (s *Server) GetPddSessionsInfo(ctx context.Context, in *proto.RequestGetPddSessionsInfo, out *proto.ResponseGetPddSessionsInfo) error {
	logs.Info("Server.GetPddSessionsInfo")
	info := models.GetPddSessionsInfo(uint(in.Id))
	returnInfo := TransPddSessions(info)
	out.PddSessions = returnInfo
	return nil
}

//列表
func (s *Server) GetPddSessionsList(ctx context.Context, in *proto.RequestGetPddSessionsList, out *proto.ResponseGetPddSessionsList) error{
	logs.Info("Server.GetPddSessionsList")
	list, total := models.GetPddSessionsList(int(in.Page), int(in.PerPage))
	returnList := []*proto.PddSessions{}
	for _, val := range list {
		returnList = append(returnList, TransPddSessions(val))
	}
	out.PddSessions = returnList
	out.Total = int32(total)
	return nil
}
