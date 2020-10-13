use std::fmt;

#[derive(Debug, PartialEq)]
pub struct Clock {
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let hours = hours.rem_euclid(24);
        let minutes = minutes + 60 * hours;
        Clock { minutes }
    }

    pub fn new_minutes_only(minutes: i32) -> Self{
        Clock{ minutes }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        let total_minutes = self.minutes + minutes;
        return Clock { minutes: total_minutes };
    }

    fn get_hours_and_minutes(&self) -> (i32, i32) {
        let minutes = self.minutes.rem_euclid(60);
        let hours = self.minutes.div_euclid(60).rem_euclid(24);

        return (hours, minutes);
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let (hours, minutes) = self.get_hours_and_minutes();
        write!(f, "{:02}:{:02}", hours, minutes)
    }
}
