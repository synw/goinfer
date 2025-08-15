package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/synw/goinfer/state"
	"github.com/synw/goinfer/types"
	"gopkg.in/yaml.v3"
)

func SaveTask(task types.Task) error {
	if state.IsVerbose {
		fmt.Println("Saving task", task)
	}
	// Ensure task name has .yml extension and extract base name
	taskName := task.Name
	// Extract base name (remove directory path)
	baseName := filepath.Base(taskName)
	// Ensure base name has .yml extension
	if !strings.HasSuffix(baseName, ".yml") {
		baseName += ".yml"
	}

	filePath := filepath.Join(state.TasksDir, taskName)
	// But we need to ensure the file path has .yml extension
	if !strings.HasSuffix(filePath, ".yml") {
		// Extract directory and base name, then add .yml to base name
		dir := filepath.Dir(filePath)
		base := filepath.Base(filePath)
		if !strings.HasSuffix(base, ".yml") {
			base += ".yml"
		}
		filePath = filepath.Join(dir, base)
	}

	// Create a copy of the task with normalized name for YAML serialization
	taskForYAML := task
	taskForYAML.Name = baseName

	// Create directory if it doesn't exist
	err := os.MkdirAll(filepath.Dir(filePath), 0o755)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	enc := yaml.NewEncoder(file)
	err = enc.Encode(&taskForYAML)
	if err != nil {
		return err
	}

	// Don't forget to close the file
	err = file.Close()
	if err != nil {
		return err
	}

	if state.IsVerbose {
		fmt.Println("Task", task.Name, "file successfully")
	}

	return nil
}
