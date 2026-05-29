#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum Unit {
    Temp(TempUnit),
    Length(LengthUnit),
}
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TempUnit {
    Celsius,
    Fahrenheit,
    Kelvin,
}
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum LengthUnit {
    Centimeter,
    Meter,
    Kilometer,
    Mile,
}

pub fn parse_unit(s: &str) -> Result<Unit, String> {
    match s {
        "C" => Ok(Unit::Temp(TempUnit::Celsius)),
        "F" => Ok(Unit::Temp(TempUnit::Fahrenheit)),
        "K" => Ok(Unit::Temp(TempUnit::Kelvin)),
        "cm" => Ok(Unit::Length(LengthUnit::Centimeter)),
        "m" => Ok(Unit::Length(LengthUnit::Meter)),
        "km" => Ok(Unit::Length(LengthUnit::Kilometer)),
        "mi" => Ok(Unit::Length(LengthUnit::Mile)),
        _ => Err(format!("unknown unit '{s}'")),
    }
}

pub fn convert(from: &str, to: &str, value: f64) -> Result<f64, String> {
    let from_unit: Unit = parse_unit(from)?;
    let to_unit: Unit = parse_unit(to)?;

    match (from_unit, to_unit) {
        (Unit::Temp(a), Unit::Temp(b)) => Ok(celsius_to_temp(temp_to_celsius(value, a), b)),
        (Unit::Length(a), Unit::Length(b)) => Ok(meters_to_length(length_to_meters(value, a), b)),
        _ => Err(format!(
            "cannot convert from '{from}' to '{to}': different unit families"
        )),
    }
}

fn temp_to_celsius(v: f64, t: TempUnit) -> f64 {
    match t {
        TempUnit::Celsius => v,
        TempUnit::Fahrenheit => (v - 32.0) * 5.0 / 9.0,
        TempUnit::Kelvin => v - 273.15,
    }
}

fn celsius_to_temp(v: f64, t: TempUnit) -> f64 {
    match t {
        TempUnit::Celsius => v,
        TempUnit::Fahrenheit => v * 9.0 / 5.0 + 32.0,
        TempUnit::Kelvin => v + 273.15,
    }
}

fn length_to_meters(v: f64, l: LengthUnit) -> f64 {
    match l {
        LengthUnit::Centimeter => v / 100.0,
        LengthUnit::Meter => v,
        LengthUnit::Kilometer => v * 1000.0,
        LengthUnit::Mile => v * 1609.344,
    }
}
fn meters_to_length(v: f64, l: LengthUnit) -> f64 {
    match l {
        LengthUnit::Centimeter => v * 100.0,
        LengthUnit::Meter => v,
        LengthUnit::Kilometer => v / 1000.0,
        LengthUnit::Mile => v / 1609.344,
    }
}
