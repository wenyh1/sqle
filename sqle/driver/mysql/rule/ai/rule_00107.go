package ai

import (
	"fmt"
	"strconv"

	rulepkg "github.com/actiontech/sqle/sqle/driver/mysql/rule"
	driverV2 "github.com/actiontech/sqle/sqle/driver/v2"
	"github.com/actiontech/sqle/sqle/pkg/params"
	"github.com/pingcap/parser/ast"
)

const (
	SQLE00107 = "SQLE00107"
)

func init() {
	rh := rulepkg.RuleHandler{
		Rule: driverV2.Rule{
			Name:       SQLE00107,
			Desc:       "在 MySQL 中, 建议将过长的SQL分解成几个简单的SQL",
			Annotation: "过长的SQL可读性较差，难以维护，且容易引发性能问题。",
			Level:      driverV2.RuleLevelNotice,
			Category:   rulepkg.RuleTypeDMLConvention,
			Params: params.Params{
				&params.Param{
					Key:   rulepkg.DefaultSingleParamKeyName,
					Value: "1024",
					Desc:  "句子长度限制",
					Type:  params.ParamTypeInt,
				},
			},
		},
		Message:      "在 MySQL 中, 建议将过长的SQL分解成几个简单的SQL",
		AllowOffline: true,
		Func:         RuleSQLE00107,
	}
	rulepkg.RuleHandlers = append(rulepkg.RuleHandlers, rh)
	rulepkg.RuleHandlerMap[rh.Rule.Name] = rh
}

/*
==== Prompt start ====
在 MySQL 中，您应该检查 SQL 是否违反了规则(SQLE00107): "在 MySQL 中，建议将过长的SQL分解成几个简单的SQL.默认参数描述: 句子长度限制, 默认参数值: 1024 "
您应遵循以下逻辑：
1. 对于所有DML语句，计算SQL语句的字符串长度，如果大于等于阈值，则报告违反规则。

==== Prompt end ====
*/

// ==== Rule code start ====
// 规则函数实现开始
func RuleSQLE00107(input *rulepkg.RuleHandlerInput) error {
	// 确保输入节点是DML语句
	if _, ok := input.Node.(ast.DMLNode); !ok {
		return nil
	}

	// 步骤1: 检查SQL语句的字符串长度
	param := input.Rule.Params.GetParam(rulepkg.DefaultSingleParamKeyName)
	if param == nil {
		return fmt.Errorf("param %s not found", rulepkg.DefaultSingleParamKeyName)
	}

	lengthThreshold, err := strconv.Atoi(param.Value)
	if err != nil {
		return fmt.Errorf("param must be an integer, got '%s'", param.Value)
	}

	sqlText := input.Node.Text()
	sqlSize := len(sqlText)
	if sqlText[sqlSize-1] == ';' {
		if sqlSize-1 >= lengthThreshold {
			rulepkg.AddResult(input.Res, input.Rule, SQLE00107)
		}
	} else {
		if sqlSize >= lengthThreshold {
			rulepkg.AddResult(input.Res, input.Rule, SQLE00107)
		}
	}

	return nil
}

// 规则函数实现结束
// ==== Rule code end ====
