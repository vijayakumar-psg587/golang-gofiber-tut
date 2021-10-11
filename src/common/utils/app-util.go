package utils

import "github.com/golang-module/carbon"

func GetTimestamp() int64  {
  return carbon.Now().Timestamp()
}