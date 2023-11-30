package clients

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Mitmadhu/commons/dto/response"
	"github.com/Mitmadhu/commons/handle_http"
	mysqlDto "github.com/Mitmadhu/mysqlDB/dto"
)

type MysqlClientImpl struct {
}

func (m *MysqlClientImpl) Login(url string, method string, req mysqlDto.ValidateUserRequest) (*mysqlDto.ValidateUserResponse, error) {

	byteResp, err := handle_http.Call(url+"/user-exists", http.MethodPost, mysqlDto.ValidateUserRequest{
		BaseRequest: mysqlDto.BaseRequest{MsgID: req.MsgID},
		Username:    req.Username,
		Password:    req.Password,
	})
	if err != nil {
		return nil, err
	}
	genericResp := &response.SuccessResponse{}
	err = json.Unmarshal(byteResp, genericResp)
	if err != nil {
		return nil, err
	}
	loginResp, ok := genericResp.Response.(mysqlDto.ValidateUserResponse)
	if !ok {
		return nil, errors.New("fail to convert into login respose type")
	}
	return &loginResp, nil
}
