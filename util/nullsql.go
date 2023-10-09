package util

import (
	"database/sql"
	"simplebank/pb"
	"time"

	"github.com/shopspring/decimal"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// LoadConfig reads configuration from file or environment variables.
func NullString() sql.NullString {
	return sql.NullString{String: "", Valid: false}
}

func NullBool() sql.NullBool {
	return sql.NullBool(sql.NullBool{Bool: false, Valid: false})
}

func NullInt32() sql.NullInt32 {
	return sql.NullInt32{Int32: 0, Valid: false}
}

func NullInt64() sql.NullInt64 {
	return sql.NullInt64{Int64: 0, Valid: false}
}

func NullTime() sql.NullTime {
	return sql.NullTime{Time: time.Time{}, Valid: false}
}

func NullFloat64() sql.NullFloat64 {
	return sql.NullFloat64{Float64: 0, Valid: false}
}

func SetNullString(str string) sql.NullString {
	return sql.NullString{String: str, Valid: true}
}

func SetNullBool(b bool) sql.NullBool {
	return sql.NullBool(sql.NullBool{Bool: b, Valid: true})
}

func SetNullInt16(i int16) sql.NullInt16 {
	return sql.NullInt16{Int16: i, Valid: true}
}

func SetNullInt32(i int32) sql.NullInt32 {
	return sql.NullInt32{Int32: i, Valid: true}
}

func SetNullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: true}
}

func SetNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: true}
}

func SetNullFloat64(f float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: f, Valid: true}
}

func SetDecimal(f string) decimal.Decimal {
	d, e := decimal.NewFromString(f)
	if e == nil {
		return d
	} else {
		return decimal.NewFromInt(0)
	}
}

func SetNullDecimal(f string) decimal.NullDecimal {
	return decimal.NullDecimal{Decimal: SetDecimal(f), Valid: true}
}

func NullString2Proto(v sql.NullString) *pb.NullString {
	return &pb.NullString{
		Value: v.String,
		Valid: v.Valid,
	}
}

func NullBool2Proto(v sql.NullBool) *pb.NullBool {
	return &pb.NullBool{
		Value: v.Bool,
		Valid: v.Valid,
	}
}

func NullInt642Proto(v sql.NullInt64) *pb.NullInt64 {
	return &pb.NullInt64{
		Value: v.Int64,
		Valid: v.Valid,
	}
}

func NullTime2Proto(v sql.NullTime) *pb.NullTime {
	return &pb.NullTime{
		Value: timestamppb.New(v.Time),
		Valid: v.Valid,
	}
}

func NullFloat642Proto(v sql.NullFloat64) *pb.NullFloat64 {
	return &pb.NullFloat64{
		Value: v.Float64,
		Valid: v.Valid,
	}
}

func NullDecimal2Proto(v decimal.NullDecimal) *pb.Decimal {
	return &pb.Decimal{
		Value: v.Decimal.String(),
		Valid: v.Valid,
	}
}

func Decimal2Proto(v decimal.Decimal) *pb.Decimal {
	return &pb.Decimal{
		Value: v.String(),
		Valid: true,
	}
}

func NullProto2String(v *pb.NullString) sql.NullString {
	return sql.NullString{
		String: v.Value,
		Valid:  v.Valid,
	}
}

func NullProto2Decimal(v *pb.Decimal) decimal.NullDecimal {
	return decimal.NullDecimal{Decimal: SetDecimal(v.Value), Valid: v.Valid}
}

func Proto2Decimal(v *pb.Decimal) decimal.Decimal {
	return SetDecimal(v.Value)
}

func NullProto2Bool(v *pb.NullBool) sql.NullBool {
	return sql.NullBool{
		Bool:  v.Value,
		Valid: v.Valid,
	}
}

func NullProto2Int64(v *pb.NullInt64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: v.Value,
		Valid: v.Valid,
	}
}

func NullProto2Time(v *pb.NullTime) sql.NullTime {
	return sql.NullTime{
		Time:  v.Value.AsTime(),
		Valid: v.Valid,
	}
}

func NullProto2Float64(v *pb.NullFloat64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: v.Value,
		Valid:   v.Valid,
	}
}
