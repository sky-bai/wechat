package service

import (
	"github.com/google/wire"
	"wx-base/internal/service/mini_program_service"
	"wx-base/internal/service/official_account_service"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(official_account_service.NewOfficialAccountService, mini_program_service.NewMiniProgramService)
