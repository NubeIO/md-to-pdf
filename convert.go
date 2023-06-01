package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func (app *application) convert(ctx context.Context, inputFile []byte, writeToHomeDir bool) ([]byte, error) {
	tmpdir := path.Join(os.TempDir(), fmt.Sprintf("pandocserver_%s", randStringRunes(10)))

	if err := os.Mkdir(tmpdir, permDir); err != nil {
		return nil, fmt.Errorf("could not create dir %q: %w", tmpdir, err)
	}
	defer os.RemoveAll(tmpdir)

	inputFileName := filepath.Join(tmpdir, fmt.Sprintf("%s.md", randStringRunes(10)))
	if err := os.WriteFile(inputFileName, inputFile, permWrite); err != nil {
		return nil, fmt.Errorf("could not create inputfile: %w", err)
	}

	outputDir := path.Join(tmpdir, "output")
	if err := os.Mkdir(outputDir, permDir); err != nil {
		return nil, fmt.Errorf("could not create output directory: %w", err)
	}
	pdfFilename := fmt.Sprintf("%s.pdf", randStringRunes(10))
	outputFilename := filepath.Join(outputDir, pdfFilename)

	args := []string{
		inputFileName,
		fmt.Sprintf("--output=%s", outputFilename),
		fmt.Sprintf("--data-dir=%s", app.pandocDataDir),
		"--from=markdown+yaml_metadata_block+raw_html+emoji",
	}
	commandCtx, cancel := context.WithTimeout(ctx, app.commandTimeout)
	defer cancel()

	var out bytes.Buffer
	var stderr bytes.Buffer
	logger.Infof("path: %s", app.pandocPath)
	logger.Infof("going to call pandoc with the following args: %v", args)
	cmd := exec.CommandContext(commandCtx, "pandoc", args...)
	cmd.Dir = tmpdir
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		app.killProcessIfRunning(cmd)
		return nil, fmt.Errorf("could not execute command %w: %s", err, stderr.String())
	}

	app.killProcessIfRunning(cmd)

	content, err := os.ReadFile(outputFilename)
	if err != nil {
		return nil, fmt.Errorf("could not read output file: %w", err)
	}
	pdfFile := fmt.Sprintf("%s/%s", app.userHome, pdfFilename)
	if writeToHomeDir {
		logger.Infof("write pdf to dir: %s", pdfFile)
		ioutil.WriteFile(pdfFile, content, 0644)
	}

	return content, nil
}

func (app *application) convertFromRead(ctx context.Context, inputFile string) (bool, error) {
	pdfFilename := fmt.Sprintf("%s.pdf", randStringRunes(10))
	outputFilename := filepath.Join(app.userHome, pdfFilename)
	args := []string{ // works with MD and images
		inputFile,
		fmt.Sprintf("--output=%s", outputFilename),
		fmt.Sprintf("--data-dir=%s", app.pandocDataDir),
		"--from=markdown-implicit_figures",
	}

	commandCtx, cancel := context.WithTimeout(ctx, app.commandTimeout)
	defer cancel()

	var out bytes.Buffer
	var stderr bytes.Buffer
	logger.Infof("path: %s", app.pandocPath)
	logger.Infof("going to call pandoc with the following args: %v", args)
	cmd := exec.CommandContext(commandCtx, "pandoc", args...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		app.killProcessIfRunning(cmd)
		return false, fmt.Errorf("could not execute command %w: %s", err, stderr.String())
	}
	app.killProcessIfRunning(cmd)
	return true, nil
}
