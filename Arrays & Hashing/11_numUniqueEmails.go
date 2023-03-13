package main

import "strings"

func numUniqueEmails(emails []string) int {
	res, maps := 0, make(map[string]int)
	for _, email := range emails {
		address := uniqueEmail(email)
		if _, ok := maps[address]; !ok {
			res++
			maps[address] = 1
		}
	}
	return res
}

func uniqueEmail(email string) string {
	parts := strings.Split(email, "@")
	local, domain := parts[0], parts[1]
	local = strings.ReplaceAll(local, ".", "")

	plusIndex := strings.Index(local, "+")
	if plusIndex > -1 {
		local = local[:plusIndex]
	}

	return local + "@" + domain
}
