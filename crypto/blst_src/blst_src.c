// +build relic

// keygen.c is not included as it is imported by dkg_core and is not needed
// by bls12_381_utils
#include "hash_to_field.c"
#include "e1.c"
#include "map_to_g1.c"
#include "e2.c"
#include "map_to_g2.c"
#include "fp12_tower.c"
#include "pairing.c"
#include "aggregate.c"
#include "exp.c"
#include "sqrt.c"
#include "recip.c"
#include "bulk_addition.c"
#include "multi_scalar.c"
#include "consts.c"
#include "vect.c"
#include "exports.c"

