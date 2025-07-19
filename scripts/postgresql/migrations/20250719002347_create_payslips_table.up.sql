CREATE TABLE payslips (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    period_id UUID NOT NULL REFERENCES attendance_periods(id),
    total_attendance_days INTEGER NOT NULL,
    total_attendance_salary NUMERIC(15, 2) NOT NULL,
    total_overtime_hours NUMERIC(10, 2) NOT NULL,
    total_overtime_salary NUMERIC(15, 2) NOT NULL,
    total_reimbursement NUMERIC(15, 2) NOT NULL,
    total_take_home NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (user_id, period_id)
);