CREATE TABLE reimbursements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    amount NUMERIC(15,2) NOT NULL,
    description TEXT,
    date DATE NOT NULL,

    created_by UUID,
    updated_by UUID,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    ip_address VARCHAR(45),

    CONSTRAINT fk_reimbursements_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_reimbursements_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT fk_reimbursements_updated_by FOREIGN KEY (updated_by) REFERENCES users(id) ON DELETE SET NULL
);
