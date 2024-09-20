package mysql

import (
	"iv-test/library/logger"

	"github.com/gogf/gf/database/gdb"
)

// 执行事务，一般用于修改，删除、插入，查询操作不用
func ExecTransaction(db gdb.DB, f func(tx *gdb.TX) error) error {
	tx, err := db.Begin()
	if err != nil {
		logger.Error(err)
		return err
	}
	// 固定执行Rollback，成功时释放连接，失败时回滚
	defer tx.Rollback()

	err = f(tx)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
