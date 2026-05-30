-- name: CreateTransaction :one
INSERT INTO transactions (
    user_id,
    category_id,
    type,
    amount,
    note,
    transaction_date
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetTransactionByID :one
SELECT *
FROM transactions
WHERE id = $1
AND user_id = $2;

-- name: ListTransactions :many
SELECT
    t.id,
    t.user_id,
    t.category_id,
    c.name AS category_name,
    t.type,
    t.amount,
    t.note,
    t.transaction_date,
    t.created_at
FROM transactions t
JOIN categories c
ON c.id = t.category_id
WHERE t.user_id = $1
ORDER BY t.transaction_date DESC;

-- name: UpdateTransaction :one
UPDATE transactions
SET
    category_id = $3,
    type = $4,
    amount = $5,
    note = $6,
    transaction_date = $7,
    updated_at = NOW()
WHERE id = $1
AND user_id = $2
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1
AND user_id = $2;

-- name: GetSummary :one
SELECT
    COALESCE(
        SUM(CASE WHEN type = 'income' THEN amount END),
        0
    ) AS total_income,

    COALESCE(
        SUM(CASE WHEN type = 'expense' THEN amount END),
        0
    ) AS total_expense

FROM transactions
WHERE user_id = $1;

-- name: GetMonthlySummary :many
SELECT
    DATE_TRUNC('month', transaction_date) AS month,

    COALESCE(
        SUM(CASE WHEN type = 'income' THEN amount END),
        0
    ) AS income,

    COALESCE(
        SUM(CASE WHEN type = 'expense' THEN amount END),
        0
    ) AS expense

FROM transactions
WHERE user_id = $1
GROUP BY month
ORDER BY month DESC;

-- name: GetCategoryBreakdown :many
SELECT
    c.name,
    SUM(t.amount) AS total
FROM transactions t
JOIN categories c
ON c.id = t.category_id
WHERE
    t.user_id = $1
    AND t.type = 'expense'
GROUP BY c.name
ORDER BY total DESC;