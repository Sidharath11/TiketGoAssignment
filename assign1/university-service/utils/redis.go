package utils

import (
	"encoding/json"
	"time"
	"university/domain"
	"university/domain/dto"
)

func Get(key string) (string, error) {
	val, err := domain.RedisClient.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func Store(key string, value string, expiration time.Duration) error {
	return domain.RedisClient.Set(key, value, expiration).Err()
}
func ToInfoJson(val []byte) dto.UniversityInfo {
	user := dto.UniversityInfo{}
	err := json.Unmarshal(val, &user)
	if err != nil {
		panic(err)
	}
	return user
}
