package ai

import (
	rulepkg "github.com/actiontech/sqle/sqle/driver/mysql/rule"
	util "github.com/actiontech/sqle/sqle/driver/mysql/rule/ai/util"
	driverV2 "github.com/actiontech/sqle/sqle/driver/v2"
	"github.com/actiontech/sqle/sqle/pkg/params"
	"github.com/pingcap/parser/ast"
)

const (
	SQLE00090 = "SQLE00090"
)

func init() {
	rh := rulepkg.RuleHandler{
		Rule: driverV2.Rule{
			Name:       SQLE00090,
			Desc:       "在 MySQL 中, 建议使用UNION ALL替代UNION",
			Annotation: "union会对结果集进行去重，union all只是简单的将两个结果合并后就返回，从效率上看，union all 要比union快很多；如果合并的两个结果集中允许包含重复数据的话，建议开启此规则，使用union all替代union",
			Level:      driverV2.RuleLevelNotice,
			Category:   rulepkg.RuleTypeDMLConvention,
			Params:     params.Params{},
		},
		Message:      "在 MySQL 中, 建议使用UNION ALL替代UNION",
		AllowOffline: true,
		Func:         RuleSQLE00090,
	}
	rulepkg.RuleHandlers = append(rulepkg.RuleHandlers, rh)
	rulepkg.RuleHandlerMap[rh.Rule.Name] = rh
}

/*
==== Prompt start ====
在 MySQL 中，您应该检查 SQL 是否违反了规则(SQLE00090): "在 MySQL 中，建议使用UNION ALL替代UNION."
您应遵循以下逻辑：
1. 对于 dml 语句:
  1、检查句子中是否存在SELECT子句，如果存在，则进入下一步检查。
  2、检查SELECT子句中是否存在UNION语法节点，如果存在，但不是UNION ALL，则报告违反规则。

1. 对于UNION...语句, 对于其中的所有SELECT子句进行与SELECT语句相同的检查。
==== Prompt end ====
*/

// ==== Rule code start ====
// 规则函数实现开始
func RuleSQLE00090(input *rulepkg.RuleHandlerInput) error {
	// 子查询中
	for _, subquery := range util.GetSubquery(input.Node) {
		if unionStmt, ok := subquery.Query.(*ast.UnionStmt); ok {
			if util.CheckUnionNotAll(unionStmt) {
				rulepkg.AddResult(input.Res, input.Rule, SQLE00090)
				return nil
			}
		}
	}

	switch stmt := input.Node.(type) {
	case *ast.UnionStmt, *ast.SelectStmt:
		if util.CheckUnionNotAll(stmt) {
			rulepkg.AddResult(input.Res, input.Rule, SQLE00090)
			return nil
		}
	case *ast.UpdateStmt, *ast.DeleteStmt:
		for _, selectStmt := range util.GetSelectStmt(stmt) {
			if selectStmt.From != nil {
				if util.CheckUnionNotAll(selectStmt) {
					rulepkg.AddResult(input.Res, input.Rule, SQLE00090)
					return nil
				}
			}
		}
	case *ast.InsertStmt:
		if stmt.Select != nil {
			if unionStmt, ok := stmt.Select.(*ast.UnionStmt); ok {
				if util.CheckUnionNotAll(unionStmt) {
					rulepkg.AddResult(input.Res, input.Rule, SQLE00090)
					return nil
				}
			}
		}
	}
	return nil
}

// 规则函数实现结束
// ==== Rule code end ====
