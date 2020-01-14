package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"github.com/ninjaneers-team/uropa/state"
	"github.com/ninjaneers-team/uropa/utils"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// WriteConfig holds settings to use to write the state file.
type WriteConfig struct {
	Workspace  string
	SelectTags []string
	Filename   string
	FileFormat Format
	WithID     bool
}

func compareID(obj1, obj2 id) bool {
	return strings.Compare(obj1.id(), obj2.id()) < 0
}

// OpaStateToFile writes a state object to file with filename.
// It will omit timestamps and IDs while writing.
func OpaStateToFile(OpaState *state.OpaState, config WriteConfig) error {
	// TODO break-down this giant function
	var file Content

	file.Workspace = config.Workspace
	// hardcoded as only one version exists currently
	file.FormatVersion = "1.1"

	selectTags := config.SelectTags
	if len(selectTags) > 0 {
		file.Info = &Info{
			SelectorTags: selectTags,
		}
	}

	policies, err := OpaState.Policies.GetAll()
	if err != nil {
		return err
	}
	for _, s := range policies {
		s := FPolicy{Policy: s.Policy}

		zeroOutID(&s, s.Name, config.WithID)
		zeroOutTimestamps(&s)
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

func zeroOutTimestamps(obj interface{}) {
	zeroOutField(obj, "CreatedAt")
	zeroOutField(obj, "UpdatedAt")
}

var zero reflect.Value

func zeroOutField(obj interface{}, field string) {
	ptr := reflect.ValueOf(obj)
	if ptr.Kind() != reflect.Ptr {
		return
	}
	v := reflect.Indirect(ptr)
	ts := v.FieldByName(field)
	if ts == zero {
		return
	}
	ts.Set(reflect.Zero(ts.Type()))
}

func zeroOutID(obj interface{}, altName *string, withID bool) {
	// withID is set, export the ID
	if withID {
		return
	}
	// altName is not set, export the ID
	if utils.Empty(altName) {
		return
	}
	// zero the ID field
	zeroOutField(obj, "ID")
}
