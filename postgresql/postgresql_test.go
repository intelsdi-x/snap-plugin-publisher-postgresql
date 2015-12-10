// +build unit

/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package postgresql

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core/ctypes"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSliceToString(t *testing.T) {
	expl1 := []string{"intel", "os", "vmstat"}
	expl2 := []string{"intel", "os"}
	Convey("TestSliceToString", t, func() {
		sp := sliceToString(expl1)
		So(sp, ShouldEqual, "intel, os, vmstat")
		sp2 := sliceToString(expl2)
		So(sp2, ShouldEqual, "intel, os")
	})
}

func TestSliceToNamespace(t *testing.T) {
	expl1 := []string{"intel", "os", "vmstat"}
	expl2 := []string{"intel", "os"}
	Convey("TestSliceToNamespace", t, func() {
		sp := sliceToNamespace(expl1)
		So(sp, ShouldEqual, "intel.os.vmstat")
		sp2 := sliceToNamespace(expl2)
		So(sp2, ShouldEqual, "intel.os")
	})
}

func TestInterfaceToString(t *testing.T) {
	expl1 := []int{1, 2}
	expl2 := []string{"intel", "os"}
	expl3 := 1
	expl4 := 1.12
	expl5 := "snap"
	expl6 := make(map[float64]float64)
	expl7 := []int{}
	expl8 := []int{1}
	Convey("TestInterfaceToString", t, func() {
		sp, err := interfaceToString(expl1)
		So(sp, ShouldEqual, "1, 2")
		So(err, ShouldBeNil)
		sp, err = interfaceToString(expl2)
		So(sp, ShouldEqual, "intel, os")
		So(err, ShouldBeNil)
		sp, err = interfaceToString(expl3)
		So(sp, ShouldEqual, "1")
		So(err, ShouldBeNil)
		sp, err = interfaceToString(expl4)
		So(sp, ShouldEqual, "1.12")
		So(err, ShouldBeNil)
		sp, err = interfaceToString(expl5)
		So(sp, ShouldEqual, "snap")
		So(err, ShouldBeNil)
		sp, err = interfaceToString(expl6)
		So(sp, ShouldResemble, "")
		So(err, ShouldNotBeNil)
		sp, err = interfaceToString(expl7)
		So(sp, ShouldResemble, "")
		So(err, ShouldBeNil)
		sp, err = interfaceToString(expl8)
		So(sp, ShouldResemble, "1")
		So(err, ShouldBeNil)
	})
}

func TestGetConfigPolicy(t *testing.T) {
	Convey("TestGetConfigPolicy", t, func() {
		sp := NewPostgreSQLPublisher()
		expl, err := sp.GetConfigPolicy()
		So(expl, ShouldNotBeNil)
		So(err, ShouldBeNil)
		text := expl.Get([]string{""})
		So(fmt.Sprintf("%s", reflect.TypeOf(text)), ShouldResemble, "*cpolicy.ConfigPolicyNode")
	})
}

func TestGetPostgreSQLConn(t *testing.T) {
	config := make(map[string]ctypes.ConfigValue)
	config["username"] = ctypes.ConfigValueStr{Value: "root"}
	config["password"] = ctypes.ConfigValueStr{Value: "root"}
	config["database"] = ctypes.ConfigValueStr{Value: "SNAP_TEST"}
	config["table_name"] = ctypes.ConfigValueStr{Value: "info"}
	GetPostgreSQLConn := func(config map[string]ctypes.ConfigValue) (*sql.DB, error) {
		db, err := GetSQLMock()
		return db, err
	}
	GetPostgreSQLConn(config)

	Convey("TestGetPostgreSQL", t, func() {
		sp, err := GetPostgreSQLConn(config)
		So(sp, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}

func TestCreateTable(t *testing.T) {
	config := make(map[string]ctypes.ConfigValue)
	config["username"] = ctypes.ConfigValueStr{Value: "root"}
	config["password"] = ctypes.ConfigValueStr{Value: "root"}
	config["database"] = ctypes.ConfigValueStr{Value: "SNAP_TEST"}
	config["table_name"] = ctypes.ConfigValueStr{Value: "info"}
	GetPostgreSQLConn := func(config map[string]ctypes.ConfigValue) (*sql.DB, error) {
		db, err := GetSQLMock()
		return db, err
	}
	GetPostgreSQLConn(config)
	tableName := "info"

	Convey("TestGetPostgreSQL", t, func() {
		conn, err := GetPostgreSQLConn(config)

		sp, err := createTable(conn, tableName)
		So(sp, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}

func GetSQLMock() (*sql.DB, error) {
	db, mock, err := sqlmock.New()
	mock.ExpectExec("^CREATE TABLE IF NOT EXISTS (.+)$").WillReturnResult(sqlmock.NewResult(0, 1))
	if err != nil {
		fmt.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectExec("^CREATE INDEX key_index on (.+)$").WillReturnResult(sqlmock.NewResult(0, 1))
	return db, err
}

func TestPostgreSQLPublish(t *testing.T) {
	var buf bytes.Buffer
	//mock.ExpectBegin()
	expTime := time.Now()
	metrics := []plugin.PluginMetricType{
		*plugin.NewPluginMetricType([]string{"test_string"}, expTime, "", nil, nil, "example_string"),
		*plugin.NewPluginMetricType([]string{"test_int"}, expTime, "", nil, nil, 1),
		*plugin.NewPluginMetricType([]string{"test_int"}, expTime, "", nil, nil, true),
		*plugin.NewPluginMetricType([]string{"test_float"}, expTime, "", nil, nil, 1.12),
		*plugin.NewPluginMetricType([]string{"test_string_slice"}, expTime, "", nil, nil, []string{"str1", "str2"}),
		*plugin.NewPluginMetricType([]string{"test_string_slice"}, expTime, "", nil, nil, []int{1, 2}),
		*plugin.NewPluginMetricType([]string{"test_uint8"}, expTime, "", nil, nil, uint8(1)),
	}
	config := make(map[string]ctypes.ConfigValue)
	enc := gob.NewEncoder(&buf)
	enc.Encode(metrics)

	Convey("TestPostgreSQLPublish", t, func() {
		config["hostname"] = ctypes.ConfigValueStr{Value: "localhost"}
		config["port"] = ctypes.ConfigValueInt{Value: 5432}
		config["username"] = ctypes.ConfigValueStr{Value: "postgres"}
		config["password"] = ctypes.ConfigValueStr{Value: ""}
		config["database"] = ctypes.ConfigValueStr{Value: "snap_test"}
		config["table_name"] = ctypes.ConfigValueStr{Value: "info"}
		sp := NewPostgreSQLPublisher()
		So(sp, ShouldNotBeNil)
		err := sp.Publish("", buf.Bytes(), config)
		So(err, ShouldResemble, errors.New("Unknown content type ''"))
		err = sp.Publish(plugin.SnapGOBContentType, buf.Bytes(), config)
		meta := Meta()
		So(meta, ShouldNotBeNil)
	})
}
