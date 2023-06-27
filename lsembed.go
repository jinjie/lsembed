package lsembed

import (
	"context"
	"log"
	"os"

	"github.com/benbjohnson/litestream"
)

var ctx = context.Background()

func Replicate(r *litestream.Replica) (*litestream.DB, error) {
	lsdb := r.DB()
	lsdb.Replicas = append(lsdb.Replicas, r)

	if err := Restore(r); err != nil {
		return nil, err
	}

	if err := lsdb.Open(); err != nil {
		return nil, err
	}

	return lsdb, nil
}

func Restore(r *litestream.Replica) (err error) {
	// skip restore if local database already exists
	if _, err := os.Stat(r.DB().Path()); err == nil {
		log.Println("local database already exists, skipping restore")
		return nil
	}

	opt := litestream.NewRestoreOptions()
	opt.OutputPath = r.DB().Path()

	if opt.Generation, _, err = r.CalcRestoreTarget(ctx, opt); err != nil {
		return err
	}

	if opt.Generation == "" {
		log.Println("no generation to restore")
	}

	log.Printf("restoring %s to %s", opt.Generation, opt.OutputPath)
	if err := r.Restore(ctx, opt); err != nil {
		return err
	}

	log.Println("restore complete")
	return nil
}
