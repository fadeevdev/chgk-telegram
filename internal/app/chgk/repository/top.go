package repository

import (
	"context"
	"fmt"
	"github.com/lib/pq"
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
					array_length(answered_questions, 1) as questions,
        			row_number() over(
          		 order by array_length(answered_questions, 1) desc
        		) as position 
				from correct_answers
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
		insert into correct_answers (
			user_id,
			answered_questions
		) VALUES (
			$1, $2
		) ON CONFLICT (user_id) do
			update set answered_questions = array_append(correct_answers.answered_questions, $2);
	`
	_, err = r.pool.Exec(ctx, query,
		uID,
		pq.Array(qID),
	)

	fmt.Println(err)

	return
}
