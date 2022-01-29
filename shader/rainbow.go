package shader

var rainbowShader = []byte(`
package main

// Uniform variables.
var Time float
var Cursor vec2
var ScreenSize vec2

// Fragment is the entry point of the fragment shader.
// Fragment returns the color value for the current position.
func Fragment(position vec4, texCoord vec2, color vec4) vec4 {

	// Normalized pixel coordinates (from 0 to 1)
	uv := (position.xy-.5*ScreenSize)/ScreenSize.y

	//r := fract(sin(Time+uv.x*332.0+uv.y*5532.0)*3442.0)

	// Time varying pixel color
	col := 0.5 + 0.5*cos(Time+uv.xyx+vec3(0,2,4))
	//col = col + imageSrc0At(texCoord).xyz
	//col = col*sin(r)

	// You can treat multiple source images by
	// imageSrc[N]At or imageSrc[N]UnsafeAt.
	return vec4(col,1.0)
}

`)
