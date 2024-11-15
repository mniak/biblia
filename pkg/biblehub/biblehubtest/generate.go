package biblehubtest

import (
	_ "go.uber.org/mock/gomock"
)

//go:generate mockgen -package biblehubtest -destination downloader.go .. Downloader
