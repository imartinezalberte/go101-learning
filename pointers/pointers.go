// pointers
//
// All variables are addressablek, whereas constants, function calls and explicit conversion results are all unaddressable.
//
// A pointer value can't be converted to an arbitrary pointer type
// 1. The underlying types of type T1 and T2 are identical (ignoring struct tags), in particular if either T1 and T2 is a 
//		unnamed type and their unnamed types are identical (considering struct tags), then the conversion can be implicit. 
// 2. Type T1 and T2 are both unnamed pointer types and the underlying types of their base types are identical (ignoring struct tags).
package pointers

type (
	Ptr  *int
	PPtr *Ptr

	MyInt int64
	Ta *int64
	Tb *MyInt
)
