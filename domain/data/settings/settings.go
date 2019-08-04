package settings

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"github.com/pkg/errors"

	"github.com/josephbudd/crud/domain/data/filepaths"
	"github.com/josephbudd/crud/domain/types"
	"github.com/josephbudd/crudsitepack"
)

// NewApplicationSettings makes a new ApplicationSettings.
// Returns a pointer to the ApplicationSettings and the error.
func NewApplicationSettings() (settings *types.ApplicationSettings, err error) {
	var fpath string
	var contents []byte
	var found bool
	fpath = filepaths.GetShortSettingsPath()
	if contents, found = crudsitepack.Contents(fpath); !found {
		err = errors.New(fmt.Sprintf("can't find %q", fpath))
		return
	}
	settings = &types.ApplicationSettings{}
	err = yaml.Unmarshal(contents, settings)
	return
}
