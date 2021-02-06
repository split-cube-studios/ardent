// Package assetutil contains utilities for generating, parsing, and managing assets.
package assetutil

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func parseConfigFiles(dir string) (confs []config, err error) {
	err = filepath.Walk(
		dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if filepath.Ext(path) != ".yml" {
				return nil
			}

			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()

			var conf config
			if err := yaml.NewDecoder(f).Decode(&conf); err != nil {
				return err
			}

			// NOTE: sounds handle filepaths within the configuration
			conf.filepath = strings.Replace(path, ".yml", ".png", 1)
			confs = append(confs, conf)

			return nil
		},
	)

	return //nolint:nakedret
}

// CreateAssets takes a directory path that should contain a yml config file and uses it to produce a .asset file.
func CreateAssets(dir string) {
	confs, err := parseConfigFiles(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, conf := range confs {
		asset, err := conf.toAsset()
		if err != nil {
			log.Fatal(err)
		}

		d, err := asset.Marshal()
		if err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile(
			strings.Replace(conf.filepath, ".png", ".asset", 1),
			d,
			0o600,
		); err != nil {
			log.Fatal(err)
		}
	}
}
