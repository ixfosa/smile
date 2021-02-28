package main

import (
	"fmt"
	"net/url"
	"strings"
)

/*
type URL struct {
	Scheme     string
	Opaque     string    // encoded opaque data
	User       *Userinfo // username and password information
	Host       string    // host or host:port
	Path       string    // path (relative paths may omit leading slash)
	RawPath    string    // encoded path hint (see EscapedPath method)
	ForceQuery bool      // append a query ('?') even if RawQuery is empty
	RawQuery   string    // encoded query values, without '?'
	Fragment   string    // fragment for references, without '#'
}
*
 */
func main() {
	// 这个 URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// func Parse(rawurl string) (*URL, error)
	parse, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(parse.Scheme) 		   	// postgres

	fmt.Println(parse.User)   		   	// user:pass
	fmt.Println(parse.User.Username())	// user
	fmt.Println(parse.User.Password())	// pass true

	fmt.Println(parse.Path) 			// /path
	fmt.Println(parse.Fragment)			// f

	fmt.Println(parse.Host) 			// host.com:5432
	h := strings.Split(parse.Host, ":")
	fmt.Println(h[0]) 					// host.com
	fmt.Println(h[1])					// 5432


	fmt.Println(parse.RawQuery) 			// k=v
	query, _ := url.ParseQuery(parse.RawQuery)
	fmt.Println(query) // map[k:[v]]

	fmt.Println(query["k"]) 			// [v]
	fmt.Println(query["k"][0]) 			// v

}
