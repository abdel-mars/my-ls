package flags

// import "fmt"
import "strings"


type Flags struct {
    Recursive bool
	All bool
}

func Parse(args []string) (Flags, []string) {
	flags := Flags{}
	paths := []string{}

	// if !flags.All {
	// 	fmt.Println("All is not true")
	// }

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") && len(arg) > 1 {

			for _, flag := range arg[1:] {
				switch flag {
				case 'R':
					flags.Recursive = true
				case 'a':
					flags.All = true
				}
			}
		} else {
			paths = append(paths, arg)
		}
}
return flags, paths
}

