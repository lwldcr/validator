/*
Package validator provides a negroni(github.com/urfave/negroni) compatible
middleware to do data validating before handles requests.

requests should be like:
${host}:${port}/${some route}/?sign=xxxx

sign gives current request data signature

test/test.go implements a simple demo showing how to use.
 */
package validator
