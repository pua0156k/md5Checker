package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {

	str := getFileMd5("./logstash-8.2.2-linux-x86_64.tar.gz")
	answer := "0e27568bcebc4cb705c96ab6855742fe"
	assert.Equal(t, answer, str)
	str = getFileMd5("./kops-darwin-amd64")
	answer = "4694d21778837215d47babe9319dd054"
	assert.Equal(t, answer, str)

}
