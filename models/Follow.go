package models

import "time"

type Follow struct {
	ID           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID       int64     `json:"userId"`
	FollowUserID int64     `json:"followUserid"`
	CreateTime   time.Time `json:"createTime"`
}

func SaveFollow(follow *Follow) {
	db.Table("tb_follow").Create(*follow)
}

func RemoveFollow(userid int64, followId int64) {
	db.Table("tb_follow").Where("user_id = ? and follow_user_id = ?", userid, followId).Delete(&Follow{})
}

func IsFollow(uid int64, followUid int64) int {
	var count int
	db.Table("tb_follow").Where("user_id = ? and follow_user_id = ?", uid, followUid).Count(&count)
	return count
}
