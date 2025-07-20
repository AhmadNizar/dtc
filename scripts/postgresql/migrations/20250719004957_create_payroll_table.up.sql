CREATE TABLE payrolls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    period_id UUID NOT NULL,

    base_salary NUMERIC(12,2) NOT NULL,
    overtime_pay NUMERIC(12,2) NOT NULL,
    bonus NUMERIC(12,2) NOT NULL,
    deductions NUMERIC(12,2) NOT NULL,
    net_pay NUMERIC(12,2) NOT NULL,
    approved BOOLEAN DEFAULT FALSE,
    generated_at TIMESTAMPTZ NOT NULL,

    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    created_by UUID,
    updated_by UUID,
    request_ip INET,
    request_id TEXT,

    CONSTRAINT fk_payrolls_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_payrolls_period FOREIGN KEY (period_id) REFERENCES attendance_periods(id) ON DELETE CASCADE,
    CONSTRAINT fk_payrolls_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT fk_payrolls_updated_by FOREIGN KEY (updated_by) REFERENCES users(id) ON DELETE SET NULL
);
