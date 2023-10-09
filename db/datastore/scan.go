package db

import (
	"time"

	"github.com/google/uuid"
)

type MetalScanner struct {
	valid bool
	value interface{}
}

func (scanner *MetalScanner) getBytes(src interface{}) []byte {
	if a, ok := src.([]uint8); ok {
		return a
	}
	return nil
}

func (scanner *MetalScanner) Scan(src interface{}) error {

	switch src.(type) {
	case int64:
		if value, ok := src.(int64); ok {
			scanner.value = value
			scanner.valid = true
		} else {
			scanner.value = 0
			scanner.valid = false
		}
	case float64:
		if value, ok := src.(float64); ok {
			scanner.value = value
			scanner.valid = true
		}
	case bool:
		if value, ok := src.(bool); ok {
			scanner.value = value
			scanner.valid = true
		}
	case string:
		//log.Println("[=", scanner.getBytes(src), "=] [", src, "]")
		if src != nil {
			scanner.value = (src)
			scanner.valid = true
		} else {
			scanner.value = ""
			scanner.valid = false
		}

	case []byte:
		value := scanner.getBytes(src)
		scanner.value = value
		scanner.valid = true
	case time.Time:
		if value, ok := src.(time.Time); ok {
			scanner.value = value
			scanner.valid = true
		}
	case nil:
		//log.Println("===[", scanner.value, "]=")
		scanner.value = nil
		scanner.valid = true
	default:
		value := scanner.getBytes(src)
		scanner.value = string(value)
		scanner.valid = true
	}
	return nil
}

func toInt(fld interface{}) int {
	switch fld.(type) {
	case int64:
		return int(fld.(int64))
	default:
		return 0
	}
}

func toInt64(fld interface{}) int64 {
	switch fld.(type) {
	case int64:
		return fld.(int64)
	default:
		return 0
	}
}

func toTime(fld interface{}) time.Time {
	switch fld.(type) {
	case time.Time:
		return fld.(time.Time)
	default:
		return time.Time{}
	}
}

func toBool(fld interface{}) bool {
	switch fld.(type) {
	case bool:
		return fld.(bool)
	default:
		return false
	}
}

func tofloat64(fld interface{}) float64 {
	switch fld.(type) {
	case float64:
		return fld.(float64)
	default:
		return 0
	}
}

func toUUID(fld interface{}) uuid.UUID {
	switch fld.(type) {
	case []byte:
		//suuid := string(fld.([]uint8)[:])
		_uuid, er := uuid.ParseBytes(fld.([]uint8))
		if er != nil {
			return uuid.UUID{}
		}
		return _uuid
	default:
		return uuid.UUID{}
	}
}

func toString(fld interface{}) string {
	switch fld.(type) {
	case string:
		return fld.(string)
	case []byte:
		//suuid := string(fld.([]uint8)[:])
		return string(fld.([]uint8))
	default:
		return ""
	}
}
