package hosts

import (
	"bufio"
	"io"
	"os"
	"strings"
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
	f, err := open("/etc/hosts")
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
	f, err := open("/etc/hosts")
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
	line := 1
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
			// set the IP address
			record.IP = row[0]
			// append all hosts assigned to the IP
			record.Hosts = row[1:]

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
