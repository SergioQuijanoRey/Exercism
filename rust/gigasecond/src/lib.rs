use chrono::{DateTime, Utc};

// Returns a Utc DateTime one billion seconds after start.
pub fn after(start: DateTime<Utc>) -> DateTime<Utc> {
    let one_billion = 1_000_000_000;
    let one_billion_seconds = chrono::Duration::seconds(one_billion);

    return start + one_billion_seconds;
}
