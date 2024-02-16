package main

//
func InchToFoot(i Inch) Foot {
  return Foot(i / InToFt)
}

func InchToMeter(i Inch) Meter {
  return FootToMeter(InchToFoot(i))
}

// FootToMeter converts a length in Feet to Meters
func FootToMeter(f Foot) Meter {
  return Meter(f * FtToM)
}

// MeterToFoot converts a length in Meters to Feet
func MeterToFoot(m Meter) Foot {
  return Foot(m * MToFt)
}

// PoundToKilogram converts a weight in Pounds to Kilograms
func PoundToKilogram(p Pound) Kilogram {
  return Kilogram(p * LbsToKg)
}

// KilogramToPound converts a weight in Kilograms to Pounds
func KilogramToPound(k Kilogram) Pound {
  return Pound(k * KgToLbs)
}
