package piscine

func ConcatParams(args []string) string {
	smyrep := ""
	for _, str := range args[:len(args)-1] {
		smyrep += str + "\n"
	}
	return smyrep + args[len(args)-1]
}
