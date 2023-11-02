package codekata

import "testing"

func assert_equal(t *testing.T, expected, got int) {
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestChop(t *testing.T) {
	t.Run("Emtpy slice", func(t *testing.T) {
		assert_equal(t, -1, Chop(3, []int{}))
	})
	t.Run("One element slice without target value", func(t *testing.T) {
		assert_equal(t, -1, Chop(3, []int{1}))
	})
	t.Run("One element slice containing target value", func(t *testing.T) {
		assert_equal(t, 0, Chop(1, []int{1}))
	})
	t.Run("Three element slice containing target value on position 0", func(t *testing.T) {
		assert_equal(t, 0, Chop(1, []int{1, 3, 5}))
	})
	t.Run("Three element slice containing target value on position 1", func(t *testing.T) {
		assert_equal(t, 1, Chop(3, []int{1, 3, 5}))
	})
	t.Run("Three element slice containing target value on position 2", func(t *testing.T) {
		assert_equal(t, 2, Chop(5, []int{1, 3, 5}))
	})
	t.Run("Three element slice not containing target", func(t *testing.T) {
		assert_equal(t, -1, Chop(0, []int{1, 3, 5}))
	})
	t.Run("Three element slice not containing target", func(t *testing.T) {
		assert_equal(t, -1, Chop(2, []int{1, 3, 5}))
	})
	t.Run("Three element slice not containing target", func(t *testing.T) {
		assert_equal(t, -1, Chop(4, []int{1, 3, 5}))
	})
	t.Run("Three element slice not containing target", func(t *testing.T) {
		assert_equal(t, -1, Chop(6, []int{1, 3, 5}))
	})
	t.Run("Four element slice containing target at position 0", func(t *testing.T) {
		assert_equal(t, 0, Chop(1, []int{1, 3, 5, 7}))
	})
	t.Run("Four element slice containing target at position 1", func(t *testing.T) {
		assert_equal(t, 1, Chop(3, []int{1, 3, 5, 7}))
	})
	t.Run("Four element slice containing target at position 2", func(t *testing.T) {
		assert_equal(t, 2, Chop(5, []int{1, 3, 5, 7}))
	})
	t.Run("Four element slice containing target at position 3", func(t *testing.T) {
		assert_equal(t, 3, Chop(7, []int{1, 3, 5, 7}))
	})
	t.Run("Four element slice not containing target", func(t *testing.T) {
		assert_equal(t, -1, Chop(0, []int{1, 3, 5, 7}))
	})
	t.Run("Four element slice not containing target", func(t *testing.T) {
		assert_equal(t, -1, Chop(2, []int{1, 3, 5, 7}))
	})
	t.Run("Four element slice not containing target", func(t *testing.T) {
		assert_equal(t, -1, Chop(4, []int{1, 3, 5, 7}))
	})
	t.Run("Four element slice not containing target", func(t *testing.T) {
		assert_equal(t, -1, Chop(6, []int{1, 3, 5, 7}))
	})
	t.Run("Four element slice not containing target", func(t *testing.T) {
		assert_equal(t, -1, Chop(8, []int{1, 3, 5, 7}))
	})

}
