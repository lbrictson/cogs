package pkg

import (
	"context"
	"github.com/lbrictson/cogs/ent"
	"time"
)

func runErroredJobCleaner(ctx context.Context, db *ent.Client) {
	for {
		f := false
		LogFromCtx(ctx).Info("Running job sweep")
		r := "running"
		e := "job expired - unknown error"
		finished := "finished"
		input := QueryHistoriesInput{
			Limit:  50,
			Offset: 0,
			Status: &r,
		}
		nonFinishedHistories, err := queryHistories(ctx, db, input)
		if err != nil {
			LogFromCtx(ctx).Error(err.Error())
			time.Sleep(5 * time.Minute)
			continue
		}
		for _, history := range nonFinishedHistories {
			if history.CreatedAt.Before(time.Now().Add(-1 * time.Hour)) {
				LogFromCtx(ctx).Info("Marking job as errored: " + history.RunID)
				dur := int(time.Now().Sub(history.CreatedAt).Seconds())
				i := UpdateHistoryInput{
					Success:         &f,
					Status:          &finished,
					Output:          &e,
					DurationSeconds: &dur,
				}
				_, err := updateHistory(ctx, db, history.ID, i)
				if err != nil {
					LogFromCtx(ctx).Error(err.Error())
				}
			}
		}
		time.Sleep(5 * time.Minute)
	}
}
