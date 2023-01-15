package badger

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

type ZeroLogAdpater struct{}

func (adpater *ZeroLogAdpater) Errorf(format string, v ...interface{}) {
	log.Error().Msg(strings.Trim(fmt.Sprintf(format, v...), "\n"))
}

func (adpater *ZeroLogAdpater) Warningf(format string, v ...interface{}) {
	log.Warn().Msg(strings.Trim(fmt.Sprintf(format, v...), "\n"))
}

func (adpater *ZeroLogAdpater) Infof(format string, v ...interface{}) {
	log.Info().Msg(strings.Trim(fmt.Sprintf(format, v...), "\n"))
}

func (adpater *ZeroLogAdpater) Debugf(format string, v ...interface{}) {
	log.Debug().Msg(strings.Trim(fmt.Sprintf(format, v...), "\n"))
}
