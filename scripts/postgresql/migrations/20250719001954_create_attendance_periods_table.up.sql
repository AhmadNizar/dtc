CREATE TABLE attendance_periods (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,

    created_by UUID,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),

    CONSTRAINT fk_attendance_period_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL
);
