package mysql

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// 获取插入sql，只生成非零值的属性
func InsertSql(table string, data interface{}) (string, []interface{}) {
	if data == nil {
		return "", nil
	}

	oType, oValue := getReflect(data)

	columnsStr := strings.Builder{}
	valuesStr := strings.Builder{}
	values := ([]interface{})(nil)

	for i := 0; i < oType.NumField(); i++ {
		filed := oValue.Field(i)
		if filed.IsValid() && filed.IsZero() {
			continue
		}

		field := oType.Field(i)
		ormTag := field.Tag.Get("orm")
		if ormTag == "" {
			continue
		}
		ormTag = strings.Split(ormTag, ",")[0]
		if ormTag == "-" {
			continue
		}

		// 判断gtime.time 类型 零值判断跳过
		if filed.Kind().String() == "ptr" {
			if filed.Interface().(*gtime.Time).IsZero() {
				continue
			}
		}

		if columnsStr.Len() > 0 {
			columnsStr.WriteString(",")
			valuesStr.WriteString(",")
		}

		columnsStr.WriteString("`")
		columnsStr.WriteString(ormTag)
		columnsStr.WriteString("`")
		valuesStr.WriteString("?")

		values = append(values, filed.Interface())
	}

	sqlStr := fmt.Sprintf("insert into %s (%s) values(%s)", table, columnsStr.String(), valuesStr.String())

	return sqlStr, values
}

// 获取更新sql，只生成非零值的属性，当condition存在时则生成相应的where条件
func UpdateSql(table string, data interface{}, condition g.Slice) (string, []interface{}) {
	if data == nil {
		return "", nil
	}

	oType, oValue := getReflect(data)

	// 解析set内容
	setStr := strings.Builder{}
	values := ([]interface{})(nil)
	for i := 0; i < oType.NumField(); i++ {
		filed := oValue.Field(i)
		field := oType.Field(i)
		ormTag := field.Tag.Get("orm")
		if filed.IsValid() && filed.IsZero() && !strings.Contains(ormTag, ",needModifyZeroWhenUpdate") {
			continue
		}

		ormTag = strings.Split(ormTag, ",")[0]

		if ormTag == "-" {
			continue
		}

		if ormTag == "" {
			continue
		}

		// 判断gtime.time 类型 零值判断跳过
		if filed.Kind().String() == "ptr" {
			var time interface{} = (*gtime.Time)(nil)
			if filed.Interface() != nil && filed.Interface() != time {
				if filed.Interface().(*gtime.Time).IsZero() {
					continue
				}
			}
		}

		if setStr.Len() > 0 {
			setStr.WriteString(" , ")
		}

		setStr.WriteString("`")
		setStr.WriteString(ormTag)
		setStr.WriteString("` = ?")

		values = append(values, filed.Interface())
	}

	// 解析where条件
	whereStr := strings.Builder{}

	if condition != nil {
		for i, v := range condition {
			// 偶数
			if i%2 == 0 {
				whereStr.WriteString(" and ")
				whereStr.WriteString(v.(string))
			} else { // 奇数
				values = append(values, v)
			}
		}
	}

	sqlStr := fmt.Sprintf("update %s set %s where 1=1 %s", table, setStr.String(), whereStr.String())

	return sqlStr, values
}

// 乐观锁更新sql，只生成非零值的属性，当condition存在时则生成相应的where条件
func UpdateSqlByOptimisticLock(table string, data interface{}, condition g.Slice,
	optimisticLock string) (string, []interface{}) {
	if data == nil {
		return "", nil
	}

	oType, oValue := getReflect(data)

	// 解析where条件
	whereStr := strings.Builder{}

	// 解析set内容
	setStr := strings.Builder{}
	values := ([]interface{})(nil)
	for i := 0; i < oType.NumField(); i++ {
		fieldValue := oValue.Field(i)
		field := oType.Field(i)
		ormTag := field.Tag.Get("orm")
		ormTag = strings.Split(ormTag, ",")[0]

		if fieldValue.IsValid() && fieldValue.IsZero() {
			// 乐观锁字段进行自增处理
			//if optimisticLock == ormTag {
			//	if setStr.Len() > 0 {
			//		setStr.WriteString(" , ")
			//	}
			//
			//	setStr.WriteString("`")
			//	setStr.WriteString(ormTag)
			//	setStr.WriteString("` = ")
			//	setStr.WriteString("`")
			//	setStr.WriteString(ormTag)
			//	setStr.WriteString("` + 1 ")
			//}
			continue
		}

		if setStr.Len() > 0 {
			setStr.WriteString(" , ")
		}

		// 乐观锁字段进行自增处理
		if optimisticLock == ormTag {
			//if setStr.Len() > 0 {
			//	setStr.WriteString(" , ")
			//}

			setStr.WriteString("`")
			setStr.WriteString(ormTag)
			setStr.WriteString("` = ")
			setStr.WriteString("`")
			setStr.WriteString(ormTag)
			setStr.WriteString("` + 1 ")

			// where条件中自动加上乐观锁限制
			whereStr.WriteString(" and ")
			whereStr.WriteString("`")
			whereStr.WriteString(ormTag)
			whereStr.WriteString("` = ")
			whereStr.WriteString(strconv.FormatInt(fieldValue.Int(), 10))
		} else {
			setStr.WriteString("`")
			setStr.WriteString(ormTag)
			setStr.WriteString("` = ?")
		}

		if optimisticLock != ormTag {
			values = append(values, fieldValue.Interface())
		}

	}

	if condition != nil {
		for i, v := range condition {
			// 偶数
			if i%2 == 0 {
				whereStr.WriteString(" and ")
				whereStr.WriteString(v.(string))
			} else { // 奇数
				values = append(values, v)
			}
		}
	}

	sqlStr := fmt.Sprintf("update %s set %s where 1=1 %s", table, setStr.String(), whereStr.String())

	return sqlStr, values
}

// 查询
func SelectSql(table string, pageInfo bool, condition g.SliceStr, orderStr ...string) (string, string) {
	// 解析where条件
	whereStr := strings.Builder{}

	if condition != nil {
		for _, v := range condition {
			whereStr.WriteString(" and ")
			whereStr.WriteString(v)
		}
	}

	sqlStr := fmt.Sprintf("select * from %s where 1=1%s", table, whereStr.String())
	pageSqlStr := fmt.Sprintf("select * from %s where 1=1%s", table, whereStr.String())

	if orderStr != nil {
		sqlStr = sqlStr + " " + orderStr[0]
		pageSqlStr = pageSqlStr + " " + orderStr[0]
	}

	if pageInfo {
		limit := " limit ?,?"
		pageSqlStr = pageSqlStr + limit
	}

	return sqlStr, pageSqlStr
}

// 查询 count
func SelectCountSql(table string, pageInfo bool, condition g.SliceStr, orderStr ...string) (string, string) {
	// 解析where条件
	whereStr := strings.Builder{}

	if condition != nil {
		for _, v := range condition {
			whereStr.WriteString(" and ")
			whereStr.WriteString(v)
		}
	}

	sqlStr := fmt.Sprintf("select count(*) from %s where 1=1%s", table, whereStr.String())
	pageSqlStr := fmt.Sprintf("select count(*) from %s where 1=1%s", table, whereStr.String())

	if orderStr != nil {
		sqlStr = sqlStr + " " + orderStr[0]
		pageSqlStr = pageSqlStr + " " + orderStr[0]
	}

	if pageInfo {
		limit := " limit ?,?"
		pageSqlStr = pageSqlStr + limit
	}

	return sqlStr, pageSqlStr
}

// 删除 -- 生成语句
func DeleteSql(table string, condition g.Slice) (string, []interface{}) {
	values := ([]interface{})(nil)

	// 解析where条件
	whereStr := strings.Builder{}

	if condition != nil {
		for i, v := range condition {
			if i%2 == 0 {
				whereStr.WriteString(" and ")
				whereStr.WriteString(fmt.Sprintf("%v", v))
			} else { // 奇数
				values = append(values, v)
			}
		}
	}

	sqlStr := fmt.Sprintf("delete from %s where 1=1%s", table, whereStr.String())

	return sqlStr, values
}

// 获取反射信息
func getReflect(o interface{}) (reflect.Type, reflect.Value) {
	oType := reflect.TypeOf(o)
	oValue := reflect.ValueOf(o)
	return oType, oValue
}
