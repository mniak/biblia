package transliterators

func (l Letter) CanDageshLene() bool {
	return l.IsBegadKephat()
}

func (l Letter) CanDageshForte() bool {
	return !l.IsGutural()
}

func (l Letter) IsDageshLene(hasVowelBefore bool) bool {
	if !l.IsBegadKephat() {
		return false
	}
	return !hasVowelBefore
}
