package services

import (
	"university/domain/dao"
	"university/domain/dto"
	"university/utils"
)

func CreateUniversityInfo(info *dto.UniversityInfo) (*dto.UniversityInfo, *utils.RestErr) {
	info, restErr := dao.Create(info)
	if restErr != nil {
		return nil, restErr

	}
	return info, nil
}

func FindUniversityInfo(universityid int) (*dto.UniversityInfo, *utils.RestErr) {
	info, restErr := dao.Find(universityid)
	if restErr != nil {
		return nil, restErr

	}
	return info, nil
}

func DeleteUniversityInfo(universityid int) *utils.RestErr {
	restErr := dao.Delete(universityid)
	if restErr != nil {
		return restErr

	}
	return nil
}
func UpdateUniversityInfo(info *dto.UniversityInfo) (*dto.UniversityInfo, *utils.RestErr) {
	info, restErr := dao.Update(info)
	if restErr != nil {
		return nil, restErr

	}
	return info, nil
}
