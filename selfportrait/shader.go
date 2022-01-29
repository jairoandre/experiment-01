package selfportrait

var shaderCode = []byte(`
package main

// Uniform variables.
var Time float
var Cursor vec2
var ScreenSize vec2
var Gradient [256*3]float

// Fragment is the entry point of the fragment shader.
// Fragment returns the color value for the current position.
func Fragment(position vec4, texCoord vec2, color vec4) vec4 {

	// Normalized pixel coordinates (from 0 to 1)
	//uv := position.xy/ScreenSize

	//col := 0.5 + 0.5*cos(Time+uv.xyx+vec3(0,2,4))
	col := imageSrc0At(texCoord).xyz

	return vec4(col,1.0)
}

`)
