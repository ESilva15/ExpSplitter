package repo

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	dec "github.com/shopspring/decimal"
)

func decimalToNumeric(d dec.Decimal) (pgtype.Numeric, error) {
	var num pgtype.Numeric
	err := num.Scan(d.Truncate(2).String())
	if err != nil {
		return pgtype.Numeric{}, err
	}
	num.Valid = true

	return num, nil
}

func boolToPgBool(v bool) (pgtype.Bool, error) {
	var b pgtype.Bool
	err := b.Scan(v)
	if err != nil {
		return pgtype.Bool{}, err
	}
	b.Valid = true

	return b, nil
}

func timeToTimestamp(t *time.Time) (pgtype.Timestamp, error) {
	var ts pgtype.Timestamp

	if t == nil {
		ts.Valid = false
	} else {
		err := ts.Scan(*t)
		if err != nil {
			return pgtype.Timestamp{}, err
		}
	}

	return ts, nil
}

func pgNumericToDecimal(n pgtype.Numeric) dec.Decimal {
	return dec.NewFromBigInt(n.Int, n.Exp)
}

func pgBoolToBool(n pgtype.Bool) bool {
	return n.Bool
}

func pgTimestampToTime(ts pgtype.Timestamp) time.Time {
	return ts.Time
}

func pgTextToString(s pgtype.Text) string {
	return s.String
}
