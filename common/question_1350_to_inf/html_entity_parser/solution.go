package html_entity_parser

import "strings"

// https://leetcode.cn/problems/html-entity-parser/
func entityParser(text string) string {
	new0 := strings.Replace(text, "&quot;", "\"", -1)
	new1 := strings.Replace(new0, "&apos;", "'", -1)
	new2 := strings.Replace(new1, "&amp;", "&", -1)
	new3 := strings.Replace(new2, "&gt;", ">", -1)
	new4 := strings.Replace(new3, "&lt;", "<", -1)
	new5 := strings.Replace(new4, "&frasl;", "/", -1)
	return new5
}

// 双引号：字符实体为 &quot; ，对应的字符是 " 。
// 单引号：字符实体为 &apos; ，对应的字符是 ' 。
// 与符号：字符实体为 &amp; ，对应对的字符是 & 。
// 大于号：字符实体为 &gt; ，对应的字符是 > 。
// 小于号：字符实体为 &lt; ，对应的字符是 < 。
// 斜线号：字符实体为 &frasl; ，对应的字符是 / 。
