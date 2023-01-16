package translator

import (
	"github.com/imdario/mergo"
	"github.com/scarlet0725/prism-api/model"
)

func MergeUser(base, patch *model.User) (*model.User, error) {
	baseUser := *base
	patchData := *patch
	err := mergo.Merge(&patchData, baseUser)
	if err != nil {
		return nil, err
	}

	//mergoの仕様上ゼロ値を上書きしないのでture -> falseの場合はfalseにする
	if !baseUser.IsAdminVerified && patchData.IsAdminVerified {
		patchData.IsAdminVerified = true
	}

	if !baseUser.DeleteProtected && patchData.DeleteProtected {
		patchData.DeleteProtected = true
	}

	return &patchData, nil
}
