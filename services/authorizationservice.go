package services

import(
	"model"
)

func IsUserAuthorizedForResource (userId string, resourceId string) bool {
	if err := model.Db.Where("user_id = ? and resource_id = ?", userId, resourceId).First().Error(); err != nil {
		rows, err := model.Db.Table("resource_authorizations").Select("parent_resource_id").
		Joins("left join resource_groups on resource_groups.parent_resource_id = resource_authorizations.resource_id").
		Where("resource_groups.resource_id = ? and resource_authorizations.user_id = ?", resourceId, userId).Rows()
		if err != nil{
			return false
		}else if rows.Next(){
			return true
		}
	}
	return true
}