package stringPrac

import "bytes"

func replaceSpace(s string) string {
	buf := make([]byte, 0, len(s)*3)
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == ' ' {
			buf = append(buf, "%20"...)
		} else {
			buf = append(buf, bytes[i])
		}
	}
	return string(buf)
}

func comma(s string) string {
	buffer := bytes.Buffer{}
	i := len(s) % 3
	if i > 0 {
		buffer.WriteString(s[:(len(s) % 3)])
		buffer.WriteString(",")
	}
	for ; i+3 < len(s); i += 3 {
		buffer.WriteString(s[i : i+3])
		buffer.WriteString(",")
	}
	if i < len(s) {
		buffer.WriteString(s[i:])
	}
	return buffer.String()
}
