package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tinyhole/user-srv-passport/dal/db"
	"github.com/tinyhole/user-srv-passport/global"
	"github.com/tinyhole/user-srv-passport/idl/base"
	"github.com/tinyhole/user-srv-passport/idl/platform/id/srv/snowflake"
	srv "github.com/tinyhole/user-srv-passport/idl/platform/user/srv/passport"
	"github.com/tinyhole/user-srv-passport/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
)

type Handler struct {}

func (h *Handler) SignUp(ctx context.Context, req *srv.SignUpReq, rsp *srv.SignUpRsp) error {
	rsp.BaseRsp = &srv.BaseRsp{
		Code: srv.Code_Failed,
	}
	mobile, err := util.Mobile(req.Mobile)
	if err != nil {
		// 这种错误是bug类型错误，不应该出现
		logrus.Debugf("bad mobile, err:%v", err)
		rsp.BaseRsp.Msg = err.Error()
		return nil
	}

	// 检验手机号是否已被注册


	// TODO 后期再验证短信验证码
	//if !format.IsDigit(req.Code) {
	//	logrus.Debugf("bad code %v", req.Code)
	//	return errors.New("bad code")
	//}

	if err = util.Password(req.Password); err != nil {
		// bug 错误
		logrus.Debugf("bad password. %v", err)
		rsp.BaseRsp.Msg = err.Error()
		return nil
	}

	// 生成密码
	passwd, salt := util.Make([]byte(req.Password))

	cl := snowflake.NewSnowFlakeService("platform.id.srv.snowflake", global.PassportSvc.Client())
	idRsp, err := cl.GetID(ctx, &base.Empty{})
	if err != nil {
		// 网络或框架错误，通知框架层错误，做出重试
		logrus.Errorf("platform.id.srv.snowflake GetID error: %v", err)
		return err
	}

	ua := db.UserPassport{
		UID:		  idRsp.ID,
		Name:         req.Name,
		Mobile:       mobile,
		Email:        req.Email,
		Password:     string(passwd),
		Salt:         string(salt),
		WeChatID: 	  "",
	}
	err = db.UserPassportOp.Insert(ctx, &ua)
	if err != nil {
		// uid 重复插入失败错误，如何处理
		logrus.Error("Failed to insert user auth", err)
		return err
	}

	rsp.UserID = idRsp.ID
	rsp.BaseRsp.Code = srv.Code_Ok
	return nil
}


type WeChatOpenCode2SessionResp struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// 微信登陆
func (h *Handler) WeChatSignIn(ctx context.Context, req *srv.WeChatSignInReq, rsp *srv.WeChatSignInRsp) error {
	rsp.BaseRsp = &srv.BaseRsp{
		Code: srv.Code_Ok,
	}

	if len(req.AppID) == 0 || len(req.Code) == 0 {
		logrus.Warnf("invalid parameter")
		rsp.BaseRsp.Code = srv.Code_Failed
		rsp.BaseRsp.Msg = "invalid AppID or invalid code"
		return nil
	}

	// 根据appid获取私钥
	secret, ok := util.DefaultConfig.WeChatOpenConfig.Secrets[req.AppID]
	if !ok {
		logrus.Warnf("invalid appid")
		rsp.BaseRsp.Code = srv.Code_Failed
		rsp.BaseRsp.Msg = "invalid AppID"
		return nil
	}

	// 向微信开放接口请求，登陆校验
	reqStr := fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		util.DefaultConfig.WeChatOpenConfig.URL, req.AppID, secret, req.Code)
	resp1, err := http.Get(reqStr)
	if err != nil {
		logrus.Warnf("get %s error. %s", reqStr, err)
		return err
	}
	defer resp1.Body.Close()

	body, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		logrus.Warnf("read body from resp error. %s", err)
		return err
	}

	jsonResp := &WeChatOpenCode2SessionResp{}
	err = json.Unmarshal(body, jsonResp)

	if jsonResp.ErrCode != 0 {
		logrus.Warnf("request wechat open api return errcode:%d", jsonResp.ErrCode)
		rsp.BaseRsp.Code = srv.Code_Failed
		rsp.BaseRsp.Msg = fmt.Sprintf("request wechat open api return errcode:%d", jsonResp.ErrCode)
		return nil
	}

	logrus.Debugf("jsonResp:%+v", jsonResp)

	// 检查该openid是否已经注册过
	u, err := db.UserPassportOp.Get(ctx, bson.D{{"wechat_id",jsonResp.OpenID}})
	if err != nil {
		// 没有找到
		if err == mongo.ErrNoDocuments {
			// 请求一个uid
			cl := snowflake.NewSnowFlakeService("platform.id.srv.snowflake", global.PassportSvc.Client())
			idRsp, err := cl.GetID(ctx, &base.Empty{})
			if err != nil {
				logrus.Errorf("platform.id.srv.snowflake GetID error: %v", err)
				return err
			}

			// TODO: 自动生成用户名称
			u = &db.UserPassport{
				UID:       idRsp.ID,
				WeChatID:  jsonResp.OpenID,
			}
			// 插入一条用户通行证信息

			err = db.UserPassportOp.Insert(ctx, u)
			if err != nil {
				// uid 重复插入失败错误，如何处理
				logrus.Error("Failed to insert user auth", err)
				return err
			}
		} else {
			logrus.Warnf("mongo error:%v", err)
			return err
		}
	}

	rsp.UserID = u.UID

	return nil
}
//
//func (h *Handler) MobileSignIn(ctx context.Context, req *srv.MobileSignInReq, rsp *srv.MobileSignInResp) error {
//
//	return nil
//}
//
//func (h *Handler) EmailSignIn(ctx context.Context, req *srv.EmailSignInReq, rsp *srv.EmailSignInResp) error {
//
//	return nil
//}
//
//func (h *Handler) UserNameSignIn(ctx context.Context, req *srv.UserNameSignInReq, rsp *srv.UserNameSignInResp) error {
//
//	return nil
//}
//
//func (h *Handler) TokenSignIn(ctx context.Context, req *srv.TokenSignInReq, rsp *srv.TokenSignInResp) error {
//
//	return nil
//}
//
//func (h *Handler) SignOut(ctx context.Context, req *srv.SignOutReq, rsp *srv.SignOutResp) error {
//
//	return nil
//}
//
//func (h *Handler) ChangePassword(ctx context.Context, req *srv.ChangePasswordReq, rsp *srv.ChangePasswordResp) error {
//
//	return nil
//}
//
//
//func (h *Handler) ResetPassword(ctx context.Context, req *srv.ResetPasswordReq, rsp *srv.ResetPasswordResp) error {
//
//	return nil
//}
//
//func (h *Handler) BindMobile(ctx context.Context, req *srv.BindMobileReq, rsp *srv.BindMobileResp) error {
//
//	return nil
//}
//
//func (h *Handler) UnbindMobile(ctx context.Context, req *srv.UnbindMobileReq, rsp *srv.UnbindMobileResp) error {
//
//	return nil
//}
