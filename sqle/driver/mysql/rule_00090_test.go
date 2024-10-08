package mysql

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	rulepkg "github.com/actiontech/sqle/sqle/driver/mysql/rule"
	"github.com/actiontech/sqle/sqle/driver/mysql/rule/ai"
	"github.com/actiontech/sqle/sqle/driver/mysql/session"
)

// ==== Rule test code start ====
func TestRuleSQL00090(t *testing.T) {
	ruleName := ai.SQLE00090
	rule := rulepkg.RuleHandlerMap[ruleName].Rule

	runAIRuleCase(rule, t, "case 1: SELECT 语句中包含 UNION", "SELECT id FROM exist_db.exist_tb_1 UNION SELECT id FROM exist_db.exist_tb_2;",
		nil, /*mock context*/
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 2: SELECT 语句中包含 UNION ALL", "SELECT id FROM exist_db.exist_tb_1 UNION ALL SELECT id FROM exist_db.exist_tb_2;",
		nil, /*mock context*/
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 3: 简单的 SELECT 语句", "SELECT id FROM exist_db.exist_tb_1;",
		nil, /*mock context*/
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 4: 复杂的 SELECT 语句中包含 UNION", "SELECT id FROM exist_db.exist_tb_1 WHERE id IN (SELECT id FROM exist_db.exist_tb_2 UNION SELECT id FROM table3);",
		nil, /*mock context*/
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 5: 复杂的 SELECT 语句中包含 UNION ALL", "SELECT id FROM exist_db.exist_tb_1 WHERE id IN (SELECT id FROM exist_db.exist_tb_2 UNION ALL SELECT id FROM table3);",
		nil, /*mock context*/
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 6: UNION 语句包含多个 SELECT 子句", "SELECT id FROM exist_db.exist_tb_1 UNION SELECT id FROM exist_db.exist_tb_2 UNION SELECT id FROM table3;",
		nil, /*mock context*/
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 7: UNION ALL 语句包含多个 SELECT 子句", "SELECT id FROM exist_db.exist_tb_1 UNION ALL SELECT id FROM exist_db.exist_tb_2 UNION ALL SELECT id FROM table3;",
		nil, /*mock context*/
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 8: INSERT 语句中包含 UNION", "INSERT INTO exist_db.exist_tb_1 (id) SELECT id FROM exist_db.exist_tb_2 UNION SELECT id FROM table3;",
		nil, /*mock context*/
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 9: INSERT 语句中包含 UNION ALL", "INSERT INTO exist_db.exist_tb_1 (id) SELECT id FROM exist_db.exist_tb_2 UNION ALL SELECT id FROM table3;",
		nil, /*mock context*/
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 10: UPDATE 语句中包含 UNION", "UPDATE exist_db.exist_tb_1 SET id = (SELECT id FROM exist_db.exist_tb_2 UNION SELECT id FROM table3);",
		nil, /*mock context*/
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 11: UPDATE 语句中包含 UNION ALL", "UPDATE exist_db.exist_tb_1 SET id = (SELECT id FROM exist_db.exist_tb_2 UNION ALL SELECT id FROM table3);",
		nil, /*mock context*/
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 12: DELETE 语句中包含 UNION", "DELETE FROM exist_db.exist_tb_1 WHERE id IN (SELECT id FROM exist_db.exist_tb_2 UNION SELECT id FROM table3);",
		nil, /*mock context*/
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 13: DELETE 语句中包含 UNION ALL", "DELETE FROM exist_db.exist_tb_1 WHERE id IN (SELECT id FROM exist_db.exist_tb_2 UNION ALL SELECT id FROM table3);",
		nil, /*mock context*/
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 14: SELECT 语句中包含 UNION (示例)", "SELECT name, city FROM customers UNION SELECT name, city FROM suppliers;",
		session.NewAIMockContext().WithSQL("CREATE TABLE customers(id INT(11) NOT NULL, name VARCHAR(32) DEFAULT '', sex TINYINT NOT NULL, city VARCHAR(32) NOT NULL, age SMALLINT(4) NOT NULL, PRIMARY KEY (id)); INSERT INTO customers VALUES(1,'xiaoli',1,'shanghai',18); INSERT INTO customers SELECT ID + (SELECT count(1) FROM customers),concat('t',(ID + (SELECT count(1) FROM customers))),1,'shanghai',18 FROM customers; INSERT INTO customers SELECT ID + (SELECT count(1) FROM customers),concat('t',(ID + (SELECT count(1) FROM customers))),1,'shanghai',18 FROM customers; INSERT INTO customers SELECT ID + (SELECT count(1) FROM customers),concat('t',(ID + (SELECT count(1) FROM customers))),1,'shanghai',18 FROM customers;"),
		[]*AIMockSQLExpectation{
			{
				Query: "EXPLAIN SELECT name, city FROM customers UNION SELECT name, city FROM suppliers;",
				Rows:  sqlmock.NewRows(nil),
			},
		}, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 15: SELECT 语句中包含 UNION ALL (示例)", "SELECT name, city FROM customers UNION ALL SELECT name, city FROM suppliers;",
		session.NewAIMockContext().WithSQL("CREATE TABLE customers(id INT(11) NOT NULL, name VARCHAR(32) DEFAULT '', sex TINYINT NOT NULL, city VARCHAR(32) NOT NULL, age SMALLINT(4) NOT NULL, PRIMARY KEY (id)); INSERT INTO customers VALUES(1,'xiaoli',1,'shanghai',18); INSERT INTO customers SELECT ID + (SELECT count(1) FROM customers),concat('t',(ID + (SELECT count(1) FROM customers))),1,'shanghai',18 FROM customers; INSERT INTO customers SELECT ID + (SELECT count(1) FROM customers),concat('t',(ID + (SELECT count(1) FROM customers))),1,'shanghai',18 FROM customers; INSERT INTO customers SELECT ID + (SELECT count(1) FROM customers),concat('t',(ID + (SELECT count(1) FROM customers))),1,'shanghai',18 FROM customers;"),
		[]*AIMockSQLExpectation{
			{
				Query: "EXPLAIN SELECT name, city FROM customers UNION ALL SELECT name, city FROM suppliers;",
				Rows:  sqlmock.NewRows(nil),
			},
		}, newTestResult())
}

// ==== Rule test code end ====
