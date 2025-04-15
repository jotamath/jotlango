# Math Library

The `math` library provides mathematical constants and functions for JotLang.

## Import

```jt
import "math"
```

## Constants

| Name | Type | Description |
|------|------|-------------|
| `PI` | `float` | Mathematical constant π (pi) |
| `E` | `float` | Mathematical constant e (Euler's number) |

## Functions

### Basic Operations

```jt
fn sqrt(x float) : float
```
Returns the square root of x.

```jt
fn pow(base float, exponent float) : float
```
Returns base raised to the power of exponent.

```jt
fn abs(x float) : float
```
Returns the absolute value of x.

### Trigonometric Functions

```jt
fn sin(x float) : float
```
Returns the sine of x (x in radians).

```jt
fn cos(x float) : float
```
Returns the cosine of x (x in radians).

```jt
fn tan(x float) : float
```
Returns the tangent of x (x in radians).

### Rounding Functions

```jt
fn round(x float) : int
```
Returns x rounded to the nearest integer.

```jt
fn ceil(x float) : int
```
Returns the smallest integer greater than or equal to x.

```jt
fn floor(x float) : int
```
Returns the largest integer less than or equal to x.

### Statistical Functions

```jt
fn max(x float, y float) : float
```
Returns the larger of x and y.

```jt
fn min(x float, y float) : float
```
Returns the smaller of x and y.

## Examples

```jt
import "math"
import "io"

// Basic operations
root = math.sqrt(16)
io.println("Square root of 16: {root}")  // Output: 4

power = math.pow(2, 3)
io.println("2^3: {power}")  // Output: 8

// Trigonometry
angle = math.PI / 4  // 45 degrees
sin = math.sin(angle)
io.println("sin(45°): {sin}")  // Output: ~0.7071

// Rounding
x = 3.7
rounded = math.round(x)
io.println("3.7 rounded: {rounded}")  // Output: 4

// Statistics
max = math.max(10, 20)
io.println("Maximum: {max}")  // Output: 20
``` 