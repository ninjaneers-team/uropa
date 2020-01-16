package file

import (
	"encoding/json"
	"fmt"
	"github.com/ninjaneers-team/uropa/state"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

// WriteConfig holds settings to use to write the state file.
type WriteConfig struct {
	Filename   string
	FileFormat Format
}

func compareID(obj1, obj2 id) bool {
	return strings.Compare(obj1.id(), obj2.id()) < 0
}

// OpaStateToFile writes a state object to file with filename.
// It will omit timestamps and IDs while writing.
func OpaStateToFile(OpaState *state.OpaState, config WriteConfig) error {
	var file Content

	// hardcoded as only one version exists currently
	file.FormatVersion = "1.1"

	policies, err := OpaState.Policies.GetAll()
	if err != nil {
		return err
	}
	for _, s := range policies {
		s := FPolicy{Policy: s.Policy}

		file.Policies = append(file.Policies, s)
	}
	sort.SliceStable(file.Policies, func(i, j int) bool {
		return compareID(file.Policies[i], file.Policies[j])
	})

	return writeFile(file, config.Filename, config.FileFormat)
}

func writeFile(content Content, filename string, format Format) error {
	var c []byte
	var err error
	switch format {
	case YAML:
		c, err = yaml.Marshal(content)
		if err != nil {
			return err
		}
	case JSON:
		c, err = json.MarshalIndent(content, "", "  ")
		if err != nil {
			return err
		}
	default:
		return errors.New("unknown file format: " + string(format))
	}

	if filename == "-" {
		_, err = fmt.Print(string(c))
	} else {
		filename = addExtToFilename(filename, string(format))
		err = ioutil.WriteFile(filename, c, 0600)
	}
	if err != nil {
		return errors.Wrap(err, "writing file")
	}
	return nil
}

func addExtToFilename(filename, format string) string {
	if filepath.Ext(filename) == "" {
		filename = filename + "." + strings.ToLower(format)
	}
	return filename
}
