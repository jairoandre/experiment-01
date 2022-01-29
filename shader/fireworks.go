package shader

var fireWorks = []byte(`
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

	col := vec3(0)
	brightness := .004

	for i := 0; i < 20; i++ {
		dir := Hash12(i)-.5
		d := length(uv-dir*Time*0.4)
		col += brightness/d
	}

	// You can treat multiple source images by
	// imageSrc[N]At or imageSrc[N]UnsafeAt.
	return vec4(col,1.0)
}

func Hash12(t float) vec2 {
	x := fract(sin(t*674.3)*453.2)
	y := fract(sin((t+x)*714.3)*263.2)
	return vec2(x, y)
}

`)
