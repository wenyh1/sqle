package mysql

import (
	"testing"

	rulepkg "github.com/actiontech/sqle/sqle/driver/mysql/rule"
	"github.com/actiontech/sqle/sqle/driver/mysql/rule/ai"
	"github.com/actiontech/sqle/sqle/driver/mysql/session"
)

// ==== Rule test code start ====
func TestRuleSQL00089(t *testing.T) {
	ruleName := ai.SQLE00089
	rule := rulepkg.RuleHandlerMap[ruleName].Rule

	// case 1: INSERT INTO 使用 SELECT 和 FROM 子句
	runAIRuleCase(rule, t, "case 1: INSERT INTO 使用 SELECT 和 FROM 子句",
		"INSERT INTO table1 (column1) SELECT column2 FROM table2;",
		session.NewAIMockContext().WithSQL("CREATE TABLE table1 (column1 INT); CREATE TABLE table2 (column2 INT);"),
		nil, newTestResult().addResult(ruleName))

	// case 2: INSERT INTO 使用 SELECT 但没有 FROM 子句
	runAIRuleCase(rule, t, "case 2: INSERT INTO 使用 SELECT 但没有 FROM 子句",
		"INSERT INTO table1 (column1) SELECT 1;",
		session.NewAIMockContext().WithSQL("CREATE TABLE table1 (column1 INT);"),
		nil, newTestResult())

	// case 3: INSERT INTO 使用 VALUES 而非 SELECT
	runAIRuleCase(rule, t, "case 3: INSERT INTO 使用 VALUES 而非 SELECT",
		"INSERT INTO table1 (column1) VALUES (1);",
		session.NewAIMockContext().WithSQL("CREATE TABLE table1 (column1 INT);"),
		nil, newTestResult())

	// case 4: INSERT INTO 使用子查询但没有 FROM 子句
	runAIRuleCase(rule, t, "case 4: INSERT INTO 使用子查询但没有 FROM 子句",
		"INSERT INTO table1 (column1) SELECT (SELECT 1);",
		session.NewAIMockContext().WithSQL("CREATE TABLE table1 (column1 INT);"),
		nil, newTestResult())

	// case 5: INSERT INTO 使用复杂 SELECT 包含 FROM 子句
	runAIRuleCase(rule, t, "case 5: INSERT INTO 使用复杂 SELECT 包含 FROM 子句",
		"INSERT INTO table1 (column1) SELECT column2 FROM table2 WHERE column2 > 10;",
		session.NewAIMockContext().WithSQL("CREATE TABLE table1 (column1 INT); CREATE TABLE table2 (column2 INT);"),
		nil, newTestResult().addResult(ruleName))

	// case 6: INSERT INTO 使用嵌套 SELECT 和 FROM 子句
	runAIRuleCase(rule, t, "case 6: INSERT INTO 使用嵌套 SELECT 和 FROM 子句",
		"INSERT INTO table1 (column1) SELECT column2 FROM (SELECT column2 FROM table2) AS subquery;",
		session.NewAIMockContext().WithSQL("CREATE TABLE table1 (column1 INT); CREATE TABLE table2 (column2 INT);"),
		nil, newTestResult().addResult(ruleName))

	// case 7: INSERT INTO 使用 SELECT 和 JOIN 子句
	runAIRuleCase(rule, t, "case 7: INSERT INTO 使用 SELECT 和 JOIN 子句",
		"INSERT INTO table1 (column1) SELECT t1.column2 FROM table2 t1 JOIN table3 t2 ON t1.id = t2.id;",
		session.NewAIMockContext().WithSQL("CREATE TABLE table1 (column1 INT); CREATE TABLE table2 (id INT, column2 INT); CREATE TABLE table3 (id INT);"),
		nil, newTestResult().addResult(ruleName))

	// case 8: INSERT INTO 使用简单 SELECT 没有 FROM 子句
	runAIRuleCase(rule, t, "case 8: INSERT INTO 使用简单 SELECT 没有 FROM 子句",
		"INSERT INTO table1 (column1) SELECT (1 + 1);",
		session.NewAIMockContext().WithSQL("CREATE TABLE table1 (column1 INT);"),
		nil, newTestResult())
}

// ==== Rule test code end ====
