package base

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/context"
)

// 查询条件
type QueryOptions struct {
	Objid   string
	Name    string
	Type    string
	Time    string
	Starttm string
	Endtm   string
	Status  string
	Limit   int
	Offset  int
	SortBy  string
	Order   string
}

type baseParamParser struct {
	Limit    int
	MaxLimit int
}

func NewBaseParamParser(limit, maxLimit int) *baseParamParser {
	return &baseParamParser{
		Limit:    limit,
		MaxLimit: maxLimit,
	}
}
func (p *baseParamParser) BaseParam(input *context.BeegoInput) *QueryOptions {
	qo := new(QueryOptions)

	inObjid := input.Query("objid")
	if inObjid != "" {
		qo.Objid = inObjid
	}

	inName := input.Query("name")
	if inName != "" {
		qo.Name = inName
	}

	inType := input.Query("type")
	if inType != "" {
		qo.Type = inType
	}

	inTime := input.Query("time")
	if inTime != "" {
		//将时间字符串转换为int64类型
		int64Time, err := strconv.ParseInt(inTime, 10, 64)
		if err != nil {
			qo.Time = ""
		} else {
			ti := time.Unix(int64Time, 0)
			qo.Time = ti.Format("2006-01-02 15:04:05")
		}
	}

	inStatus := input.Query("status")
	if inStatus != "" {
		qo.Status = inStatus
	}

	inSortby := input.Query("sortby")
	if inSortby != "" {
		inOrder := input.Query("order")
		if inOrder != "" && inOrder == "desc" {
			qo.SortBy = "-" + inSortby
		} else {
			qo.SortBy = inSortby
		}
	}

	startTm := input.Query("starttm")
	if startTm != "" {
		//将时间字符串转换为int64类型
		int64Time, err := strconv.ParseInt(startTm, 10, 64)
		if err != nil {
			qo.Starttm = ""
		} else {
			ti := time.Unix(int64Time, 0)
			qo.Starttm = ti.Format("2006-01-02 15:04:05")
		}
	}

	endTm := input.Query("endtm")
	if endTm != "" {
		//将时间字符串转换为int64类型
		int64Time, err := strconv.ParseInt(endTm, 10, 64)
		if err != nil {
			qo.Endtm = ""
		} else {
			ti := time.Unix(int64Time, 0)
			qo.Endtm = ti.Format("2006-01-02 15:04:05")
		}
	}

	inLimit := input.Query("limit")
	qo.Limit = p.Limit
	if inLimit != "" {
		qo.Limit, _ = strconv.Atoi(inLimit)
		if qo.Limit < 0 {
			qo.Limit = p.MaxLimit
		} else if qo.Limit == 0 {
			qo.Limit = p.Limit
		}
	}

	inOffset := input.Query("offset")
	if inOffset != "" {
		qo.Offset, _ = strconv.Atoi(inOffset)
		if qo.Offset <= 0 {
			qo.Offset = 0
		}

	}
	return qo
}

// 获取分页信息
func (opt *QueryOptions) GetPageInfo(total int) *Page {
	return NewPage(total, opt.Offset, opt.Limit)
}
