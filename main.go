package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {
		data := reader.Text()
		ipv6_regex := `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
		ipv4_regex := `\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`
		ipv4 := regexp.MustCompile(ipv4_regex)
		ipv6 := regexp.MustCompile(ipv6_regex)
		if ipv4.MatchString(data) {fmt.Println("IP found")}
		censoredData := ipv4.ReplaceAllLiteralString(data, "[IP REDACTED]")
		censoredData = ipv6.ReplaceAllLiteralString(censoredData, "[IP REDACTED]")
		fmt.Println(censoredData)
	}
}
