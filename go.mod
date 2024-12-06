module github.com/musicpulpite/koanf/v2

go 1.18

require (
	github.com/go-viper/mapstructure/v2 v2.2.1
	github.com/knadh/koanf/maps v0.1.1
	github.com/mitchellh/copystructure v1.2.0
)

require github.com/mitchellh/reflectwalk v1.0.2 // indirect

retract v2.0.2 // Tagged as minor version, but contains breaking changes.
