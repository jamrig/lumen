package lumen

import (
	"encoding/json"
	"fmt"
)

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
	Hit(r Ray, t Interval) *HitRecord
}

type HittableList struct {
	Objects []Hittable
}

func NewHittableList() HittableList {
	return HittableList{}
}

func (h HittableList) String() string {
	pretty, _ := json.MarshalIndent(h, "", "  ")
	return fmt.Sprintf("HittableList: %v", string(pretty))
}

func (h *HittableList) Clear() {
	h.Objects = make([]Hittable, 0)
}

func (h *HittableList) Add(object Hittable) {
	h.Objects = append(h.Objects, object)
}

func (h *HittableList) Hit(r Ray, t Interval) *HitRecord {
	closest := NewInterval(t.Min, t.Max)
	var record *HitRecord

	for _, obj := range h.Objects {
		if newRecord := obj.Hit(r, closest); newRecord != nil {
			record = newRecord
			closest.Max = record.T
		}
	}

	return record
}
