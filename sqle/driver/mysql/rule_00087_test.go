package mysql

import (
	"testing"

	rulepkg "github.com/actiontech/sqle/sqle/driver/mysql/rule"
	"github.com/actiontech/sqle/sqle/driver/mysql/rule/ai"
	"github.com/actiontech/sqle/sqle/driver/mysql/session"
)

// ==== Rule test code start ====
func TestRuleSQL00087(t *testing.T) {
	ruleName := ai.SQLE00087
	rule := rulepkg.RuleHandlerMap[ruleName].Rule

	runAIRuleCase(rule, t, "case 1: DELETE语句，WHERE子句中IN参数个数超过阈值",
		"DELETE FROM employees WHERE id IN (1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51);",
		session.NewAIMockContext().WithSQL("CREATE TABLE employees (id INT, name VARCHAR(100));"),
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 2: DELETE语句，WHERE子句中IN参数个数未超过阈值",
		"DELETE FROM employees WHERE id IN (1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50);",
		session.NewAIMockContext().WithSQL("CREATE TABLE employees (id INT, name VARCHAR(100));"),
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 3: DELETE语句，无WHERE子句",
		"DELETE FROM employees;",
		session.NewAIMockContext().WithSQL("CREATE TABLE employees (id INT, name VARCHAR(100));"),
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 4: INSERT语句，SELECT子句中WHERE子句的IN参数个数超过阈值",
		"INSERT INTO archive SELECT * FROM employees WHERE id IN (1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51);",
		session.NewAIMockContext().WithSQL("CREATE TABLE employees (id INT, name VARCHAR(100)); CREATE TABLE archive (id INT, name VARCHAR(100));"),
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 5: INSERT语句，SELECT子句中WHERE子句的IN参数个数未超过阈值",
		"INSERT INTO archive SELECT * FROM employees WHERE id IN (1, 2, 3, 50);",
		session.NewAIMockContext().WithSQL("CREATE TABLE employees (id INT, name VARCHAR(100)); CREATE TABLE archive (id INT, name VARCHAR(100));"),
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 6: UPDATE语句，WHERE子句中IN参数个数超过阈值",
		"UPDATE employees SET status = 'inactive' WHERE id IN (1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51);",
		session.NewAIMockContext().WithSQL("CREATE TABLE employees (id INT, name VARCHAR(100), status VARCHAR(20));"),
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 7: UPDATE语句，WHERE子句中IN参数个数未超过阈值",
		"UPDATE employees SET status = 'inactive' WHERE id IN (1, 2, 3, 50);",
		session.NewAIMockContext().WithSQL("CREATE TABLE employees (id INT, name VARCHAR(100), status VARCHAR(20));"),
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 8: SELECT语句，WHERE子句中IN参数个数超过阈值",
		"SELECT * FROM employees WHERE id IN (1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51);",
		session.NewAIMockContext().WithSQL("CREATE TABLE employees (id INT, name VARCHAR(100));"),
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 9: SELECT语句，WHERE子句中IN参数个数未超过阈值",
		"SELECT * FROM employees WHERE id IN (1, 2, 3, 50);",
		session.NewAIMockContext().WithSQL("CREATE TABLE employees (id INT, name VARCHAR(100));"),
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 10: UNION语句，第一个SELECT子句WHERE子句中IN参数个数超过阈值",
		"SELECT * FROM employees WHERE id IN (1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51) UNION SELECT * FROM managers;",
		session.NewAIMockContext().WithSQL("CREATE TABLE employees (id INT, name VARCHAR(100)); CREATE TABLE managers (id INT, name VARCHAR(100));"),
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 11: UNION语句，第二个SELECT子句WHERE子句中IN参数个数超过阈值",
		"SELECT * FROM employees UNION SELECT * FROM managers WHERE id IN (1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51);",
		session.NewAIMockContext().WithSQL("CREATE TABLE employees (id INT, name VARCHAR(100)); CREATE TABLE managers (id INT, name VARCHAR(100));"),
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 12: UNION语句，所有SELECT子句WHERE子句中IN参数个数未超过阈值",
		"SELECT * FROM employees WHERE id IN (1, 2, 3, 50) UNION SELECT * FROM managers WHERE id IN (1, 2, 3, 50);",
		session.NewAIMockContext().WithSQL("CREATE TABLE employees (id INT, name VARCHAR(100)); CREATE TABLE managers (id INT, name VARCHAR(100));"),
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 13: SELECT语句，WHERE子句中IN参数个数超过阈值，使用t2表",
		"SELECT * FROM t2 WHERE name IN ('a', 'b', 'c',1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,'zz');",
		session.NewAIMockContext().WithSQL("CREATE TABLE t2 (id INT, name VARCHAR(100), type VARCHAR(50), addr VARCHAR(100));"),
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 14: SELECT语句，WHERE子句中IN参数个数未超过阈值，使用t2表",
		"SELECT * FROM t2 WHERE name IN ('t1', '2t');",
		session.NewAIMockContext().WithSQL("CREATE TABLE t2 (id INT, name VARCHAR(100), type VARCHAR(50), addr VARCHAR(100));"),
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 15: INSERT语句，SELECT子句中WHERE子句的IN参数个数未超过阈值，使用t2表",
		"INSERT INTO t2 SELECT id,name,type,addr FROM t2 WHERE name IN ('t1','a3');",
		session.NewAIMockContext().WithSQL("CREATE TABLE t2 (id INT, name VARCHAR(100), type VARCHAR(50), addr VARCHAR(100));"),
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 16: UPDATE语句，WHERE子句中IN参数个数未超过阈值，使用t2表",
		"UPDATE t2 SET type ='3' WHERE name IN ('t1','2t');",
		session.NewAIMockContext().WithSQL("CREATE TABLE t2 (id INT, name VARCHAR(100), type VARCHAR(50), addr VARCHAR(100));"),
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 17: UPDATE语句，WHERE子句中IN参数个数超过阈值，使用t2表",
		"UPDATE t2 set type ='3' WHERE name IN ('a', 'b', 'c',1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51, 'zz');",
		session.NewAIMockContext().WithSQL("CREATE TABLE t2 (id INT, name VARCHAR(100), type VARCHAR(50), addr VARCHAR(100));"),
		nil, newTestResult().addResult(ruleName))

	runAIRuleCase(rule, t, "case 18: DELETE语句，WHERE子句中IN参数个数未超过阈值，使用t2表",
		"DELETE FROM t2 WHERE name IN ('t1','2t');",
		session.NewAIMockContext().WithSQL("CREATE TABLE t2 (id INT, name VARCHAR(100), type VARCHAR(50), addr VARCHAR(100));"),
		nil, newTestResult())

	runAIRuleCase(rule, t, "case 19: DELETE语句，WHERE子句中IN参数个数超过阈值，使用t2表",
		"DELETE FROM t2 WHERE name IN ('a', 'b', 'c',1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51, 'zz');",
		session.NewAIMockContext().WithSQL("CREATE TABLE t2 (id INT, name VARCHAR(100), type VARCHAR(50), addr VARCHAR(100));"),
		nil, newTestResult().addResult(ruleName))
}

// ==== Rule test code end ====
