package mysql

import (
	"testing"

	rulepkg "github.com/actiontech/sqle/sqle/driver/mysql/rule"
	"github.com/actiontech/sqle/sqle/driver/mysql/rule/ai"
	"github.com/actiontech/sqle/sqle/driver/mysql/session"
)

// ==== Rule test code start ====
func TestRuleSQLE00089(t *testing.T) {
	ruleName := ai.SQLE00089
	rule := rulepkg.RuleHandlerMap[ruleName].Rule

	runAIRuleCase(rule, t, "case 1: INSERT INTO 使用 SELECT 语句",
		"INSERT INTO target_table (column1) SELECT column1 FROM source_table;",
		session.NewAIMockContext().WithSQL(
			"CREATE TABLE target_table (column1 INT);CREATE TABLE source_table (column1 INT);",
		),
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 2: INSERT INTO 使用 WITH 语句",
		"INSERT INTO target_table (column1) WITH cte AS (SELECT column1 FROM source_table) SELECT column1 FROM cte;",
		session.NewAIMockContext().WithSQL(
			"CREATE TABLE target_table (column1 INT);CREATE TABLE source_table (column1 INT);",
		),
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 3: INSERT INTO 使用 VALUES 语句",
		"INSERT INTO target_table (column1) VALUES ('value1');",
		session.NewAIMockContext().WithSQL(
			"CREATE TABLE target_table (column1 VARCHAR(255));",
		),
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 4: UNION 语句中包含 INSERT INTO ... SELECT",
		"SELECT column1 FROM table1 UNION SELECT column1 FROM table2;",
		session.NewAIMockContext().WithSQL(
			"CREATE TABLE table1 (column1 INT);CREATE TABLE table2 (column1 INT);",
		),
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 5: UNION 语句中不包含 INSERT INTO ... SELECT",
		"SELECT column1 FROM table1 UNION SELECT column1 FROM table2;",
		session.NewAIMockContext().WithSQL(
			"CREATE TABLE table1 (column1 INT);CREATE TABLE table2 (column1 INT);",
		),
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 6: INSERT INTO 使用 JOIN 语句(从xml中补充)",
		"INSERT INTO target_table (column1) SELECT a.column1 FROM table1 a JOIN table2 b ON a.id = b.id;",
		session.NewAIMockContext().WithSQL(
			"CREATE TABLE target_table (column1 INT);CREATE TABLE table1 (id INT, column1 INT);CREATE TABLE table2 (column1 INT);",
		),
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 7: INSERT INTO 使用子查询(从xml中补充)",
		"INSERT INTO target_table (column1) SELECT (SELECT column1 FROM table2 WHERE id = 1) AS subquery_result;",
		session.NewAIMockContext().WithSQL(
			"CREATE TABLE target_table (column1 INT);CREATE TABLE table2 (id INT, column1 INT);",
		),
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 8: INSERT INTO 使用常量(从xml中补充)",
		"INSERT INTO target_table (column1) VALUES (42);",
		session.NewAIMockContext().WithSQL(
			"CREATE TABLE target_table (column1 INT);",
		),
		nil, newTestResult())
}

// ==== Rule test code end ====
