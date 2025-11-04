package maths

type ScatteredRay struct {
	Ray         Ray
	Attenuation Color
}

func NewScatteredRay(ray Ray, attenuation Color) ScatteredRay {
	return ScatteredRay{
		Ray:         ray,
		Attenuation: attenuation,
	}
}
