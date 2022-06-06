package repository

import (
	"context"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/models"
)

func (r *repository) GetTopPlayers(ctx context.Context, n uint64) ([]models.User, error) {
	//TBD
	return nil, nil
}

func (r *repository) GetTopPosition(ctx context.Context, uID uint64) (position models.TopUser, err error) {
	const query = `
		select position, questions 
			from (
   				select *,
					array_length(answered_questions) as questions,
        			row_number() over(
          		 order by array_length(answered_questions) desc
        		) as position 
				from answered_questions
		) result 
		where user_id = $1;
	`
	err = r.pool.QueryRow(ctx, query,
		uID,
	).Scan(&position.Position, &position.Questions)
	u, err := r.GetUser(ctx, uID)
	if err != nil {
		return
	}

	position.FirstName = u.FirstName

	return
}

func (r *repository) AddToTop(ctx context.Context, uID uint64, qID uint64) (err error) {
	const query = `
		update correct_answers (
			id,
			answered_questions
		) VALUES (
			$1, $2
		) ON CONFLICT (id) do
			update set answered_questions = array_prepend(answered_questions, $2);
	`
	r.pool.QueryRow(ctx, query,
		uID,
		qID,
	)

	return
}
