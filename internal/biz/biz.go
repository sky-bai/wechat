package biz

import (
	"github.com/google/wire"
	"wx-base/internal/biz/mini_program_biz"
	"wx-base/internal/biz/official_account_biz"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(official_account_biz.NewOfficialAccountUseCase, mini_program_biz.NewMiniProgramUseCase)
