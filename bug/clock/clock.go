package clock

import (
	"fmt"
	"github.com/MichaelMure/git-bug/repository"
	"io/ioutil"
	"path"
)

const ClockFile = ".git/git-bug-clock"

var BugClock LamportClock

func readClock(repo repository.Repo) error {
	filePath := path.Join(repo.GetPath(), ClockFile)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var value uint64
	n, err := fmt.Sscanf(string(content), "%d", &value)

	if err != nil {
		return err
	}

	if n != 1 {
		return fmt.Errorf("could not read the clock")
	}

	BugClock = LamportClock{counter: value}

	fmt.Println(value)

	return nil
}

func WriteClock(repo repository.Repo) error {
	filePath := path.Join(repo.GetPath(), ClockFile)

	data := []byte(fmt.Sprintf("%d", BugClock.Time()))
	return ioutil.WriteFile(filePath, data, 0644)
}

//func Witness(bug bug.Bug) {
//
//}
