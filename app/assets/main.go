package assets

import (
	"github.com/elazarl/go-bindata-assetfs"
)

func NewAssetFileSys() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:    Asset,
		AssetDir: AssetDir,
		Prefix:   "assets"}
}
