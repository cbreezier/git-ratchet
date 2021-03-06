package cmd

import (
	"encoding/csv"
	"github.com/iangrunert/git-ratchet/store"
	log "github.com/spf13/jwalterweatherman"
	"io"
	"strconv"
)

func Dump(prefix string, output io.Writer) int {
	log.INFO.Println("Reading measures stored in git")
	gitlog := store.CommitMeasureCommand(prefix)

	readStoredMeasure, err := store.CommitMeasures(gitlog)
	if err != nil {
		log.FATAL.Println(err)
		return 20
	}

	for {
		cm, err := readStoredMeasure()

		// Empty state of the repository - no stored metrics.
		if err == io.EOF {
			break
		} else if err != nil {
			log.FATAL.Println(err)
			return 40
		}

		out := csv.NewWriter(output)

		for _, measure := range cm.Measures {
			out.Write([]string{cm.Timestamp.String(), measure.Name, strconv.Itoa(measure.Value), strconv.Itoa(measure.Baseline)})
		}
		out.Flush()
	}

	log.INFO.Println("Finished reading measures stored in git")
	return 0
}
