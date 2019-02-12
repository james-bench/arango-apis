//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

import "sort"

// GetAllEmails returns all email addresses of the given user.
func (u *User) GetAllEmails() []string {
	if u == nil {
		return nil
	}
	result := u.GetAdditionalEmails()
	if email := u.GetEmail(); email != "" {
		result = append(result, email)
	}
	// Remove duplicates
	sort.Strings(result)
	for i, v := range result {
		if (i > 0) && v == result[i-1] {
			// Remove element at i
			result = append(result[:i], result[i+1:]...)
		}
	}
	return result
}
