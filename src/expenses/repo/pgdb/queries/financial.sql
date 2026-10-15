-- name: GetMonthlySummary :many
WITH monthly_summary AS (
  SELECT
    TO_CHAR(e."ExpDate", 'MM/YYYY') AS month_year,
    DATE_TRUNC('month', e."ExpDate") AS month_sort,
    SUM(CASE WHEN e."TypeID" = 0 THEN ep."Payed" ELSE 0 END)::NUMERIC AS despesa,
    SUM(CASE WHEN e."TypeID" = 1 THEN ep."Payed" ELSE 0 END)::NUMERIC AS salario,
    SUM(CASE WHEN e."TypeID" = 2 THEN ep."Payed" ELSE 0 END)::NUMERIC AS cartao_refeicao,
    SUM(CASE WHEN e."TypeID" = 3 THEN ep."Payed" ELSE 0 END)::NUMERIC AS reembolso,
    SUM(CASE WHEN e."TypeID" = 4 THEN ep."Payed" ELSE 0 END)::NUMERIC AS prenda,
    SUM(CASE WHEN e."TypeID" = 5 THEN ep."Payed" ELSE 0 END)::NUMERIC AS investment,
    SUM(CASE WHEN e."TypeID" = 6 THEN ep."Payed" ELSE 0 END)::NUMERIC AS adjustment,
    SUM(CASE WHEN e."TypeID" = 7 THEN ep."Payed" ELSE 0 END)::NUMERIC AS adjustment_liquid
  FROM "expensesPayments" ep
  JOIN expenses e ON ep."ExpID" = e."ExpID"
  WHERE ep."UserID" = $1
    AND (sqlc.narg(StartDate)::timestamp IS NULL OR e."ExpDate" >= sqlc.narg(StartDate)::timestamp)
    AND (sqlc.narg(EndDate)::timestamp IS NULL OR e."ExpDate" <= sqlc.narg(EndDate)::timestamp)
  GROUP BY
    TO_CHAR(e."ExpDate", 'MM/YYYY'),
    DATE_TRUNC('month', e."ExpDate")
),
calculated_summary AS (
  SELECT
    month_year,
    month_sort,
    despesa,
    salario,
    cartao_refeicao,
    reembolso,
    prenda,
    investment,
    adjustment,
    adjustment_liquid,
    ((salario + cartao_refeicao + reembolso) - (despesa + prenda))::NUMERIC AS poupanca,
    ((salario + cartao_refeicao + reembolso) - (despesa + investment + prenda))::NUMERIC AS poupanca_liquid
  FROM monthly_summary
)
SELECT
  month_year,
  despesa,
  salario,
  cartao_refeicao,
  reembolso,
  prenda,
  investment,
  adjustment,
  adjustment_liquid,
  poupanca,
  poupanca_liquid,
  (SUM(poupanca) OVER (ORDER BY month_sort ASC) + SUM(adjustment) OVER (ORDER BY month_sort ASC) + SUM(adjustment_liquid) OVER (ORDER BY month_sort ASC))::NUMERIC AS running_balance,
  (SUM(poupanca_liquid) OVER (ORDER BY month_sort ASC) + SUM(adjustment_liquid) OVER (ORDER BY month_sort ASC))::NUMERIC AS running_balance_liquid
FROM calculated_summary
ORDER BY month_sort ASC;

-- name: GetMonthlySavingsAverage :one
WITH monthly_summary AS (
  SELECT
    TO_CHAR(e."ExpDate", 'MM/YYYY') AS month_year,
    DATE_TRUNC('month', e."ExpDate") AS month_sort,
    SUM(CASE WHEN e."TypeID" = 0 THEN ep."Payed" ELSE 0 END)::NUMERIC AS despesa,
    SUM(CASE WHEN e."TypeID" = 1 THEN ep."Payed" ELSE 0 END)::NUMERIC AS salario,
    SUM(CASE WHEN e."TypeID" = 2 THEN ep."Payed" ELSE 0 END)::NUMERIC AS cartao_refeicao,
    SUM(CASE WHEN e."TypeID" = 3 THEN ep."Payed" ELSE 0 END)::NUMERIC AS reembolso,
    SUM(CASE WHEN e."TypeID" = 4 THEN ep."Payed" ELSE 0 END)::NUMERIC AS prenda,
    SUM(CASE WHEN e."TypeID" = 5 THEN ep."Payed" ELSE 0 END)::NUMERIC AS investment,
    SUM(CASE WHEN e."TypeID" = 6 THEN ep."Payed" ELSE 0 END)::NUMERIC AS adjustment,
    SUM(CASE WHEN e."TypeID" = 7 THEN ep."Payed" ELSE 0 END)::NUMERIC AS adjustment_liquid
  FROM "expensesPayments" ep
  JOIN expenses e ON ep."ExpID" = e."ExpID"
  WHERE ep."UserID" = 1
  GROUP BY
    TO_CHAR(e."ExpDate", 'MM/YYYY'),
    DATE_TRUNC('month', e."ExpDate")
),
calculated_summary AS (
  SELECT
    ((salario + cartao_refeicao + reembolso) - (despesa + prenda))::NUMERIC AS poupanca,
    ((salario + cartao_refeicao + reembolso) - (despesa + investment + prenda))::NUMERIC AS poupanca_liquid
  FROM monthly_summary
)
SELECT
  TRUNC(AVG(poupanca), 2)::NUMERIC AS avg_poupanca,
  TRUNC(AVG(poupanca_liquid), 2)::NUMERIC AS avg_poupanca_liquid
FROM calculated_summary;
