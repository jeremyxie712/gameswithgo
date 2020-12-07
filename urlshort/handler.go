package urlshort

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/boltdb/bolt"
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
	ymlParsed, err := yamlParser(yml)
	if err != nil {
		return nil, err
	}
	builtMap := mapBuilder(ymlParsed)
	return MapHandler(builtMap, fallback), err
}

func jsonParser(jsonData []byte) (jsonParsed []pathUrl, err error) {
	if len(jsonData) == 0 {
		err := errors.New("JSON is empty")
		return jsonParsed, err
	}
	err = json.Unmarshal(jsonData, &jsonParsed)
	return
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//     [
//	     {
//         "path": "/some-path",
//         "url": "https://www.some-url.com/demo"
//       }
//	   ]
//
// The only errors that can be returned all related to having
// invalid JSON data.
func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	jsonParsed, err := jsonParser(json)
	if err != nil {
		return nil, err
	}
	mapJSON := mapBuilder(jsonParsed)
	return MapHandler(mapJSON, fallback), err
}

//BoltDBHandler will use the given Bolt DB and return an http.HandlerFunc
//that will attempt to map any paths to their corresponding URL.
//If the path is not provided in the DB, then the fallback http.Handler
//will be called instead.
func BoltDBHandler(db *bolt.DB, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var url string
		err := db.View(func(tx *bolt.Tx) error {
			buc := tx.Bucket([]byte("paths"))
			bucGet := buc.Get([]byte(r.URL.Path))
			if bucGet != nil {
				url = string(bucGet)
			}
			return nil
		})
		if err != nil || url == "" {
			fallback.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, url, http.StatusFound)
		}
	})
}
