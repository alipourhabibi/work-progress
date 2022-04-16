package files

import "os"

var (
	PARRENT string
	WORK    string
)

func Init() {
	a, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	PARRENT = a + "/work-progress/"
	WORK = a + "/work-progress/works.json"
}
