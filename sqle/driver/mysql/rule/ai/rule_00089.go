package ai

import (
	rulepkg "github.com/actiontech/sqle/sqle/driver/mysql/rule"
	util "github.com/actiontech/sqle/sqle/driver/mysql/rule/ai/util"
	driverV2 "github.com/actiontech/sqle/sqle/driver/v2"
	"github.com/actiontech/sqle/sqle/pkg/params"
	"github.com/pingcap/parser/ast"
)

const (
	SQLE00089 = "SQLE00089"
)

func init() {
	rh := rulepkg.RuleHandler{
		Rule: driverV2.Rule{
			Name:       SQLE00089,
			Desc:       "在 MySQL 中, 禁止INSERT ... SELECT",
			Annotation: "使用 INSERT ... SELECT 在默认事务隔离级别下，可能会导致对查询的表施加表级锁",
			Level:      driverV2.RuleLevelWarn,
			Category:   rulepkg.RuleTypeDMLConvention,
			Params:     params.Params{},
		},
		Message:      "在 MySQL 中, 禁止INSERT ... SELECT",
		AllowOffline: true,
		Func:         RuleSQLE00089,
	}
	rulepkg.RuleHandlers = append(rulepkg.RuleHandlers, rh)
	rulepkg.RuleHandlerMap[rh.Rule.Name] = rh
}

/*
==== Prompt start ====
在 MySQL 中，您应该检查 SQL 是否违反了规则(SQLE00089): "在 MySQL 中，禁止INSERT ... SELECT."
您应遵循以下逻辑：
1. 对于 insert... 语句，
  1、解析语法树，检查INSERT INTO语句中是否包含SELECT子句。
  2、如果SELECT子句中包含FROM子句，则报告违反规则。
==== Prompt end ====
*/

// ==== Rule code start ====
func RuleSQLE00089(input *rulepkg.RuleHandlerInput) error {
	switch stmt := input.Node.(type) {
	case *ast.InsertStmt:
		// 对于 "INSERT..." 语句
		for _, selectStmt := range util.GetSelectStmt(stmt) {
			// 检查 SELECT 子句中是否包含 FROM 子句
			if selectStmt.From != nil && selectStmt.From.TableRefs != nil {
				// 违反规则，报告结果
				rulepkg.AddResult(input.Res, input.Rule, SQLE00089)
				return nil
			}
		}
	}
	return nil
}

// ==== Rule code end ====
