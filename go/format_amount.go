package handlers

import (
	"fmt"
	"strconv"
	"strings"
)

// It trims trailing 0s.
func formatAmount(num interface{}, decimal int) string {
	var ns string
	var no int64
	var err error
	switch t := num.(type) {
	case string:
		ns = t
		no, err = strconv.ParseInt(t, 10, 64)
		if err != nil {
			return ""
		}
	case int:
		ns = strconv.FormatInt(int64(t), 10)
		no = int64(t)
	case int64:
		ns = strconv.FormatInt(t, 10)
		no = t
	default:
		return ""
	}
	if ns == "0" {
		return "0"
	}

	sign := ""
	if no < 0 {
		sign = "-"
	}
	absstr := strings.TrimPrefix(ns, "-")
	monetary := "0"
	fraction := ""
	if len(absstr) > decimal {
		integer := absstr[0 : len(absstr)-decimal]
		fraction = absstr[len(absstr)-decimal:]
		i1 := len(integer) % 3
		m := []string{}
		if i1 > 0 {
			m = append(m, integer[0:i1])
		}
		for i2 := 0; i2 < (len(integer)-i1)/3; i2++ {
			m = append(m, integer[i1+i2*3:i1+i2*3+3])
		}
		monetary = strings.Join(m, ",")
	} else if len(absstr) == decimal {
		fraction = absstr
	} else {
		fraction = strings.Repeat("0", decimal-len(absstr)) + absstr
	}
	fraction = strings.TrimRight(fraction, "0")
	p := "" // decimal point
	if len(fraction) > 0 {
		p = "."
	}
	return fmt.Sprintf("%s%s%s%s", sign, monetary, p, fraction)
}

