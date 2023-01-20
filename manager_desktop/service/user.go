package service

import (
	"context"
	"google.golang.org/grpc"
	"simple_gateway/dto"
	"simple_gateway/global"
	"simple_gateway/global/form"
	"simple_gateway/model"
	"simple_gateway/protoc"
)

func UserList(req form.AppListReq) (*model.UserList, error) {
	info := &model.User{}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	all, total := info.PageSelect(req)
	offset := (req.PageNum - 1) * req.PageSize
	limit := offset + req.PageSize
	if limit > total {
		limit = total
	}
	infoList := make([]*model.UserListItem, 0)
	list := &model.UserList{Total: total}
	for _, v := range all[offset:limit] {
		dial, err := grpc.Dial(global.DebugFullConfig.GRPCServer.Host, grpc.WithInsecure())
		client := protoc.NewFlowCountClient(dial)
		rsp, err := client.GetUserFlowCount(context.Background(), &protoc.FlowCountRequest{ServiceName: v.AppID})
		if err != nil {
			return nil, err
		}
		item := &model.UserListItem{
			Id:     int(v.ID),
			AppId:  v.AppID,
			Name:   v.Name,
			Secret: v.Secret,
			QPS:    int(rsp.Qps),
			QPD:    int(rsp.Qpd),
		}
		infoList = append(infoList, item)
	}
	list.List = infoList
	return list, nil
}

func AddUser(req form.AddUserReq) error {
	u := model.User{}
	return u.Save(req)
}

func UserStat(req form.UserStatReq) (*dto.UserStatOutput, error) {
	u := &model.User{}
	err := u.FindById(req.Id)
	if err != nil {
		return nil, err
	}
	output := &dto.UserStatOutput{
		Name: u.Name,
	}
	dial, err := grpc.Dial(global.DebugFullConfig.GRPCServer.Host, grpc.WithInsecure())
	client := protoc.NewFlowCountClient(dial)
	rsp, err := client.GetUserFlowCount(context.Background(), &protoc.FlowCountRequest{ServiceName: u.AppID})
	if err != nil {
		return nil, err
	}

	output.Today = rsp.TodayCount
	output.Yesterday = rsp.YesterdayCount
	return output, nil
}

func UpdateUser(req form.UpdateUserReq) error {
	u := model.User{}
	return u.Update(req)
}

func DeleteUser(req form.DeleteUserReq) error {
	u := model.User{}
	findU := &model.User{}
	err := findU.FindById(req.Id)
	if err != nil {
		return err
	}
	err = u.Delete(req)
	if err != nil {
		return err
	}
	model.DeleteUserSync(findU.AppID)
	return nil
}

func AppDetail(req form.UserDetailReq) (*model.User, error) {
	u := &model.User{}
	err := u.FindById(req.Id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
