package loaders

import "github.com/mniak/biblia/pkg/bible"

type Loader interface {
	Load() (bible.Testament, error)
}
