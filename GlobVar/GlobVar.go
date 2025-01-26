package GlobVar

type Room struct {
	Links        []string
	IsChecked    bool
	BeforeInPath string
}

var (
	AntsNumber    int
	OriginalRooms = make(map[string]Room)
	Rooms         = make(map[string]Room)
	Start         string
	End           string
	ValidPaths    [][]string
	AllValidPaths [][][]string
)