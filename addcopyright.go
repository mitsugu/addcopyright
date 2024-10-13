package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

type Config struct {
	ImageMagickPath string `json:"imagemagick_path"`
	ExifToolPath    string `json:"exiftool_path"`
	Copyright       string `json:"copyright"`
}

func (c *Config) readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(c); err != nil {
		return nil, err
	}

	return c, nil
}

var config Config

func main() {
	app := &cli.App{
		Name:    "addcopyright",
		Usage:   "Add copyright information to an image's EXIF data",
		Version: "v1.1.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
				Value:   "addcopyright.json",
			},
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Usage:    "Path to the input image file",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "output",
				Aliases:  []string{"o"},
				Usage:    "Path to the output image file",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			inputFilePath := c.String("input")
			outputFilePath := c.String("output")

			configPath := c.String("config")
			if configPath == "" {
				configPath = "addcopyright.json"
			}
			config, err := config.readConfig(configPath)
			if err != nil {
				return fmt.Errorf("error reading config file: %w", err)
			}

			if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
				if err := copyFile(inputFilePath, outputFilePath); err != nil {
					return fmt.Errorf("error copying file: %w", err)
				}
			}

			cmd := exec.Command(config.ExifToolPath, "-Copyright="+config.Copyright, "-Artist="+config.Copyright, outputFilePath)
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("error executing exiftool: %w", err)
			}

			fmt.Println("EXIF data updated successfully!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func copyFile(src, dst string) error {
	if src == dst {
		return errors.New("source and destination files are the same")
	}

	srcExt := strings.ToLower(filepath.Ext(src))
	dstExt := strings.ToLower(filepath.Ext(dst))

	if srcExt != dstExt {
		cmd := exec.Command(config.ImageMagickPath, src, dst)
		err := cmd.Run()
		if err != nil {
			return err
		}
		return nil
	}

	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}
