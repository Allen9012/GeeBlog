package mysql

import (
	"github.com/Allen9012/gee_blog/common"
	"github.com/Allen9012/gee_blog/model"
	"go.uber.org/zap"
)

type postDAO struct {
}

var Post = new(postDAO)

func (postDAO) Insert(p *model.Post) (err error) {
	err = common.GEE_DB.Create(p).Error
	if err != nil {
		zap.L().Error("[dao mysql post] insert post error ", zap.Error(err))
	}
	return
}

func (postDAO) QueryPostByPostId(id int64) (*model.Post, error) {
	u := new(model.Post)
	err := common.GEE_DB.First(u, "post_id = ?", id).Error
	if err != nil {
		zap.L().Warn("[dao mysql post] query post by postId error ", zap.Error(err))
		return nil, err
	}
	return u, nil
}

func (postDAO) QueryPostVOByPostId(id int64) (*model.PostVO, error) {
	u := new(model.PostVO)
	err := common.GEE_DB.First(&model.Post{}, "post_id = ?", id).Scan(u).Error
	if err != nil {
		zap.L().Warn("[dao mysql post] query post vo by postId error ", zap.Error(err))
		return nil, err
	}
	return u, nil
}

func (postDAO) QueryPostList(params *model.ListParams) ([]*model.Post, error) {
	var u []*model.Post
	err := common.GEE_DB.Unscoped().Limit(params.Size).Offset(params.Page - 1).Find(&u).Error
	if err != nil {
		zap.L().Warn("[dao mysql post] query post list error ", zap.Error(err))
		return nil, err
	}
	return u, nil
}

func (postDAO) QueryPostVOList(params *model.ListParams) ([]*model.PostVO, error) {
	var u []*model.PostVO
	err := common.GEE_DB.Limit(params.Size).Offset(params.Page - 1).Model(&model.Post{}).Scan(&u).Error
	if err != nil {
		zap.L().Warn("[dao mysql post] query post vo error ", zap.Error(err))
		return nil, err
	}
	return u, nil
}

func (postDAO) QueryPostVOListByUserId(params *model.ListParams, userId int64) ([]*model.PostVO, error) {
	var u []*model.PostVO
	err := common.GEE_DB.Limit(params.Size).Offset(params.Page-1).Model(&model.Post{}).Where("user_id = ?", userId).Scan(&u).Error
	if err != nil {
		zap.L().Warn("[dao mysql post] query post vo error ", zap.Error(err))
		return nil, err
	}
	return u, nil
}

func (postDAO) UpdatePostBySelf(u *model.PostDTOUpdateBySelf) (err error) {
	err = common.GEE_DB.Take(&model.Post{}, "post_id = ?", u.PostId).Updates(model.Post{
		Title:   u.Title,
		Content: u.Content,
	}).Error
	return err
}

func (postDAO) UpdatePostByAdmin(u *model.PostDTOUpdateByAdmin) (err error) {
	err = common.GEE_DB.Take(&model.Post{}, "post_id = ?", u.PostId).Updates(model.Post{
		Title:   u.Title,
		Content: u.Content,
	}).Error
	return
}

func (postDAO) DeletePostByPostId(postId int64) (err error) {
	err = common.GEE_DB.Delete(&model.Post{}, "post_id = ?", postId).Error
	return
}
