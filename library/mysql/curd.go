package mysql

import (
	"database/sql"

	"github.com/gogf/gf/database/gdb"
)

//// 通用主库执行SQL（可用于插入、更新、删除操作）
//func ExecQuery(sqlStr string, data []interface{}, tx ...*gdb.TX) (*sql.Rows, error) {
//	var result sql.Result
//	var err error
//
//	if tx != nil { // 开启事务
//		result, err = tx[0].Query(sqlStr, data...)
//	} else {
//		result, err = ds.Write.Query(sqlStr, data...)
//	}
//
//	return result, err
//}

// 通用主库执行SQL（可用于插入、更新、删除操作）
func ExecWrite(sqlStr string, data []interface{}, tx ...*gdb.TX) (sql.Result, error) {
	var result sql.Result
	var err error

	if tx != nil { // 开启事务
		result, err = tx[0].Exec(sqlStr, data...)
	} else {
		result, err = ds.Write.Exec(sqlStr, data...)
	}

	return result, err
}

// 通用查询SQL
func ExecReadList(sqlStr string, data []interface{}, tx ...*gdb.TX) (gdb.Result, error) {
	var result gdb.Result
	var err error

	if tx != nil { // 开启事务
		result, err = tx[0].GetAll(sqlStr, data...)
	} else {
		result, err = ds.Read.GetAll(sqlStr, data...)
	}

	return result, err
}

// 通用查询SQL
func ExecReadOne(sqlStr string, data []interface{}, tx ...*gdb.TX) (gdb.Record, error) {
	var record gdb.Record
	var err error

	if tx != nil { // 开启事务
		record, err = tx[0].GetOne(sqlStr, data...)
	} else {
		record, err = ds.Read.GetOne(sqlStr, data...)
	}

	return record, err
}

// 通用统计SQL
func ExecReadCount(sqlStr string, data []interface{}, tx ...*gdb.TX) (int, error) {
	var count int
	var err error

	if tx != nil { // 开启事务
		count, err = tx[0].GetCount(sqlStr, data...)
	} else {
		count, err = ds.Read.GetCount(sqlStr, data...)
	}

	return count, err
}
