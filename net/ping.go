package net

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
)

// Ping a host using the linux ping utility
func Ping(host string, ipv6 bool) (string, error) {
	matched, err := regexp.Match(`^[\w._:-]+$`, []byte(host))
	if err != nil {
		return "", err
	}
	if !matched {
		return "", errors.New("invalid host/IP")
	}
	six := ""
	if ipv6 {
		six = "6"
	}
	// -c: packet count, -w: timeout in seconds
	out, err := exec.Command("ping"+six, "-c", "1", "-w", "3", "--", host).Output()
	if err != nil {
		errs := fmt.Sprintf("%s", err)
		if errs == "exit status 1" {
			return "", errors.New("timeout")
		}
		if errs == "exit status 2" {
			return "", errors.New("unknown host")
		}
		return "", err
	}
	r, err := regexp.Compile(`\d+ bytes from .*`)
	if err != nil {
		return "", err
	}
	line := r.Find(out)
	if line == nil {
		return "", errors.New("cannot parse ping output")
	}
	return string(line), nil
}
