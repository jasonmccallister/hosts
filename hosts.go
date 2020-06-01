package hosts

import (
	"bufio"
	"io"
	"net"
	"os"
	"runtime"
	"strings"
)

const (
	WindowsHostsFileLocation = ""
	NixHostsFilesLocation    = "/etc/hosts"
)

// Record represents a line in the hosts file. The record
// can represent a comment or a normal "record" of the
// ip address and hosts assigned to that ip.
type Record struct {
	Line        int
	IsCommented bool
	IP          string
	Hosts       []string
}

func FindIP(ip string) ([]Record, error) {
	var rcds []Record

	file := NixHostsFilesLocation
	if runtime.GOOS == "windows" {
		file = WindowsHostsFileLocation
	}

	f, err := open(file)
	if err != nil {
		return rcds, err
	}

	records, err := Read(f)
	if err != nil {
		return records, err
	}

	for _, record := range records {
		if record.IP == ip {
			rcds = append(rcds, record)
		}
	}

	return rcds, nil
}

func FindHost(host string) ([]Record, error) {
	var rcds []Record

	file := NixHostsFilesLocation
	if runtime.GOOS == "windows" {
		file = WindowsHostsFileLocation
	}

	f, err := open(file)
	if err != nil {
		return rcds, err
	}

	records, err := Read(f)
	if err != nil {
		return records, err
	}

	for _, record := range records {
		for _, h := range record.Hosts {
			if h == host {
				rcds = append(rcds, record)
			}
		}
	}

	return rcds, nil
}

func Read(rdr io.Reader) ([]Record, error) {
	var records []Record

	sc := bufio.NewScanner(rdr)
	var line int
	for sc.Scan() {
		txt := sc.Text()
		line = line + 1
		record := Record{Line: line}

		// is this a commented file?
		if strings.HasPrefix(txt, "#") {
			record.IsCommented = true
		}

		// if this is an empty row
		if row := strings.Fields(txt); len(row) == 0 {
			continue
		} else {
			r := row[0]
			h := 1

			if record.IsCommented {
				r = row[1]
				h = 2
			}

			// set the IP address only if a valid ip
			ip := net.ParseIP(r)
			if ip != nil {
				record.IP = r

				// append all hosts assigned to the IP only if the IP is valid
				record.Hosts = row[h:]
			}

			records = append(records, record)
		}
	}

	return records, nil
}

func open(f string) (*os.File, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	return file, nil
}
