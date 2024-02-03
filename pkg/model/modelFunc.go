package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-blog/pkg/vo"

	"gorm.io/gorm"
)

type PublicFunc interface {
	UpdateFunc(model interface{}, cond, values map[string]interface{}, globalUpdate bool) error
	HasExist(model interface{}, cond map[string]interface{}) bool
	QueryFunc(model interface{}, cond map[string]interface{}, isList bool) (interface{}, error)
	ParseJSONToStruct(jsonStr interface{}, resultStruct interface{}) error
	DeleteFunc(model interface{}, cond map[string]interface{}) error
}

type QueryStrategy interface {
	Execute(*gorm.DB) (QueryResult, error)
}

type OFunc struct {
	l  *log.Helper
	db *gorm.DB
}

type QueryResult interface{}

type ListQueryStrategy struct{}

type SingleQueryStrategy struct{}

func (s *ListQueryStrategy) Execute(query *gorm.DB) (QueryResult, error) {
	var list []map[string]interface{}
	if err := query.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *SingleQueryStrategy) Execute(query *gorm.DB) (QueryResult, error) {
	var data map[string]interface{}
	if err := query.First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewOFunc(l *log.Helper, db *gorm.DB) PublicFunc {
	return &OFunc{
		l:  l,
		db: db,
	}
}

// UpdateFunc :dev
func (o *OFunc) UpdateFunc(model interface{}, cond, values map[string]interface{}, globalUpdate bool) error {
	var err error
	if len(values) == 0 {
		err = errors.New("values map is empty")
	}

	updateQuery := o.db.Model(model)

	// Allow global updates
	if globalUpdate {
		updateQuery = updateQuery.Session(&gorm.Session{AllowGlobalUpdate: true})
	}

	if len(cond) != 0 {
		for cd, va := range cond {
			c := fmt.Sprintf("%s = ?", cd)
			updateQuery = updateQuery.Where(c, va)
		}
	}

	for key, value := range values {
		updateQuery = updateQuery.Update(key, value)
	}

	if e := updateQuery.Error; e != nil {
		err = e
	}
	return err
}

// HasExist :dev Determine whether it exists according to the conditions
func (o *OFunc) HasExist(model interface{}, cond map[string]interface{}) bool {
	var (
		res bool
		err error
	)
	if len(cond) == 0 {
		err = errors.New(vo.LIST_EMPTY)
		res = false
		return res
	}
	hasQuery := o.db.Model(&model)
	for cd, va := range cond {
		cd := fmt.Sprintf("%s = ?", cd)
		hasQuery = hasQuery.Where(cd, va)
	}
	var da map[string]interface{}
	if err = hasQuery.First(&da).Error; err != nil {
		o.l.Log(log.LevelError, err)
		res = false
		return res
	}
	res = true
	return res
}

// QueryFunc :dev search based on conditions
func (o *OFunc) QueryFunc(model interface{}, cond map[string]interface{}, isList bool) (interface{}, error) {
	var (
		data interface{}
		err  error
		rule = "create_time DESC"
	)

	query := o.db.Model(model).Order(rule)

	if len(cond) != 0 {
		for cd, va := range cond {
			c := fmt.Sprintf("%s = ?", cd)
			query = query.Where(c, va)
		}
	}

	var strategy QueryStrategy
	if isList {
		strategy = &ListQueryStrategy{}
	} else {
		strategy = &SingleQueryStrategy{}
	}

	data, err = strategy.Execute(query)
	if err != nil {
		o.l.Errorf("Error querying data: %s", err)
	}

	return data, err
}

// ParseJSONToStruct :dev parse the string into the struct
func (o *OFunc) ParseJSONToStruct(jsonStr interface{}, resultStruct interface{}) error {
	f, _ := json.Marshal(jsonStr)
	err := json.Unmarshal(f, resultStruct)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %s", err)
	}
	return nil
}

// DeleteFunc :dev Incoming condition deletion
func (o *OFunc) DeleteFunc(model interface{}, cond map[string]interface{}) error {
	var err error
	if len(cond) == 0 {
		err = errors.New(vo.LIST_EMPTY)
		return err
	}
	deleteQuery := o.db
	for cd, va := range cond {
		c := fmt.Sprintf("%s = ?", cd)
		deleteQuery = deleteQuery.Where(c, va)
	}
	err = deleteQuery.Delete(&model).Error

	return err
}
