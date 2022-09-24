package main

import (
	"fmt"
	"strings"
)

func securityCheck(in string) error {
	in = strings.Replace(in, "/r/n", "/n", -1)

	c := strings.Split(in, anchorStr)
	if len(c) != 2 {
		return fmt.Errorf("don't modify anchor")
	}

	upper, lower := c[0], c[1]

	// upper half must starts with 'from diagrams'
	uls := strings.Split(upper, "/n")
	for _, ul := range uls {
		if strings.HasPrefix(ul, "#") {
			continue
		}
		if !strings.HasPrefix(ul, "from diagrams") {
			return fmt.Errorf("only allow import from diagrams")
		}
	}

	// don't allow import in lower half
	if strings.Contains("import", lower) {
		return fmt.Errorf("import should be placed in above of the anchor")
	}

	return nil
}
