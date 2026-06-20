// Decodes twitter_status.json (mixed objects, strings, numbers) into Go's generic
// any value (the representation encoding/json produces for interface{}:
// map[string]any, []any, float64, string, bool, nil). This exercises
// lightning's dynamic decoder, support.DecodeValue — the path taken for any
// fields — end to end, and compares it with the other libraries decoding the
// same input into interface{}. Pairs with the typed twitter_status case to show the
// cost of decoding into any versus a concrete struct.
package bench

type Benchmark any
