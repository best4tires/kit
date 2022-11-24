// Package env provides a uniform way of dealing with environment such as .env files, os.Environ and (command line) flags.
// The goal is, that applications don't have to care about the source from a variable but just handle the values.
package env

// merge merges "from" env into "to" env, keeping already existing values
func merge(from map[string]any, to map[string]any) {
	for k, v := range from {
		if _, ok := to[k]; !ok {
			to[k] = v
		}
	}
}
