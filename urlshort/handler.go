package urlshort

import (
	"fmt"
	"net/http"

	"github.com/go-yaml/yaml"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		URL := pathsToUrls[r.URL.Path]
		if URL != "" {
			http.Redirect(w, r, URL, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

//YAMLMap is built to store path and urls
type YAMLMap struct {
	path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathsUrls, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	mapToUrls := mapBuilder(pathsUrls)

	return MapHandler(mapToUrls, fallback), nil
}

func parseYAML(ymlInput []byte) (pathsUrls []YAMLMap, err error) {
	if len(ymlInput) == 0 {
		fmt.Println("YAML file is empty.")
		return
	}
	err = yaml.Unmarshal(ymlInput, &pathsUrls)
	return
}

func mapBuilder(ymlMap []YAMLMap) map[string]string {
	if len(ymlMap) == 0 || ymlMap == nil {
		return nil
	}
	mapToUrls := make(map[string]string)
	for _, v := range ymlMap {
		mapToUrls[v.path] = v.URL
	}
	return mapToUrls
}
