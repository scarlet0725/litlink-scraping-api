package translator

import (
	"github.com/imdario/mergo"
	"github.com/scarlet0725/prism-api/model"
)

func MergeUser(base, patch *model.User) error {
	err := mergo.Merge(base, *patch)
	if err != nil {
		return err
	}
	return nil
}
