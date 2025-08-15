package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/synw/goinfer/state"
	"github.com/synw/goinfer/types"
	"gopkg.in/yaml.v3"
)

func readTemplates(m map[string]interface{}) (map[string]types.TemplateInfo, error) {
	info := map[string]types.TemplateInfo{}

	for model, conf := range m {
		c := conf.([]interface{})
		mi := types.TemplateInfo{}
		for _, vx := range c {
			val := vx.(map[string]interface{})
			for k, v := range val {
				switch k {
				case "ctx":
					mi.Ctx = v.(int)
				case "template":
					mi.Name = v.(string)
				}
			}
		}
		info[model] = mi
	}

	return info, nil
}

func ReadTemplates() (map[string]types.TemplateInfo, error) {
	m := make(map[string]interface{})
	p := filepath.Join(state.ModelsDir, "templates.yml")
	info := map[string]types.TemplateInfo{}
	// fmt.Println("Opening", p)
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		return info, fmt.Errorf("templates file not found: %s", p)
	}

	file, err := os.Open(p)
	if err != nil {
		return info, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return info, err
	}

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		return info, err
	}

	info, err = readTemplates(m)
	if err != nil {
		return info, err
	}

	return info, nil
}
