package doctor

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// WriteText renders the report for a terminal.
func WriteText(w io.Writer, r *Report) {
	fmt.Fprintf(w, "gateway doctor — %s\n\n", r.StartedAt.Format(time.RFC3339))
	for _, res := range r.Results {
		fmt.Fprintf(w, "[%-4s] %-24s %s\n", res.Status, res.Check, res.Detail)
	}
	fmt.Fprintf(w, "\nResumo: %d OK, %d WARN, %d FAIL, %d INFO. Tabelas verificadas: %d/%d\n",
		r.Summary[OK], r.Summary[Warn], r.Summary[Fail], r.Summary[Info], r.TablesChecked, r.TablesExpected)
}

// WriteJSON renders the report as indented JSON.
func WriteJSON(w io.Writer, r *Report) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(r)
}
