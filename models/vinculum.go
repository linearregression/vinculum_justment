package models

import (
	"errors"
	"strconv"
	"time"
)

// Define a pointer to INflux DB

type Vinculum map[string]interface {}

func init() {
	// Setup Influx DB connection
}

func GetOne(VinculumId string) (object *Vinculum, err error) {
	// if v, ok := Objects[ObjectId]; ok {
	// 	return v, nil
	// }
	// if we can get a valid vinculum from influxdb
	    // return it
	// else return this error
	return nil, errors.New("ObjectId Not Exist")
}