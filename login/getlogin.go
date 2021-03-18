package login

import (
	"os"

	"github.com/marguerite/go-gnulib/ttyname"
	"github.com/marguerite/go-gnulib/utmp"
)

func GetLogin() (string, error) {

	// GetLogin is based on stdin, so find the name of the current terminal
	// and return nothing if it doesn't exist. According to GNU, this is what
	// DEC Unix, SunOS, Solaris, and HP-UX all do.
	name, err := ttyname.TtyName(os.Stdin.Fd())
	if err != nil {
		return "", err
	}

	u := new(utmp.Utmp)
	_ = copy(u.Line[:], []byte(name[5:]))

	file, err := utmp.Open(utmp.UtmpxFile, utmp.Reading)
	if err != nil {
		return "", err
	}

	line, err := u.GetUtLine(file)
	if err != nil {
		return "", err
	}

	return string(line.User[:]), nil
}
