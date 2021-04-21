package verifier

import "strconv"

type CheckResult struct {
	Address     string
	Deliverable bool
	FullInbox   bool
	Disabled    bool
	Error       error
}

func (r CheckResult) StringSlice() []string {
	out := []string{
		r.Address,
		strconv.FormatBool(r.Deliverable),
		strconv.FormatBool(r.FullInbox),
		strconv.FormatBool(r.Disabled),
	}
	if r.Error != nil {
		out = append(out, r.Error.Error())
	} else {
		out = append(out, "")
	}
	return out
}
