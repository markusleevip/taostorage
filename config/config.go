package config

const (
	TempDir       = "/tmp/"
	MaxUploadSize = 1024 * 1024 * 1024 * 2 // 2GB
)
type Profile struct {
	AlbumPath string
}
var (
	PFile Profile
)