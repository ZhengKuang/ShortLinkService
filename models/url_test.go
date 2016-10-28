package models

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"testing"
)

func init() {
	mongo_url := "127.0.0.1:3003"
	mongo_database := "short_url_test"
	session, err := mgo.Dial(mongo_url)
	if err != nil {
		panic(err)
	}
	_db = session.DB(mongo_database)
}

func TestMgo(t *testing.T) {
	var url = Url{
		Id:        1,
		SourceUrl: "http://www.qq.com",
	}

	err := url.Insert()
	assert.Nil(t, err)

}

func TestGenId(t *testing.T) {
	url := &Url{}
	url.SourceUrl = "http://www.facebook2.com"
	err := url.GenId()
	assert.Nil(t, err)

	err = url.Save()
	assert.Nil(t, err)
}
