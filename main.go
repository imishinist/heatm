package heatmtool

import (
	"fmt"
)

type strs []string

func (s *strs) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *strs) Set(v string) error {
	*s = append(*s, v)
	return nil
}

var (
	files strs
)

/*
func main() {
	flag.Var(&files, "file", "")
	flag.Parse()

	if len(files) < 1 {
		return
	}

	heatmaps := make([]*HeatMap, 0, len(files))
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		var hm HeatMap
		if err := json.NewDecoder(f).Decode(&hm); err != nil {
			log.Fatal(err)
		}
		heatmaps = append(heatmaps, &hm)
	}

	dist := heatmaps[0]
	for i := 1; i < len(heatmaps); i++ {
		if err := dist.Add(heatmaps[i]); err != nil {
			log.Fatal(err)
		}
	}

	json.NewEncoder(os.Stdout).Encode(dist)
}
*/
