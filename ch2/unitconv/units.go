package main

import "fmt"

type Foot float64
type Meter float64
type Inch float64

type Pound float64
type Kilogram float64

const LbsToKg Pound = 0.4535924
const KgToLbs Kilogram = 2.204623

const InToFt Inch = 12.0
const FtToM Foot = 3.28084
const MToFt Meter = 0.3048

// Lengths

func (f Foot) String() string {
  return fmt.Sprintf("%gft", f)
}

func (m Meter) String() string {
  return fmt.Sprintf("%gm", m)
}

func (i Inch) String() string {
  return fmt.Sprintf("%gin", i)
}

// Weights

func (lb Pound) String() string {
  return fmt.Sprintf("%glb", lb)
}

func (k Kilogram) String() string {
  return fmt.Sprintf("%gkg", k)
}
