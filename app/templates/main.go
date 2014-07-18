package templates

func ParseTemplateBinaries() []string {
	var parsed []string
	for _, name := range AssetNames() {
		asset, err := Asset(name)
		if err != nil {
			panic(err)
		}
		parsed = append(parsed, string(asset[:]))
	}
	return parsed
}
