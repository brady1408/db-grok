package command

func inArgs(args []string, table string) bool {
	if len(args) < 1 {
		return true
	}
	inArgs := false
	for _, a := range args {
		if table == a {
			inArgs = true
			break
		}
	}
	return inArgs
}
