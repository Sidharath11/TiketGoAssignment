package dao

import (
	"context"
	"strconv"
	"time"
	"university/domain"
	"university/domain/dto"
	"university/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func Create(info *dto.UniversityInfo) (*dto.UniversityInfo, *utils.RestErr) {
	universityC := domain.Db.Collection("UniveristyInfo")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := universityC.InsertOne(ctx, bson.M{
		"university_id": info.University_id,
		"domain":        info.Domain,
		"web_page":      info.Web_page,
	})
	if err != nil {
		restErr := utils.InternalErr("can't insert hotelinfo to the database")
		return nil, restErr
	}
	info.ID = result.InsertedID.(primitive.ObjectID)

	return info, nil
}

func Find(university_id int) (*dto.UniversityInfo, *utils.RestErr) {
	var info dto.UniversityInfo
	universityC := domain.Db.Collection("UniveristyInfo")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	err := universityC.FindOne(ctx, bson.M{"university_id": university_id}).Decode(&info)
	if err != nil {
		restErr := utils.NotFound("university info not found with id " + strconv.Itoa(university_id))
		return nil, restErr
	}

	return &info, nil
}

func Delete(university_id int) *utils.RestErr {
	universityC := domain.Db.Collection("UniveristyInfo")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	info, err := universityC.DeleteOne(ctx, bson.M{"university_id": university_id})
	if err != nil {
		restErr := utils.NotFound("failed to delete university  with id " + strconv.Itoa(university_id))
		return restErr
	}
	if info.DeletedCount == 0 {
		restErr := utils.NotFound("university not exist with id " + strconv.Itoa(university_id))
		return restErr
	}
	return nil
}

func Update(info *dto.UniversityInfo) (*dto.UniversityInfo, *utils.RestErr) {
	universityC := domain.Db.Collection("UniveristyInfo")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := universityC.UpdateOne(ctx, bson.M{
		"university_id": info.University_id,
	}, bson.M{"$set": bson.M{"domain": info.Domain,
		"web_page": info.Web_page}})
	if err != nil {
		restErr := utils.InternalErr("can't update hotelinfo to the database")
		return nil, restErr
	}
	if result.MatchedCount == 0 {
		restErr := utils.NotFound("university not exist with id " + strconv.Itoa(info.University_id))
		return nil, restErr
	}
	if result.ModifiedCount == 0 {
		restErr := utils.NotFound("university not exist with id " + strconv.Itoa(info.University_id))
		return nil, restErr
	}
	newinfo, resterr := Find(info.University_id)
	if resterr != nil {
		return nil, resterr
	}
	return newinfo, resterr
}
