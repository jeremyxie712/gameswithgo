package urlshort

import (
	"errors"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if _, ok := pathsToUrls[url]; ok {
			http.Redirect(w, r, pathsToUrls[url], http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// Type of path that stores the path and url information for yaml
type pathUrl struct {
	path string `yaml:"path"`
	url  string `yaml:"url"`
}

func yamlParser(ymlData []byte) (ymlParsed []pathUrl, err error) {
	if len(ymlData) == 0 {
		err := errors.New("YAML is empty")
		return nil, err
	}
	err = yaml.Unmarshal(ymlData, &ymlParsed)
	return
}

func mapBuilder(ymlData []pathUrl) map[string]string {
	if len(ymlData) == 0 {
		return nil
	}
	m := make(map[string]string)
	for _, v := range ymlData {
		m[v.path] = v.url
	}
	return m
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
	if len(yml) == 0 || yml == nil {
		return nil, errors.New("YAML data is empty")
	}
	ymlParsed, err := yamlParser(yml)
	builtMap := mapBuilder(ymlParsed)
	return MapHandler(builtMap, fallback), err
}
