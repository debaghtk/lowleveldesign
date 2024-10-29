# Open-Closed Principle Problem Statement

## Background
A shipping company needs to calculate shipping costs for different types of packages. Currently, they handle regular ground shipping and express shipping.

## Problem
The company's existing shipping calculator looks like this:

```go
type Package struct {
    weight     float64
    distance   float64
    isExpress  bool
}

type ShippingCalculator struct {}

func (sc *ShippingCalculator) CalculateCost(pkg Package) float64 {
    if pkg.isExpress {
        return pkg.weight * 2.5 * pkg.distance * 1.5  // Express rate
    }
    return pkg.weight * 1.5 * pkg.distance  // Regular rate
}
```

## Issues
1. Every time a new shipping method is added (like overnight, international, or drone delivery), the `CalculateCost` method needs to be modified.
2. The code becomes increasingly complex with more if-else statements.
3. Changes to existing shipping calculations risk breaking working code.
4. Testing becomes more difficult as the number of shipping methods grows.

## Challenge
Refactor the shipping calculator to follow the Open-Closed Principle so that:
1. Adding new shipping methods doesn't require modifying existing code
2. Each shipping method can have its own calculation logic
3. The system is extensible for future shipping methods
4. The code is easier to test and maintain

## Requirements
1. Support the existing regular and express shipping methods
2. Make it easy to add new shipping methods
3. Each shipping method should be able to define its own cost calculation formula
4. Maintain clean separation of concerns
5. Ensure the solution is testable

## Bonus Challenge
Add two new shipping methods (international shipping and overnight delivery) to demonstrate how your solution supports extension without modification.