package behavior

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lapengo/lapengo-go/internal/helper"
	"gorm.io/gorm"
	"log"
)

type BaseBehavior struct {
	DB        *gorm.DB
	Validate  *validator.Validate
	TableName string
}

func NewBaseBehavior() *BaseBehavior {
	return &BaseBehavior{}
}

func (b *BaseBehavior) SimpleGetAllNew(ctx *fiber.Ctx, models interface{}, model interface{}) (err error) {

	return
}

func (b *BaseBehavior) SimpleGetAll(ctx *fiber.Ctx, models interface{}, model interface{}) (err error) {
	query, err := helper.GetQuery(ctx, b.TableName, model, false)

	if err != nil {
		log.Print(err.Error())
		return
	}

	b.DB.Raw(query).Find(models)

	return
}

func (b *BaseBehavior) SimpleGetAllWithPreload(ctx *fiber.Ctx, viewModels interface{}, preloads []string) (err error) {
	var newDBInstance gorm.DB = *b.DB
	// var newDBPtr *gorm.DB = &newDBInstance
	var newDBPtr = &newDBInstance

	newDBPtr = newDBPtr.Table(b.TableName)

	// set preloads
	for _, preload := range preloads {
		newDBPtr = newDBPtr.Preload(preload)
	}

	// set criteria
	criterias, err := helper.GetCriteria(ctx)
	for _, criteria := range criterias {
		newDBPtr = newDBPtr.Where(criteria)
	}

	// set search
	searchStatement, err := helper.GetSearchStatement(ctx)
	if searchStatement != "" {
		newDBPtr = newDBPtr.Where(searchStatement)
	}

	// set range
	rangeStatement, err := helper.GetRangeStatement(ctx)
	if rangeStatement != "" {
		newDBPtr = newDBPtr.Where(rangeStatement)
	}

	// set paging
	offset, limit, err := helper.GetPagingStatement(ctx)
	if limit != 0 {
		newDBPtr = newDBPtr.Offset(offset).Limit(limit)
	} else {
		newDBPtr = newDBPtr.Offset(offset)
	}

	// set orders
	orderStatement, err := helper.GetOrderString(ctx)
	if orderStatement != "" {
		newDBPtr = newDBPtr.Order(orderStatement + " ASC")
	}

	newDBPtr.Find(viewModels)

	return
}

func (b *BaseBehavior) SimpleGetByID(ctx *fiber.Ctx, model interface{}, ID interface{}) (err error) {
	err = b.DB.Table(b.TableName).Find(model, ID).Error

	return
}

func (b *BaseBehavior) SimpleCreate(ctx *fiber.Ctx, model interface{}, initialModel interface{}) (err error) {
	if err = ctx.BodyParser(&model); err != nil {
		log.Print(err.Error())
		err = errors.New("BAD REQUEST")
	}

	err = b.Validate.Struct(model)
	if err != nil {
		log.Print(err.Error())
		return
	}

	if initialModel != nil {
		err = b.DB.Table(b.TableName).Create(initialModel).Error
		model = initialModel
	} else {
		err = b.DB.Table(b.TableName).Create(model).Error
	}

	return
}

func (b *BaseBehavior) SimpleUpdate(ctx *fiber.Ctx, model interface{}, newData interface{}, ID interface{}) (err error) {

	err = b.DB.Table(b.TableName).Find(model, ID).Error
	if err != nil {
		return
	}
	err = b.DB.Table(b.TableName).Model(model).Updates(newData).Error

	return
}

func (b *BaseBehavior) SimpleDelete(ctx *fiber.Ctx, model interface{}, ID interface{}) (err error) {
	err = b.DB.Table(b.TableName).Delete(&model, ID).Error

	return
}
