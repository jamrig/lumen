package lumen

type HitRecord struct {
	P         Vec3
	T         float64
	Normal    Vec3
	FrontFace bool
}

func NewHitRecord(p Vec3, t float64, r Ray, outwardNormal Vec3) *HitRecord {
	record := &HitRecord{
		P: p,
		T: t,
	}

	record.FrontFace = r.Direction.Dot(outwardNormal) < 0
	record.Normal = outwardNormal
	if !record.FrontFace {
		record.Normal = record.Normal.Mul(-1.0)
	}

	return record
}

type Hittable interface {
	Hit(r Ray, tMin, tMax float64) *HitRecord
}

type HittableList struct {
	Objects []Hittable
}

func NewHittableList() HittableList {
	return HittableList{}
}

func (h *HittableList) Clear() {
	h.Objects = make([]Hittable, 0)
}

func (h *HittableList) Add(object Hittable) {
	h.Objects = append(h.Objects, object)
}

func (h *HittableList) Hit(r Ray, tMin, tMax float64) *HitRecord {
	closestSoFar := tMax
	var record *HitRecord

	for _, obj := range h.Objects {
		if newRecord := obj.Hit(r, tMin, closestSoFar); newRecord != nil {
			record = newRecord
			closestSoFar = record.T
		}
	}

	return record
}
