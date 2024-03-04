package helper

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"reflect"
	"strconv"
	"strings"
)

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	}
}

func GetCriteriaStatement(ctx *fiber.Ctx, model interface{}) (statement string, err error) {
	defer func() {
		if r := recover(); r != nil {
			// fmt.Println(r)
			err = errors.New("failed on generating criteria")
		}
	}()

	criteria := ctx.Query("criteria")

	var pairParams []string

	if criteria != "" {
		pairParams = strings.Split(criteria, ",")
	}

	for i, v := range pairParams {
		pair := strings.Split(v, ":")
		key := pair[0]
		val := pair[1]

		v := reflect.ValueOf(model)
		typeOfModel := v.Type()
		var typeOfField string
		var columnOrmTagFieldName string
		isExact := strings.Contains(key, "(exact)")
		if isExact {
			key = strings.ReplaceAll(key, "(exact)", "")
		}
		isIn := strings.Contains(key, "(in)")
		if isIn {
			key = strings.ReplaceAll(key, "(in)", "")
		}

		for j := 0; j < v.NumField(); j++ {
			if typeOfModel.Field(j).Tag.Get("search") == key {
				typeOfField = typeOfModel.Field(j).Type.Name()
				ormTag := typeOfModel.Field(j).Tag.Get("gorm")
				ormTagValues := strings.Split(ormTag, ";")

				for _, val := range ormTagValues {
					if strings.Contains(val, "column") {
						columnOrmTagFieldName = strings.Split(val, ":")[1]
					}
				}
			}
		}

		if typeOfField == "" {
			continue
		}

		if i > 0 {
			statement = statement + " and "
		}

		if typeOfField == "string" || typeOfField == "time.Time" {
			// fmt.Println(isIn)
			if isExact {
				// fmt.Println(val)
				statement = statement + fmt.Sprintf(`LOWER("%s") = LOWER('%v')`, columnOrmTagFieldName, val)
				continue
			}
			if isIn {
				// fmt.Println(val)
				// statement = statement + fmt.Sprintf(`LOWER("%s") in LOWER('%v')`, columnOrmTagFieldName, val)
				continue
			}
			statement = statement + fmt.Sprintf("LOWER(\"%s\") ~*'%v'", columnOrmTagFieldName, val)
		} else if typeOfField == "UUID" {
			statement = statement + fmt.Sprintf("%s = '%v'", columnOrmTagFieldName, val)
		} else if (typeOfField == "uint" || typeOfField == "int") && isIn {
			statement = statement + fmt.Sprintf("%s in (%v)", columnOrmTagFieldName, strings.ReplaceAll(val, "|", ","))
		} else if typeOfField == "RawMessage" {
			/* !! ONLY ONE LEVEL YET */
			if isExact {
				statement = statement + fmt.Sprintf("%s->>'%v'='%v'", columnOrmTagFieldName, strings.Split(val, "->")[0], strings.Split(val, "->")[1])
			} else {
				statement = statement + fmt.Sprintf("%s->>'%v'~*'%v'", columnOrmTagFieldName, strings.Split(val, "->")[0], strings.Split(val, "->")[1])
			}
		} else {
			statement = statement + fmt.Sprintf("%s = %v", columnOrmTagFieldName, val)
		}
	}

	return
}

func GetOrderStatement(ctx *fiber.Ctx) (statement string, err error) {
	paramValue := ctx.Query("sort")
	var orderBy string
	var orderType string

	if paramValue == "" {
		return
	}

	params := strings.Split(paramValue, ":")

	if params[0] != "" {
		orderBy = params[0]
	} else {
		return
	}

	if params[1] != "" {
		orderType = params[1]
	} else {
		orderType = "ASC"
	}

	statement = fmt.Sprintf("ORDER BY %s %s", orderBy, orderType)

	return

	// orderBy := ctx.Query("order_by")
	// orderType := ctx.Query("order_type")

	// if orderBy != "" {
	// 	if orderType != "" {
	// 		query = fmt.Sprintf("ORDER BY `%s` %s", orderBy, orderType)
	// 	} else {
	// 		query = fmt.Sprintf("ORDER BY `%s` asc", orderBy)
	// 	}
	// }

	// return
}

func GetOrderMap(ctx *fiber.Ctx) (orderMap map[string]string, err error) {
	paramValue := ctx.Query("sort")
	orderMap = map[string]string{}
	// var orderBy string
	// var orderType string

	if paramValue == "" {
		return
	}

	var pairParams []string

	if paramValue != "" {
		pairParams = strings.Split(paramValue, ",")
	}

	for _, v := range pairParams {
		pair := strings.Split(v, ":")
		// key := pair[0]
		// val := pair[1]
		orderMap[pair[0]] = pair[1]
		// orderMap = append(orderMap, })
	}

	// params := strings.Split(paramValue, ":")

	// if params[0] != "" {
	// 	orderBy = params[0]
	// } else {
	// 	return
	// }

	// if params[1] != "" {
	// 	orderType = params[1]
	// } else {
	// 	orderType = "ASC"
	// }

	// statement = fmt.Sprintf("ORDER BY %s %s", orderBy, orderType)

	return

	// orderBy := ctx.Query("order_by")
	// orderType := ctx.Query("order_type")

	// if orderBy != "" {
	// 	if orderType != "" {
	// 		query = fmt.Sprintf("ORDER BY `%s` %s", orderBy, orderType)
	// 	} else {
	// 		query = fmt.Sprintf("ORDER BY `%s` asc", orderBy)
	// 	}
	// }

	// return
}

func GetOrderString(ctx *fiber.Ctx) (orderText string, err error) {
	var orders []string
	paramValue := ctx.Query("sort")
	if paramValue == "" {
		orderText = orderText + "created_at desc"
		return
	}

	params := strings.Split(paramValue, ",")

	for _, param := range params {
		splittedParam := strings.Split(param, ":")

		order := fmt.Sprintf("`%s` %s", splittedParam[0], splittedParam[1])
		orders = append(orders, order)
	}

	orderText = strings.Join(orders[:], ",")

	orderText = orderText + ", created_at desc"

	return
}

func GetCriteria(ctx *fiber.Ctx) (criterias []string, err error) {
	defer catch()

	paramValue := ctx.Query("criteria")
	if paramValue == "" {
		return
	}

	params := strings.Split(paramValue, ",")

	for _, param := range params {
		splittedParam := strings.Split(param, ":")
		criteria := fmt.Sprintf("`%s` = '%s'", splittedParam[0], splittedParam[1])
		criterias = append(criterias, criteria)
	}

	return
}

func GetSearchStatement(ctx *fiber.Ctx) (statement string, err error) {
	var searches []string

	paramValue := ctx.Query("search")
	if paramValue == "" {
		return
	}

	params := strings.Split(paramValue, ",")

	for _, param := range params {
		splittedParam := strings.Split(param, ":")
		search := "`" + splittedParam[0] + "` LIKE '%" + splittedParam[1] + "%'"
		searches = append(searches, search)
	}

	for i, search := range searches {
		if i != 0 {
			statement += " OR " + search
		} else {
			statement += search
		}
	}

	return
}

func GetRangeStatement(ctx *fiber.Ctx) (statement string, err error) {
	var ranges []string

	paramValue := ctx.Query("range")
	if paramValue == "" {
		return
	}

	params := strings.Split(paramValue, ",")

	for _, param := range params {
		splittedParam := strings.Split(param, ":")
		rangeSplitted := strings.Split(splittedParam[1], "|")

		if rangeSplitted[0] == "" || rangeSplitted[1] == "" {
			return
		}

		rangeText := "`" + splittedParam[0] + "` BETWEEN " + rangeSplitted[0] + " AND " + rangeSplitted[1]

		if rangeSplitted[0] == "" {
			rangeText = "`" + splittedParam[0] + "` <= " + rangeSplitted[1]
		} else if rangeSplitted[1] == "" {
			rangeText = "`" + splittedParam[0] + "` >= " + rangeSplitted[0]
		}

		ranges = append(ranges, rangeText)
	}

	for i, rangeText := range ranges {
		if i != 0 {
			statement += " AND " + rangeText
		} else {
			statement += rangeText
		}
	}

	return
}

func GetPagingStatement(ctx *fiber.Ctx) (offset int, limit int, err error) {
	page := 1
	offset = 0 // default offset
	limit = 10 // default limit

	paramValue := ctx.Query("page")
	if paramValue == "" {
		return
	}

	// params[0] is page
	// params[1] is limit
	params := strings.Split(paramValue, ":")

	if params[0] != "" {
		page, err = strconv.Atoi(params[0])
	}

	if params[1] != "" {
		limit, err = strconv.Atoi(params[1])
	}

	if page != 1 {
		offset = (page - 1) * limit
	}

	return
}

func ComposeWithURLQuery(ctx *fiber.Ctx, tx *gorm.DB, model interface{}) (err error) {
	criteriaQueryStatement, err := GetCriteriaStatement(ctx, model)
	if err != nil {
		return
	}

	tx.Where(criteriaQueryStatement)

	offset, limit, err := GetPagingStatement(ctx)
	if err != nil {
		return
	}

	tx.Offset(offset).Limit(limit)

	orderQueryMap, err := GetOrderMap(ctx)
	for key, val := range orderQueryMap {
		order := fmt.Sprintf("%s %s", key, val)
		tx.Order(order)
	}

	return
}

func GetQuery(ctx *fiber.Ctx, tableName string, model interface{}, isForMetaPurpose bool) (query string, err error) {
	query = fmt.Sprintf("SELECT * FROM %s", tableName)

	criteria, err := GetCriteriaStatement(ctx, model)
	if err != nil {
		return
	}
	orderQuery, err := GetOrderStatement(ctx)
	if err != nil {
		return
	}
	offset, limit, err := GetPagingStatement(ctx)
	if err != nil {
		return
	}

	if criteria != "" {
		query = query + fmt.Sprintf(" WHERE %s", criteria)
	}

	if orderQuery != "" {
		query = query + fmt.Sprintf(" %s", orderQuery)
	}

	if limit != 0 && !isForMetaPurpose {
		query = query + fmt.Sprintf(" OFFSET %d LIMIT %d", offset, limit)
	}

	return
}

func GetLimitOffsetQueryPart(page, limit *int) (stmt string) {
	if page != nil {
		stmt += fmt.Sprintf(" offset %v", getPaginationOffset(page, limit))
	}
	if limit != nil {
		stmt += fmt.Sprintf(" limit %v", *limit)
	}

	return
}

func getPaginationOffset(page, limit *int) int {
	p := *page
	l := 10

	if limit != nil {
		l = *limit
	}

	return p*l - l
}
