CREATE TABLE payslips (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    period_id UUID NOT NULL,

    total_attendance_days INT NOT NULL,
    total_attendance_salary NUMERIC(15,2) NOT NULL,
    total_overtime_hours NUMERIC(10,2) NOT NULL,
    total_overtime_salary NUMERIC(15,2) NOT NULL,
    total_reimbursement NUMERIC(15,2) NOT NULL,
    total_take_home NUMERIC(15,2) NOT NULL,

    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    created_by UUID,
    updated_by UUID,
    request_ip INET,
    request_id TEXT,

    CONSTRAINT fk_payslips_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_payslips_period FOREIGN KEY (period_id) REFERENCES attendance_periods(id) ON DELETE CASCADE,
    CONSTRAINT fk_payslips_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT fk_payslips_updated_by FOREIGN KEY (updated_by) REFERENCES users(id) ON DELETE SET NULL
);
