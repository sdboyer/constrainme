package constrainme

import "github.com/Sirupsen/logrus"

var A = logrus.SetFormatter

func init() {
	A(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
