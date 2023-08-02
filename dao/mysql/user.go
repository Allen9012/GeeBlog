package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/Allen9012/gee_blog/common"
	"github.com/Allen9012/gee_blog/model"
	"go.uber.org/zap"
)

const secret = "https://github.com/Allen9012"

type userDAO struct {
}

var User = new(userDAO)

func (userDAO) Insert(u *model.User) (err error) {
	u.Password = encryptPassword(u.Password)
	err = common.GEE_DB.Create(u).Error
	if err != nil {
		zap.L().Warn("[dao mysql user] insert user error ", zap.Error(err))
	}
	return
}

func (userDAO) Login(account, password string) (*model.UserVO, error) {
	u := new(model.UserVO)
	err := common.GEE_DB.First(&model.User{}, "account = ? and password = ?", account, encryptPassword(password)).Scan(u).Error
	if err != nil {
		zap.L().Warn("[dao mysql user] user login error ", zap.Error(err))
		return nil, err
	}
	return u, err
}

func (userDAO) QueryUserByAccount(account string) (*model.User, error) {
	u := new(model.User)
	err := common.GEE_DB.First(u, "account = ?", account).Error
	if err != nil {
		zap.L().Warn("[dao mysql user] query user by account error ", zap.Error(err))
		return nil, err
	}
	return u, nil
}

func (userDAO) QueryUserVOByAccount(account string) (*model.UserVO, error) {
	u := new(model.UserVO)
	err := common.GEE_DB.First(&model.User{}, "account = ?", account).Scan(u).Error
	if err != nil {
		zap.L().Warn("[dao mysql user] query user vo by account error ", zap.Error(err))
		return nil, err
	}
	return u, nil
}

func (userDAO) QueryUserByUserId(id int64) (*model.User, error) {
	u := new(model.User)
	err := common.GEE_DB.First(u, "user_id = ?", id).Error
	if err != nil {
		zap.L().Warn("[dao mysql user] query user by userId error ", zap.Error(err))
		return nil, err
	}
	return u, nil
}

func (userDAO) QueryUserVOByUserId(id int64) (*model.UserVO, error) {
	u := new(model.UserVO)
	err := common.GEE_DB.First(&model.User{}, "user_id = ?", id).Scan(u).Error
	if err != nil {
		zap.L().Warn("[dao mysql user] query user vo by userId error ", zap.Error(err))
		return nil, err
	}
	return u, nil
}

func (userDAO) QueryUserList(params *model.ListParams) ([]*model.User, error) {
	var u []*model.User
	err := common.GEE_DB.Unscoped().Limit(params.Size).Offset(params.Page - 1).Find(&u).Error
	if err != nil {
		zap.L().Warn("[dao mysql user] query user list error ", zap.Error(err))
		return nil, err
	}
	return u, nil
}

func (userDAO) QueryUserVOList(params *model.ListParams) ([]*model.UserVO, error) {
	var u []*model.UserVO
	err := common.GEE_DB.Limit(params.Size).Offset(params.Page - 1).Model(&model.User{}).Scan(&u).Error
	if err != nil {
		zap.L().Warn("[dao mysql user] query user vo error ", zap.Error(err))
		return nil, err
	}
	return u, nil
}

func (u userDAO) CheckAccountExist(account string) bool {
	if _, err := u.QueryUserByAccount(account); err != nil {
		return false
	}
	return true
}

func (userDAO) UpdateUserBySelf(u *model.UserDTOUpdateBySelf) (err error) {
	err = common.GEE_DB.Take(&model.User{}, "user_id = ?", u.UserId).Updates(model.User{
		Account:     u.Account,
		Password:    encryptPassword(u.Password),
		Gender:      u.Gender,
		Phone:       u.Phone,
		Email:       u.Email,
		Description: u.Description,
	}).Error
	return err
}

func (userDAO) UpdateUserByAdmin(u *model.UserDTOUpdateByAdmin) (err error) {
	fmt.Println(*u)
	err = common.GEE_DB.Take(&model.User{}, "user_id = ?", u.UserId).Updates(model.User{
		Account:     u.Account,
		Password:    encryptPassword(u.Password),
		Gender:      u.Gender,
		Phone:       u.Phone,
		Email:       u.Email,
		Description: u.Description,
		UserRole:    u.UserRole,
	}).Error
	return
}

func (userDAO) DeleteUserByUserId(userId int64) (err error) {
	err = common.GEE_DB.Delete(&model.User{}, "user_id = ?", userId).Error
	return
}

// 密码加密
func encryptPassword(originPassword string) string {
	hash := md5.New()
	hash.Write([]byte(secret))
	hash.Write([]byte(originPassword))
	encryptString := hex.EncodeToString(hash.Sum(nil))
	return encryptString
}
