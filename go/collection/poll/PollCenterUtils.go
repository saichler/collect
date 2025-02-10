package poll

import "github.com/saichler/shared/go/share/strings"

func pollKey(name, vendor, series, family, software, hardware, version string) string {
	buff := strings.New()
	buff.Add(name)
	addToKey(vendor, buff)
	addToKey(series, buff)
	addToKey(family, buff)
	addToKey(software, buff)
	addToKey(hardware, buff)
	addToKey(version, buff)
	return buff.String()
}

func addToKey(str string, buff *strings.String) {
	if str != "" {
		buff.Add("+")
		buff.Add(str)
	}
}
