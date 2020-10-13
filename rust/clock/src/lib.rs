use std::fmt;

#[derive(Debug, PartialEq)]
pub struct Clock {
    hour: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let (converted_hours, converted_minutes) = Clock::convert_times(hours, minutes);

        Clock {
            hour: converted_hours,
            minutes: converted_minutes,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        let (new_hours, new_minutes) = Clock::convert_times(self.hour, self.minutes + minutes);

        return Clock {
            hour: new_hours,
            minutes: new_minutes,
        };
    }

    /// Given a time in the format hours:minutes, returns the same time but
    /// within the usual limits
    fn convert_times(hours: i32, minutes: i32) -> (i32, i32) {
        let mut converted_hours = hours;
        let mut converted_minutes = minutes;

        // Checking for negative values
        if converted_hours < 0 {
            return Clock::convert_times(24 - ((-hours) % 24), converted_minutes);
        }
        if converted_minutes < 0 {
            return Clock::convert_times(hours - 1, 60 - ((-minutes) % 60));
        }

        // Cheking for values too big
        if minutes >= 60 {
            converted_hours = converted_hours + (converted_minutes / 60);
            converted_minutes = converted_minutes % 60;
        }

        if converted_hours >= 24 {
            converted_hours = converted_hours % 24
        }

        return (converted_hours, converted_minutes);
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hour, self.minutes)
    }
}
