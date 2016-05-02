package services

import(
	"model"
)

func IsUserAuthorizedForResourceRead(userId string, resourceId string) bool {
	user, err:= model.GetUserById(userId)
	if err == nil && user.IsSuperUser{return true}

	resourceAuth := model.ResourceAuthorization{}

	err = model.Db.Where("user_id = ? and resource_id = ? ", userId, resourceId).First(&resourceAuth).Error
	if err != nil {
		rows, err2 := model.Db.Table("resource_authorizations").Select("parent_resource_id").
		Joins("left join resource_groups on resource_groups.parent_resource_id = resource_authorizations.resource_id").
		Where("resource_groups.resource_id = ? and resource_authorizations.user_id = ?", resourceId, userId).Rows()
		if err2 != nil{
			return false
		}else if rows.Next(){
			return true
		}
	} else {
		return true
	}

	return false
}

func IsAuthorizedForResourceEditIsAuthorizedForResourceEdit(userId string, resourceId string) bool {
	user, err:= model.GetUserById(userId)
	if err == nil && user.IsSuperUser{return true}

	resourceAuth := model.ResourceAuthorization{}
	err = model.Db.Where("user_id = ? and resource_id = ? and action = ?", userId, resourceId, model.EditResourceAction).First(&resourceAuth).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func IsUserAuthorizedForJournalRead (userId string, journalId string) bool {
	journal, err := model.GetJournalById(journalId)

	if err != nil { return false }

	if journal.IsPublic {return true}

	return IsUserAuthorizedForResourceRead(userId, journalId)
}


