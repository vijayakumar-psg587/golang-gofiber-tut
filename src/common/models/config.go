package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"

	"github.com/tidwall/gjson"
	cerror "github.com/vijayakumar-psg587/golang-fiber-tut/src/common/models/errors"
	"github.com/vijayakumar-psg587/golang-fiber-tut/src/common/utils"
	"io"
	"os"
	"path/filepath"
	"strings"

	"sync"
)

var (
	once sync.Once
	file os.File
)

type CommonConfig struct {
	AppConfig
	DatabaseConfig
}

type AppConfig struct {
	Name string
	Port int32
}
type DatabaseConfig struct {
	Connection string
	UserName   string
	Password   string
}

func (appConfig *AppConfig) isExternalized() bool {
	return true
}

func (databaseConfig *DatabaseConfig) isExternalized() bool {
	return true
}

var (
	commonConfig *CommonConfig
	appConfigModel      *AppConfig
	databaseConfigModel *DatabaseConfig
)

// GetConfigFromEnv /**
func GetConfigFromEnv() []byte {
	commonConfig := new(CommonConfig) // IMP : new is used and it returns a pointer , make is used to return the type & val
	if os.Getenv("APP_ENV") == "dev" {
		// load the env from file
		cwd, _ := os.Getwd()
		if fp, err := filepath.Abs(cwd+"/src/config/development/app-dev.json" ); err == nil {
		  	viper.SetConfigFile(fp)
			err := viper.ReadInConfig()
			if err != nil {
				return nil
			} else {
				commonConfig = &CommonConfig{
					AppConfig{
						Name: viper.GetString("app.name"),
						Port: viper.GetInt32("app.port"),
					},
					DatabaseConfig{
						Connection: viper.GetString("database.connection"),
						UserName:   viper.GetString("database.userName"),
						Password:   viper.GetString("database.password"),
					},
				}
			}

		} else {
			return nil
		}
	}
	jsonByte, _ := json.Marshal(*commonConfig)
	return jsonByte
}

func GetConfig() ([]byte, []byte) {
	commonErr :=make([]cerror.CustomError, 0)

	once.Do(func () {
	commonConfig = new(CommonConfig)
	appConfigModel = new(AppConfig)
	databaseConfigModel = new(DatabaseConfig)
	fileByte := make([]byte, 100)

	if str, pathErr := filepath.Rel("./", "src/config/development/app-dev.json"); pathErr == nil {
		fp, fpErr := os.Open(str)
		if fpErr == nil {
			resStr, err := getFileDataAsString(&fileByte, fp)
			if err == nil {
				// if u get an interface first make sure to check if we get back a string
				if value, ok := resStr.(string); ok {
					appConfigModel.Name =  gjson.Get(value, "app.name").String()
					appConfigModel.Port = int32(gjson.Get(value, "app.port").Int())
					databaseConfigModel.Connection = gjson.Get(value, "database.connection").String()
					databaseConfigModel.UserName = gjson.Get(value, "database.userName").String()
					databaseConfigModel.Password = gjson.Get(value, "database.password").String()
					commonConfig.AppConfig = *appConfigModel
					commonConfig.DatabaseConfig = *databaseConfigModel
				} else {
					commonErr = append(commonErr, cerror.CustomError{Err: errors.New("Cannot read the file or convert to json"),
						Code: "500", Status: 500, Timestamp: utils.GetTimestamp()})
				}
			} else {
				commonErr  = append(commonErr, cerror.CustomError{Err: err,
					Code: "500", Status: 500, Timestamp: utils.GetTimestamp()})
			}
			fpCloseErr := fp.Close()
			if fpCloseErr != nil {
				commonErr =  append(commonErr, cerror.CustomError{Err: fpErr,
					Code: "500", Status: 500, Timestamp: utils.GetTimestamp()})
			}
		} else {
			commonErr = append(commonErr, cerror.CustomError{Err: fpErr,
				Code: "500", Status: 500, Timestamp: utils.GetTimestamp()})
		}
	}
	})
	if len(commonErr) >=1 {
		jsonErr,_ := json.Marshal(commonErr)
		return nil, jsonErr
	} else {
		b, err := json.Marshal(commonConfig)
		if err != nil {
			fmt.Println("error:", err)
		}

		return b,nil
	}

	}

/**
Get file bytes and send it as generic interface, reason being the file can be a yaml or prop file as well
*/
func getFileDataAsString(byteData *[]byte, fileFp *os.File) (interface{}, error) {
	var readString strings.Builder
	for {
		intByte, err := fileFp.Read(*byteData)
		if err != nil {
			if err != io.EOF {
				return AppConfig{}, err
			} else {
				break
			}
		}
		readString.Write((*byteData)[:intByte])

	}
	return readString.String(), nil
}

func getDatabaseConfig(ctx context.Context) DatabaseConfig {
	return *databaseConfigModel
}
