package config

import (
	"io/ioutil"

	"github.com/gsamokovarov/jump/scoring"
)

// Temporary setups a temporary jump configuration folder that can be
// used for integration testing. Needs a dir and prefix passed directly
// to ioutil.TempDir.
func Temporary(dir, prefix string) (Config, error) {
	tempDir, err := ioutil.TempDir(dir, prefix)
	if err != nil {
		return nil, err
	}

	conf, err := Setup(tempDir)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

// Testing is an in-memory testing config which loosely follows the default
// file-based configuration behavior.
type Testing struct {
	Entries  scoring.Entries
	Search   Search
	Pins     map[string]string
	Pin      string
	Settings Settings
}

// ReadEntries implements the Config interface.
func (c *Testing) ReadEntries() (scoring.Entries, error) {
	return c.Entries, nil
}

// WriteEntries implements the Config interface.
func (c *Testing) WriteEntries(entries scoring.Entries) error {
	c.Entries = entries
	c.Entries.Sort()
	return nil
}

// ReadSettings implements the Config interface.
func (c *Testing) ReadSettings() Settings {
	return c.Settings
}

// WriteSettings implements the Config interface.
func (c *Testing) WriteSettings(settings Settings) error {
	c.Settings = settings
	return nil
}

// ReadSearch implements the Config interface.
func (c *Testing) ReadSearch() Search {
	return c.Search
}

// WriteSearch implements the Config interface.
func (c *Testing) WriteSearch(term string, index int) error {
	c.Search.Term = term
	c.Search.Index = index
	return nil
}

// ReadPins implements the Config interface.
func (c *Testing) ReadPins() (map[string]string, error) {
	return c.Pins, nil
}

// FindPin implements the Config interface.
func (c *Testing) FindPin(term string) (string, bool) {
	return c.Pin, c.Pin != ""
}

// WritePin implements the Config interface.
func (c *Testing) WritePin(_, value string) error {
	c.Pin = value
	return nil
}

// RemovePin implements the Config interface.
func (c *Testing) RemovePin(term string) error {
	c.Pin = ""
	return nil
}
