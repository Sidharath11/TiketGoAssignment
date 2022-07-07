package dao

import (
	"university/domain"
	"university/domain/dto"
	"university/utils"
)

func Createuniversity(info *dto.University) (*dto.University, *utils.RestErr) {
	db := domain.PSQLDB

	sqlStatement := `INSERT INTO University (info_id, name, locationid) VALUES ($1, $2, $3) RETURNING id`

	// the inserted id will store in this id
	var id int

	err := db.QueryRow(sqlStatement, info.Info_id, info.Name, info.Locationid).Scan(&id)

	if err != nil {
		restErr := utils.InternalErr("can't insert university data to the database")
		return nil, restErr
	}
	info.ID = id

	return info, nil
}

func Finduniversity(university_id int) (*dto.University, *utils.RestErr) {

	db := domain.PSQLDB

	sqlStatement := `Select * from University where id=$1 returning *`

	// the inserted id will store in this id
	var row dto.University

	err := db.QueryRow(sqlStatement, university_id).Scan(&row.ID, &row.Info_id, &row.Name, &row.Locationid)

	if err != nil {
		restErr := utils.InternalErr("can't get university data from the database")
		return nil, restErr
	}

	return &row, nil

}

func Deleteuniversity(university_id int) *utils.RestErr {

	db := domain.PSQLDB

	sqlStatement := `delete * from University where id=$1`

	_, err := db.Exec(sqlStatement, university_id)

	if err != nil {
		restErr := utils.InternalErr("can't delete university data from the database")
		return restErr
	}

	return nil
}

func UpdateUniversity(info *dto.University) (*dto.University, *utils.RestErr) {
	db := domain.PSQLDB

	sqlStatement := `update University set info_id=$1,name=$2,location_id=$3 where id=$4`

	_, err := db.Exec(sqlStatement, info.Info_id, info.Name, info.Locationid, info.ID)

	if err != nil {
		restErr := utils.InternalErr("can't update university data from the database")
		return nil, restErr
	}
	dt, er := Finduniversity(info.ID)
	if er != nil {
		restErr := utils.InternalErr("can't fetch  university data after update from the database")
		return nil, restErr
	}
	return dt, nil
}
