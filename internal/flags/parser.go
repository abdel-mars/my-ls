package flags

// import "fmt"
import "strings"


type Flags struct {
    Recursive bool
}

func Parse(args []string) (Flags, []string) {
	flags := Flags{}
	paths := []string{}

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") && len(arg) > 1 {

			for _, flag := range arg[1:] {
				switch flag {
				case 'R':
					flags.Recursive = true
				}
			}
		} else {
			paths = append(paths, arg)
		}
}
return flags, paths
}

