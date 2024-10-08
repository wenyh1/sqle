package ai

import (
	"fmt"

	rulepkg "github.com/actiontech/sqle/sqle/driver/mysql/rule"
	util "github.com/actiontech/sqle/sqle/driver/mysql/rule/ai/util"
	driverV2 "github.com/actiontech/sqle/sqle/driver/v2"
	"github.com/actiontech/sqle/sqle/pkg/params"
	"github.com/pingcap/parser/ast"
)

const (
	SQLE00087 = "SQLE00087"
)

func init() {
	rh := rulepkg.RuleHandler{
		Rule: driverV2.Rule{
			Name:       SQLE00087,
			Desc:       "在 MySQL 中, 避免WHERE条件内IN语句中的参数值个数过多",
			Annotation: "当IN值过多时，有可能会导致查询进行全表扫描，使得数据库性能急剧下降。",
			Level:      driverV2.RuleLevelWarn,
			Category:   rulepkg.RuleTypeDMLConvention,
			Params: params.Params{
				&params.Param{
					Key:   rulepkg.DefaultSingleParamKeyName,
					Value: "50",
					Desc:  "IN的参数值个数",
					Type:  params.ParamTypeInt,
				},
			},
		},
		Message:      "在 MySQL 中, 避免WHERE条件内IN语句中的参数值个数过多",
		AllowOffline: true,
		Func:         RuleSQLE00087,
	}
	rulepkg.RuleHandlers = append(rulepkg.RuleHandlers, rh)
	rulepkg.RuleHandlerMap[rh.Rule.Name] = rh
}

/*
==== Prompt start ====
在 MySQL 中，您应该检查 SQL 是否违反了规则(SQLE00087): "在 MySQL 中，避免WHERE条件内IN语句中的参数值个数过多.默认参数描述: IN的参数值个数, 默认参数值: 50"
您应遵循以下逻辑：
1、检查SELECT语句中是否存在WHERE子句，如果存在，则进一步检查。
2、检查WHERE子句中是否存在IN条件节点，如果存在，则进一步检查。
3、检查IN条件节点中的参数是否为常量集合，如果是，则继续检查。
4、检查IN条件节点的参数个数是否超过当前规则的阈值，如果超过阈值，则触发违反规则。

1. 对于UNION语句, 递归检查其中的所有SELECT子句，执行与SELECT语句相同的检查。

1、检查INSERT INTO语句中是否存在SELECT子句，如果存在，则进一步检查。
2、检查SELECT子句中是否存在WHERE子句，如果存在，则进一步检查。
3、检查WHERE子句中是否存在IN条件节点，如果存在，则进一步检查。
4、检查IN条件节点中的参数是否为常量集合，如果是，则继续检查。
5、检查IN条件节点的参数个数是否超过当前规则的阈值，如果超过阈值，则触发违反规则。

1、检查UPDATE语句中是否存在WHERE子句，如果存在，则进一步检查。
2、检查WHERE子句中是否存在IN条件节点，如果存在，则进一步检查。
3、检查IN条件节点中的参数是否为常量集合，如果是，则继续检查。
4、检查IN条件节点的参数个数是否超过当前规则的阈值，如果超过阈值，则触发违反规则。

1、检查DELETE语句中是否存在WHERE子句，如果存在，则进一步检查。
2、检查WHERE子句中是否存在IN条件节点，如果存在，则进一步检查。
3、检查IN条件节点中的参数是否为常量集合，如果是，则继续检查。
4、检查IN条件节点的参数个数是否超过当前规则的阈值，如果超过阈值，则触发违反规则。
==== Prompt end ====
*/

// ==== Rule code start ====
// 规则函数实现开始
func RuleSQLE00087(input *rulepkg.RuleHandlerInput) error {
	// 获取规则阈值
	param := input.Rule.Params.GetParam(rulepkg.DefaultSingleParamKeyName)
	if param == nil {
		return fmt.Errorf("param %s not found", rulepkg.DefaultSingleParamKeyName)
	}
	threshold := param.Int()

	// 扫描 WHERE 子句中的 IN 表达式并检查参数数量
	scanWhereClause := func(whereClause ast.ExprNode) bool {
		foundViolation := false
		util.ScanWhereStmt(func(expr ast.ExprNode) bool {
			switch x := expr.(type) {
			case *ast.PatternInExpr:
				inQueryParamActualNumber := len(x.List)
				if inQueryParamActualNumber > threshold {
					foundViolation = true
					return true
				}
			}
			return false
		}, whereClause)
		return foundViolation
	}

	// 处理 DML 语句 (INSERT, UPDATE, DELETE, SELECT, UNION)
	handleDMLStmt := func(node ast.Node) bool {
		whereClauses := util.GetWhereExprFromDMLStmt(node)
		for _, clause := range whereClauses {
			if scanWhereClause(clause) {
				return true
			}
		}
		return false
	}

	if handleDMLStmt(input.Node) {
		rulepkg.AddResult(input.Res, input.Rule, SQLE00087)
		return nil
	}
	return nil
}

// 规则函数实现结束
// ==== Rule code end ====
