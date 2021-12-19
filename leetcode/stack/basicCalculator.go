package stack

import "strconv"

// 表达式运算入口
func calculate(s string) int {
	tokens := getTokens(s)
	rpn := convertRPN(tokens)
	res := evalRPN(rpn)
	return res
}

// 逆波兰表达式计算
func evalRPN(tokens []string) int {
	ops := map[string]bool{"+": true, "-": true, "*": true, "/": true}
	cnt := 0
	opn := make([]int, len(tokens))
	for i := 0; i < len(tokens); i++ {
		if _, ok := ops[tokens[i]]; ok {
			top := simpleCalc(opn[cnt-2], opn[cnt-1], tokens[i])
			opn[cnt-2] = top
			cnt = cnt - 1
		} else {
			opn[cnt], _ = strconv.Atoi(tokens[i])
			cnt++
		}
	}
	return opn[0]
}

// 最基础的计算
func simpleCalc(x, y int, token string) int {
	switch token {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	default:
		return 0
	}
}

// 中缀表达式转后缀表达式
func convertRPN(tokens []string) []string {
	n := len(tokens)
	ops := make([]string, n)
	opn := make([]string, n)
	sIdx, nIdx := 0, 0
	opsMap := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "(": 0, ")": 0}
	for i := 0; i < len(tokens); i++ {
		// 数字入队
		if _, ok := opsMap[tokens[i]]; !ok {
			opn[nIdx] = tokens[i]
			nIdx++
		} else {
			// 运算符
			if tokens[i] == "(" {
				// （直接入操作符队
				ops[sIdx] = tokens[i]
				sIdx++
			} else if tokens[i] == ")" {
				// ）,把括号部分倒叙写入到数字队列
				for j := sIdx - 1; j >= 0 && ops[j] != "("; j-- {
					opn[nIdx] = ops[j]
					nIdx++
					sIdx--
				}
				sIdx--
			} else if sIdx == 0 || opsMap[tokens[i]] > opsMap[ops[sIdx-1]] {
				// 优先级高的直接入队
				ops[sIdx] = tokens[i]
				sIdx++
			} else {
				// 优先级低，先把操作符队列中优先级低的倒叙写入到数字队列
				for j := sIdx - 1; j >= 0 && opsMap[tokens[i]] <= opsMap[ops[j]]; j-- {
					opn[nIdx] = ops[j]
					nIdx++
					sIdx--
				}
				ops[sIdx] = tokens[i]
				sIdx++
			}
		}
	}
	for j := sIdx - 1; j >= 0; j-- {
		opn[nIdx] = ops[j]
		nIdx++
		sIdx--
	}
	return opn[0:nIdx]
}

func getTokens(s string) []string {
	opsMap := map[rune]int{'+': 1, '-': 1, '*': 2, '/': 2, '(': 0, ')': 0}
	n := len(s)
	token := make([]rune, 0)
	tokens := make([]string, 0)
	cnt := 0
	rs := []rune(s)
	for i := 0; i < n; i++ {
		if rs[i] == ' ' {
			continue
		} else if rs[i] == '-' && (cnt == 0 || tokens[cnt-1] == "(") && len(token) == 0 {
			// 支持负数的核心所在，检测到负数-x就在前面加0，即0-x
			tokens = append(tokens, strconv.FormatInt(0, 10), string(rs[i]))
			cnt += 2
		} else if isDigit(rs[i]) {
			token = append(token, rs[i])
		} else if _, ok := opsMap[rs[i]]; ok {
			if string(token) != "" {
				tokens = append(tokens, string(token))
				cnt++
			}
			token = []rune{rs[i]}
			tokens = append(tokens, string(token))
			token = make([]rune, 0)
			cnt++
		}
	}

	if string(token) != "" {
		tokens = append(tokens, string(token))
	}
	return tokens
}

// 判断是否数字
func isDigit(r rune) bool {
	code := r - '0'
	if code >= 0 && code <= 9 {
		return true
	}
	return false
}
