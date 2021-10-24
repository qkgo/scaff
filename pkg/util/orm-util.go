package util

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type BatchInsertDB struct {
	*gorm.DB
}

func (db *BatchInsertDB) BatchInsertWrapper(objArr *[]interface{}, dontReplaceField []string) (int64, error) {
	if len(*objArr) == 0 {
		return 0, errors.New("insert a slice length of 0")
	}
	mainObj := (*objArr)[0]
	mainScope := db.NewScope(mainObj)
	mainFields := mainScope.Fields()
	quoted := make([]string, 0, len(mainFields))
	duplicate := make([]string, 0, len(mainFields))
	for i := range mainFields {
		if (mainFields[i].IsPrimaryKey && mainFields[i].IsBlank) || (mainFields[i].IsIgnored) {
			continue
		}
		quoted = append(quoted, mainScope.Quote(mainFields[i].DBName))
		for _, fieldKey := range dontReplaceField {
			if fieldKey != mainFields[i].DBName {
				duplicate = append(duplicate, mainScope.Quote(mainFields[i].DBName)+"=VALUES("+mainScope.Quote(mainFields[i].DBName)+")")
			}
		}
	}
	placeholdersArr := make([]string, 0, len(*objArr))
	for _, obj := range *objArr {
		scope := db.NewScope(obj)
		fields := scope.Fields()
		placeholders := make([]string, 0, len(fields))
		for i := range fields {
			if (fields[i].IsPrimaryKey && fields[i].IsBlank) || (fields[i].IsIgnored) {
				continue
			}
			var vars interface{}
			if fields[i].Name == "UpdatedAt" && fields[i].IsBlank {
				vars = gorm.NowFunc()
			} else {
				vars = fields[i].Field.Interface()
			}
			placeholders = append(placeholders, scope.AddToVars(vars))
		}
		placeholdersStr := "(" + strings.Join(placeholders, ", ") + ")"
		placeholdersArr = append(placeholdersArr, placeholdersStr)
		mainScope.SQLVars = append(mainScope.SQLVars, scope.SQLVars...)
	}
	if len(duplicate) > 0 {
		mainScope.Raw(fmt.Sprintf("INSERT INTO %s (%s) VALUES %s ON DUPLICATE KEY UPDATE %s",
			mainScope.QuotedTableName(),
			strings.Join(quoted, ", "),
			strings.Join(placeholdersArr, ", "),
			strings.Join(duplicate, ", "),
		))
	} else {
		mainScope.Raw(fmt.Sprintf("INSERT INTO %s (%s) VALUES %s",
			mainScope.QuotedTableName(),
			strings.Join(quoted, ", "),
			strings.Join(placeholdersArr, ", "),
		))
	}
	if err := mainScope.Exec().DB().Error; err != nil {
		return 0, err
	}
	return mainScope.DB().RowsAffected, nil
}

func BatchAdd(DB *gorm.DB, insertDataArray *[]interface{}, dontReplaceField []string) (int64, error) {
	DB1 := &BatchInsertDB{DB: DB}
	return DB1.BatchInsertWrapper(insertDataArray, dontReplaceField)
}
