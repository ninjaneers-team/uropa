package solver

import (
	"github.com/ninjaneers-team/uropa/crud"
	"github.com/ninjaneers-team/uropa/diff"
	"github.com/ninjaneers-team/uropa/opa"
	"github.com/ninjaneers-team/uropa/print"
	"github.com/ninjaneers-team/uropa/state"
)

// Stats holds the stats related to a Solve.
type Stats struct {
	CreateOps int
	UpdateOps int
	DeleteOps int
}

// Solve generates a diff and walks the graph.
func Solve(doneCh chan struct{}, syncer *diff.Syncer,
	client *opa.Client, parallelism int, dry bool) (Stats, []error) {
	var r *crud.Registry

	r = buildRegistry(client)

	var stats Stats
	recordOp := func(op crud.Op) {
		switch op {
		case crud.Create:
			stats.CreateOps = stats.CreateOps + 1
		case crud.Update:
			stats.UpdateOps = stats.UpdateOps + 1
		case crud.Delete:
			stats.DeleteOps = stats.DeleteOps + 1
		}
	}

	errs := syncer.Run(doneCh, parallelism, func(e diff.Event) (crud.Arg, error) {
		var err error
		var result crud.Arg

		c := e.Obj.(state.ConsoleString)
		switch e.Op {
		case crud.Create:
			print.CreatePrintln("creating", e.Kind, c.Console())
		case crud.Update:
			diffString, err := getDiff(e.OldObj, e.Obj)
			if err != nil {
				return nil, err
			}
			print.UpdatePrintln("updating", e.Kind, c.Console(), diffString)
		case crud.Delete:
			print.DeletePrintln("deleting", e.Kind, c.Console())
		default:
			panic("unknown operation " + e.Op.String())
		}

		if !dry {
			// sync mode
			// fire the request to Opa
			result, err = r.Do(e.Kind, e.Op, e)
			if err != nil {
				return nil, err
			}
		} else {
			// diff mode
			// return the new obj as is
			result = e.Obj
		}
		// record operation in both: diff and sync commands
		recordOp(e.Op)

		return result, nil
	})
	return stats, errs
}

func buildRegistry(client *opa.Client) *crud.Registry {
	var r crud.Registry
	r.MustRegister("policy", &policyCRUD{client: client})
	return &r
}
