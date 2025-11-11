-- FOR TESTING
CREATE TABLE profiles (
    id      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID DEFAULT gen_random_uuid(),

    CONSTRAINT unique_user_id UNIQUE (user_id)
);

CREATE TABLE venues (
    id            SERIAL PRIMARY KEY,
    name          TEXT NOT NULL,
    address       TEXT NOT NULL,
    city          TEXT NOT NULL,
    neighbourhood TEXT,
    borough       TEXT,
    type          TEXT,
    contact       TEXT,

    CONSTRAINT unique_venue UNIQUE (name, address, city, neighbourhood, borough, type, contact)
);

CREATE TABLE venue_rooms (
    id       SERIAL PRIMARY KEY,
    venue_id INTEGER NOT NULL,
    name     TEXT,

    CONSTRAINT fk_venue FOREIGN KEY (venue_id) REFERENCES venues(id) ON DELETE CASCADE,
    CONSTRAINT unique_venue_room UNIQUE (venue_id, name)
);

CREATE TYPE week_day_in_a_month_rule AS (
    weeks_in_advance      INTEGER,
    weeks_of_month_policy BOOLEAN[5],
    week_day              INTEGER,
    start_time            TIMESTAMPTZ,
    end_time              TIMESTAMPTZ
);

-- General information about mics that can be scheduled following a pattern (occurrence rules)
CREATE TABLE mic_definitions (
    id                   SERIAL PRIMARY KEY,
    venue_room_id        INTEGER NOT NULL,
    name                 TEXT NOT NULL,
    signup_instructions  TEXT,
    notes                TEXT,
    audience_cost        TEXT NOT NULL,
    comedian_cost        TEXT NOT NULL,
    stage_time           TEXT,
    host                 TEXT,
    instagram            TEXT, -- changes and update go here
    sms                  TEXT,
    verified             TEXT,
    active               BOOLEAN NOT NULL,
    occurrence_rules     week_day_in_a_month_rule[] NOT NULL,

    CONSTRAINT fk_venue_room FOREIGN KEY (venue_room_id) REFERENCES venue_rooms(id) ON DELETE CASCADE
);

-- Mic events that are scheduled to happen
CREATE TABLE scheduled_mic_events (
    id                  SERIAL PRIMARY KEY,
    definition_id       INTEGER,
    venue_room_id       INTEGER NOT NULL,
    name                TEXT NOT NULL,
    signup_instructions TEXT,
    notes               TEXT,
    audience_cost       TEXT NOT NULL,
    comedian_cost       TEXT NOT NULL,
    stage_time          TEXT,
    host                TEXT,
    instagram           TEXT,
    sms                 TEXT,
    comedian_lineup     TEXT[] NOT NULL,
    start_date_and_time TIMESTAMPTZ NOT NULL,
    end_date_and_time   TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_definition FOREIGN KEY (definition_id) REFERENCES mic_definitions(id) ON DELETE RESTRICT,
    CONSTRAINT fk_venue_room FOREIGN KEY (venue_room_id) REFERENCES venue_rooms(id) ON DELETE RESTRICT,
    CONSTRAINT unique_definition_venue_room_start_date_and_time UNIQUE (
        definition_id, venue_room_id, start_date_and_time
    )
); -- Delete rows if the definitions is inactive

-- Mic events that have already happened and are logged here
CREATE TABLE past_mic_events (
    id                  SERIAL PRIMARY KEY,
    venue_room_id       INTEGER NOT NULL,
    name                TEXT NOT NULL,
    host                TEXT,
    audience_cost       TEXT NOT NULL,
    comedian_cost       TEXT NOT NULL,
    stage_time          TEXT,
    comedian_lineup     TEXT[] NOT NULL,
    start_date_and_time TIMESTAMPTZ NOT NULL,
    end_date_and_time   TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_venue_room FOREIGN KEY (venue_room_id) REFERENCES venue_rooms(id) ON DELETE RESTRICT
);

CREATE TABLE personal_mic_event_notes (
    id                 SERIAL PRIMARY KEY,
    comedian_id        UUID,
    past_mic_event_id  INTEGER,
    comedian_notes     TEXT,
    video_url          TEXT,

    CONSTRAINT fk_comedian_id FOREIGN KEY (comedian_id) REFERENCES profiles(user_id) ON DELETE CASCADE,
    CONSTRAINT past_mic_events FOREIGN KEY (past_mic_event_id) REFERENCES past_mic_events(id) ON DELETE CASCADE
);
