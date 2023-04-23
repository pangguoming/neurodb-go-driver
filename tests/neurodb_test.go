// Package   neurodb
// @file     neurodb_test.go
// @author   Hoss
// @contact  hth146@163.com
// @time     2023/4/23

package tests

import (
	"fmt"
	"neurodb.org/neurodb"
	"testing"
)

// TODO Hoss The development has not been completed, and the test case cannot pass
func TestDriver(t *testing.T) {
	db, err := neurodb.Open("127.0.0.1", 8839)
	if err != nil {
		t.Error(err)
	}

	resultSet ,err := db.ExecuteQuery("match (n) return n")
	if err != nil {
		t.Error(err)
	}
	for resultSet.Next() {
		fmt.Println(resultSet.Record())
	}
	resultSet,err = db.ExecuteQuery("match (n)-[r]->(m) return n,r,m")

	if err != nil {
		t.Error(err)
	}
	for resultSet.Next() {
		fmt.Println(resultSet.Record())
	}

}
