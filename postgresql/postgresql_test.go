// +build small

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
	"github.com/intelsdi-x/snap/core"
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
	Convey("TestInterfaceToString", t, func() {

		Convey("Calling function for numeric types", func() {
			expl1 := uint(1)
			expl2 := int(-2)
			expl3 := float32(-2.4)
			expl4 := float64(-2.4)

			sp, err := interfaceToString(expl1)
			So(sp, ShouldEqual, "1")
			So(err, ShouldBeNil)

			sp, err = interfaceToString(expl2)
			So(sp, ShouldEqual, "-2")
			So(err, ShouldBeNil)

			sp, err = interfaceToString(expl3)
			So(sp, ShouldEqual, "-2.4")
			So(err, ShouldBeNil)

			sp, err = interfaceToString(expl4)
			So(sp, ShouldEqual, "-2.4")
			So(err, ShouldBeNil)
		})

		Convey("Calling function for slices with numeric types", func() {
			expl1 := []uint{1, 2}
			expl2 := []int{1, -2}
			expl3 := []float32{1.5, -2.4}
			expl4 := []float64{1.5, -2.4}

			sp, err := interfaceToString(expl1)
			So(sp, ShouldEqual, "1, 2")
			So(err, ShouldBeNil)

			sp, err = interfaceToString(expl2)
			So(sp, ShouldEqual, "1, -2")
			So(err, ShouldBeNil)

			sp, err = interfaceToString(expl3)
			So(sp, ShouldEqual, "1.5, -2.4")
			So(err, ShouldBeNil)

			sp, err = interfaceToString(expl4)
			So(sp, ShouldEqual, "1.5, -2.4")
			So(err, ShouldBeNil)
		})

		Convey("Calling function for string type", func() {
			expl1 := string("snap")
			sp, err := interfaceToString(expl1)
			So(sp, ShouldEqual, "snap")
			So(err, ShouldBeNil)
		})

		Convey("Calling function for slices with string type", func() {
			expl1 := []string{"snap1", "snap2"}
			expl2 := []string{"snap snap1", "snap snap2"}
			expl3 := []string{"[]snap1[]", "[]snap2[]"}

			sp, err := interfaceToString(expl1)
			So(sp, ShouldEqual, "snap1, snap2")
			So(err, ShouldBeNil)

			sp, err = interfaceToString(expl2)
			So(sp, ShouldEqual, "snap snap1, snap snap2")
			So(err, ShouldBeNil)

			sp, err = interfaceToString(expl3)
			So(sp, ShouldEqual, "[]snap1[], []snap2[]")
			So(err, ShouldBeNil)

		})

		Convey("Calling function for bool type", func() {
			expl1 := bool(true)
			expl2 := bool(false)

			sp, err := interfaceToString(expl1)
			So(sp, ShouldEqual, "1")
			So(err, ShouldBeNil)

			sp, err = interfaceToString(expl2)
			So(sp, ShouldEqual, "0")
			So(err, ShouldBeNil)
		})

		Convey("Calling function for unsupported types", func() {
			expl1 := map[float64]float64{}
			expl2 := struct{}{}

			sp, err := interfaceToString(expl1)
			So(sp, ShouldEqual, "")
			So(err, ShouldNotBeNil)

			sp, err = interfaceToString(expl2)
			So(sp, ShouldEqual, "")
			So(err, ShouldNotBeNil)
		})
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
	metrics := []plugin.MetricType{
		*plugin.NewMetricType(core.NewNamespace("test_string"), expTime, nil, "", "example_string"),
		*plugin.NewMetricType(core.NewNamespace("test_int"), expTime, nil, "", int(-1)),
		*plugin.NewMetricType(core.NewNamespace("test_int64"), expTime, nil, "", int64(1)),
		*plugin.NewMetricType(core.NewNamespace("test_uint"), expTime, nil, "", uint(1)),
		*plugin.NewMetricType(core.NewNamespace("test_uint64"), expTime, nil, "", uint64(1)),
		*plugin.NewMetricType(core.NewNamespace("test_float"), expTime, nil, "", 1.12),
		*plugin.NewMetricType(core.NewNamespace("test_float64"), expTime, nil, "", float64(-1.23)),
		*plugin.NewMetricType(core.NewNamespace("test_bool"), expTime, nil, "", true),

		*plugin.NewMetricType(core.NewNamespace("test_string_slice"), expTime, nil, "", []string{"str1", "str2"}),
		*plugin.NewMetricType(core.NewNamespace("test_int_slice"), expTime, nil, "", []int{-1, 2}),
		*plugin.NewMetricType(core.NewNamespace("test_int64_slice"), expTime, nil, "", []int64{-1, 2}),
		*plugin.NewMetricType(core.NewNamespace("test_uint_slice"), expTime, nil, "", []uint{1, 2}),
		*plugin.NewMetricType(core.NewNamespace("test_uint64_slice"), expTime, nil, "", []uint64{1, 2}),
		*plugin.NewMetricType(core.NewNamespace("test_float64_slice"), expTime, nil, "", []float64{1.23, -1.23}),
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
