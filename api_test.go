package cp_test

import (
	. "."
	// "reflect"
	"testing"
)

func tt(t *testing.T, wat interface{}) {
	t.Logf("%T %v", wat, wat)
}

func Test(t *testing.T) {
	gravity := V(0, -100)
	tt(t, gravity)
	// pass
	space := SpaceNew()
	SpaceSetGravity(space, gravity)

	ground := SegmentShapeNew(SpaceGetStaticBody(space), V(-20, -5), V(20, -5), 0)
	ShapeSetFriction(ground, 1)
	SpaceAddShape(space, ground)

	radius := float64(5)
	mass := float64(1)

	moment := MomentForCircle(mass, 0, radius, V(0, 0))

	ballBody := SpaceAddBody(space, BodyNew(mass, moment))
	BodySetPosition(ballBody, V(0, 15))
	ballBody.(SwigcptrStruct_SS_cpBody).SetPosition(0, 15)
	//BodySetMass(ballBody, 2)

	ballShape := SpaceAddShape(space, CircleShapeNew(ballBody, radius, V(0, 0)))
	ShapeSetFriction(ballShape, 0.7)

	timeStep := float64(1.0 / 60.0)

	for time := float64(0); time < 2.0; time += timeStep {
		pos := BodyGetPosition(ballBody)
		vel := BodyGetVelocity(ballBody)
		t.Logf("Time is %5.2f. ballBody is at(%5.2f, %5.2f). It's vel(%5.2f, %5.2f)\n",
			time, pos.GetX(), pos.GetY(), vel.GetX(), vel.GetY())
		SpaceStep(space, timeStep)
	}

	tt(t, space)
	tt(t, ground)
	tt(t, ballBody)
	tt(t, ballShape)
	t.Log(ballBody.Swigcptr())

	SpaceFree(space)
	BodyFree(ballBody)
	ShapeFree(ballShape)
	ShapeFree(ground)
	t.Log(GetVersionString())
}
