package services

import (
	"university/domain/dao"
	"university/domain/dto"
	"university/utils"
)

func CreateUniversity(info *dto.University) (*dto.University, *utils.RestErr) {
	info, restErr := dao.Createuniversity(info)
	if restErr != nil {
		return nil, restErr

	}
	return info, nil
}

func FindUniversity(universityid int) (*dto.University, *utils.RestErr) {
	info, restErr := dao.Finduniversity(universityid)
	if restErr != nil {
		return nil, restErr

	}
	return info, nil
}

func DeleteUniversity(universityid int) *utils.RestErr {
	restErr := dao.Deleteuniversity(universityid)
	if restErr != nil {
		return restErr

	}
	return nil
}
func UpdateUniversity(info *dto.University) (*dto.University, *utils.RestErr) {
	info, restErr := dao.UpdateUniversity(info)
	if restErr != nil {
		return nil, restErr

	}
	return info, nil
}
