// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package sqlc

import (
	"context"
)

const listUsers = `-- name: ListUsers :many
SELECT user_id, username, name, language, is_enabled, is_enabled_for_autotrading, settings_default_chain, date_creation, referralcode FROM telegram_user
`

func (q *Queries) ListUsers(ctx context.Context) ([]TelegramUser, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TelegramUser
	for rows.Next() {
		var i TelegramUser
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Name,
			&i.Language,
			&i.IsEnabled,
			&i.IsEnabledForAutotrading,
			&i.SettingsDefaultChain,
			&i.DateCreation,
			&i.Referralcode,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
