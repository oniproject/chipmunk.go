%module cp

//%include <typemaps.i>

%nodefaultctor;
%nodefaultdtor;

%rename("%(strip:[cp])s") "";

%ignore "cpSpaceAddDefaultCollisionHandler";
%ignore "cpSpaceAddCollisionHandler";
%ignore "cpSpaceAddWildcardHandler";

%ignore "cpPolyShapeNewRaw";

%ignore "freeFunc";
%ignore "cpArrayFreeEach";

%{
//#include "chipmunk/chipmunk_private.h"
#include "chipmunk/chipmunk.h"
#include "chipmunk/chipmunk_types.h"
%}

//%include "chipmunk/chipmunk_private.h"
%include "chipmunk/chipmunk_types.h"
%include "chipmunk/chipmunk.h"

%include "chipmunk/cpVect.h"
%include "chipmunk/cpBB.h"
%include "chipmunk/cpTransform.h"
/*
%include "chipmunk/cpSpatialIndex.h"
%include "chipmunk/cpArbiter.h"
*/
%include "chipmunk/cpBody.h"
%include "chipmunk/cpShape.h"
%include "chipmunk/cpPolyShape.h"
//%include "chipmunk/cpConstraint.h"

%include "chipmunk/cpSpace.h"

//%rename(Wrapper_cpBody) cpBody;
//%go_import("fmt")

%insert(go_wrapper) %{
type Body interface {
	Struct_SS_cpBody
	SetPosition(x,y float64)
}
/*func (ptr SwigcptrStruct_SS_cpBody) SetPosition(x, y float64) {
	BodySetPosition(ptr, V(x, y))
	fmt.Println("set pos", x, y)
}*/
%}
