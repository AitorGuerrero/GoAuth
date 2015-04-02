package config

import (
	"github.com/AitorGuerrero/User/persistence/mongoDB"
)

func Get() Config {
	return Config {
		SqlDbConfig {
			Name: "BadassCity.Users",
			UserName: "root",
			Password: "",
			Host: "http://localhost",
		},
		mongoDB.MongoDBConfig {
			Server: "mongodb://dev:dev@ds037611.mongolab.com:37611/badass_city_user_manager",
		},
		KiteServiceConfig {
			Name: "BaddassCity.user",
			Version: "0.0.0",
			Port: 3635,
		},
	}
}
