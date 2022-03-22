// Code generated by "stringer"; DO NOT EDIT.

package colfetcher

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[stateInvalid-0]
	_ = x[stateInitFetch-1]
	_ = x[stateResetBatch-2]
	_ = x[stateDecodeFirstKVOfRow-3]
	_ = x[stateFetchNextKVWithUnfinishedRow-4]
	_ = x[stateFinalizeRow-5]
	_ = x[stateEmitLastBatch-6]
	_ = x[stateFinished-7]
}

const _fetcherState_name = "stateInvalidstateInitFetchstateResetBatchstateDecodeFirstKVOfRowstateFetchNextKVWithUnfinishedRowstateFinalizeRowstateEmitLastBatchstateFinished"

var _fetcherState_index = [...]uint8{0, 12, 26, 41, 64, 97, 113, 131, 144}

func (i fetcherState) String() string {
	if i < 0 || i >= fetcherState(len(_fetcherState_index)-1) {
		return "fetcherState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _fetcherState_name[_fetcherState_index[i]:_fetcherState_index[i+1]]
}
