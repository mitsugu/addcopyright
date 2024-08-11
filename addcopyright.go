package main

import (
    "encoding/json"
    "fmt"
    "io"
    "os"
    "os/exec"

    "github.com/urfave/cli/v2"
)

const configFileName = "addcopyright.json"

type Config struct {
    ExifToolPath string `json:"exiftool_path"`
    Copyright    string `json:"copyright"`
}

func main() {
    app := &cli.App{
        Name:  "addcopyright",
        Usage: "Add copyright information to an image's EXIF data",
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:     "input",
                Usage:    "Path to the input image file",
                Required: true,
            },
            &cli.StringFlag{
                Name:     "output",
                Usage:    "Path to the output image file",
                Required: true,
            },
        },
        Action: func(c *cli.Context) error {
            inputFilePath := c.String("input")
            outputFilePath := c.String("output")

            config, err := readConfig()
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

func readConfig() (*Config, error) {
    file, err := os.Open(configFileName)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var config Config
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&config); err != nil {
        return nil, err
    }

    return &config, nil
}

func copyFile(src, dst string) error {
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
