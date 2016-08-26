package models

import (
	"errors"
  "encoding/json"
  "fmt"
  "net/url"
  // "reflect"

  "github.com/astaxie/beego/logs"
	"github.com/influxdata/influxdb/client"
)

// Define a pointer to INflux DB

type Vinculum []map[string]interface {}

const (
	VinculumDB = "vinculums"
	// url = "http://localhost:8086/"
)


func init() {
	// Setup Influx DB connection
}

func GetOne(VinculumId string) (object *Vinculum, err error) {

	ob := make(Vinculum, 0)

	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		logs.Error(err)
	}

	conf := client.Config{ URL: *host }
  if err != nil { logs.Error("Error: ", err) }

  con, err := client.NewClient(conf)
	if err != nil { logs.Error(err) }

  query_string := fmt.Sprintf("SELECT * FROM json where outpath = '%s'", VinculumId)
  logs.Debug(query_string)
  q := client.Query{ Command: query_string, Database: VinculumDB }

  if response, err := con.Query(q); err == nil && response.Error() == nil {

  	retrieved_values := response.Results[0].Series[0].Values
  	// Values is a data structure that represents each row in the DB
  	// For now, it's [time <json> outpath owner]
  	for x := range retrieved_values {
  		logs.Debug(retrieved_values[x])

  		new_vinculum := make(map[string]interface {})
  		new_vinculum["time"] = retrieved_values[x][0]

  		json_string := retrieved_values[x][1]
  		var new_json interface{}
  		json_bytes := []byte(json_string.(string))
  		
  		json.Unmarshal(json_bytes, &new_json)
  		new_vinculum["data"] = new_json
  		ob = append(ob, new_vinculum)
  	}

  	return &ob, nil
  }

	return nil, errors.New("ObjectId Not Exist")
}