package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"autocrud/pkg/config"
	"autocrud/pkg/helper"

	"github.com/google/uuid"
)

func getUpdateData(data map[string]any, body io.ReadCloser, document string, isCreate bool) error {
	requestBody := map[string]any{}
	_ = json.NewDecoder(body).Decode(&requestBody)

	for field, fType := range config.Config.Application.Documents[document].Fields {
		val := fmt.Sprint(requestBody[field])
		if requestBody[field] == nil || val == "" {
			if isCreate && fType.Required {
				return errors.New(fmt.Sprintf("%s is required", field))
			}
		} else {
			value, err := getVarTypeFunc(fType.Type)(field, val, fType)
			if err != nil {
				return err
			}
			if value != nil {
				data[field] = value
			}
		}
	}
	return nil
}

func getVarTypeFunc(fType string) func(string, string, config.Field) (any, error) {
	if fType == "uuid" {
		return getUUID
	} else if fType == "int" {
		return getInt
	} else if fType == "float" {
		return getFloat
	} else if fType == "bool" {
		return getBoolean
	} else {
		return getString
	}
}

func getUUID(key string, val string, fType config.Field) (any, error) {
	value, err := uuid.Parse(val)
	return value, helper.ModifyError(err, fmt.Sprintf("%s is invalid", key))
}

func getString(key string, value string, fType config.Field) (any, error) {
	if fType.Max > 0 && len(value) > fType.Max {
		return nil, errors.New(fmt.Sprintf("%s's length is more than max length %d", key, fType.Max))
	}

	if fType.Min > 0 && len(value) < fType.Min {
		return nil, errors.New(fmt.Sprintf("%s's length is less than min length %d", key, fType.Min))
	}
	return value, nil
}

func getInt(key string, value string, fType config.Field) (any, error) {
	intVar, err := strconv.Atoi(value)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s is invalid", key))
	}
	if fType.Max > 0 && intVar > fType.Max {
		return nil, errors.New(fmt.Sprintf("max %s allowed is %d", key, fType.Max))
	}

	if fType.Min > 0 && intVar < fType.Min {
		return nil, errors.New(fmt.Sprintf("min %s allowed is %d", key, fType.Min))
	}
	return intVar, nil
}

func getBoolean(key string, value string, fType config.Field) (any, error) {
	boolValue, err := strconv.ParseBool(value)
	return boolValue, helper.ModifyError(err, fmt.Sprintf("%s is invalid", key))
}

func getFloat(key string, value string, fType config.Field) (any, error) {
	floatVar, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s is invalid", key))
	}
	if fType.Max > 0 && floatVar > float64(fType.Max) {
		return nil, errors.New(fmt.Sprintf("max %s allowed is %d", key, fType.Max))
	}

	if fType.Min > 0 && floatVar < float64(fType.Min) {
		return nil, errors.New(fmt.Sprintf("min %s allowed is %d", key, fType.Min))
	}
	return floatVar, nil
}
