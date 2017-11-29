package db

import "time"

type testContext struct {
	TableA *Set
}

type testTable struct {
	*Record
	Name string
	Age int
}

var testCtx *testContext

func newTestTable() testTable {
	return testTable{
		Record: &Record{
			ID: 0,
			Deleted: false,
			CreateDate: time.Now(),
		},
	}
}

func init(){
	testCtx = &testContext{
		TableA: NewSet(testTable{}),
	}
}