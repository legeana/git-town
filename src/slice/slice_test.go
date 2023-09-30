package slice_test

import (
	"testing"

	"github.com/git-town/git-town/v9/src/domain"
	"github.com/git-town/git-town/v9/src/slice"
	"github.com/git-town/git-town/v9/src/steps"
	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	t.Parallel()

	t.Run("AppendAllMissing", func(t *testing.T) {
		t.Parallel()
		t.Run("slice type", func(t *testing.T) {
			t.Parallel()
			list := []string{"one", "two", "three"}
			give := []string{"two", "four", "five"}
			have := slice.AppendAllMissing(list, give)
			want := []string{"one", "two", "three", "four", "five"}
			assert.Equal(t, want, have)
		})
		t.Run("aliased slice type", func(t *testing.T) {
			t.Parallel()
			list := domain.SHAs{domain.NewSHA("111111"), domain.NewSHA("222222")}
			give := domain.SHAs{domain.NewSHA("333333"), domain.NewSHA("444444")}
			have := slice.AppendAllMissing(list, give)
			want := domain.SHAs{domain.NewSHA("111111"), domain.NewSHA("222222"), domain.NewSHA("333333"), domain.NewSHA("444444")}
			assert.Equal(t, want, have)
		})
	})

	t.Run("Contains", func(t *testing.T) {
		t.Parallel()
		give := []string{"one", "two"}
		assert.True(t, slice.Contains(give, "one"))
		assert.True(t, slice.Contains(give, "two"))
		assert.False(t, slice.Contains(give, "three"))
	})

	t.Run("FindAll", func(t *testing.T) {
		t.Parallel()
		t.Run("list contains the element", func(t *testing.T) {
			t.Parallel()
			list := []int{1, 2, 1, 3, 1}
			have := slice.FindAll(list, 1)
			want := []int{0, 2, 4}
			assert.Equal(t, want, have)
		})
		t.Run("list does not contain the element", func(t *testing.T) {
			t.Parallel()
			list := []int{1, 2, 3}
			have := slice.FindAll(list, 4)
			want := []int{}
			assert.Equal(t, want, have)
		})
	})

	t.Run("FirstElementOr", func(t *testing.T) {
		t.Parallel()
		t.Run("list contains an element", func(t *testing.T) {
			t.Parallel()
			list := []string{"one"}
			have := slice.FirstElementOr(list, "other")
			want := "one"
			assert.Equal(t, want, have)
		})
		t.Run("list is empty", func(t *testing.T) {
			t.Parallel()
			list := []string{}
			have := slice.FirstElementOr(list, "other")
			want := "other"
			assert.Equal(t, want, have)
		})
	})

	t.Run("Hoist", func(t *testing.T) {
		t.Parallel()
		t.Run("already hoisted", func(t *testing.T) {
			t.Parallel()
			give := []string{"initial", "one", "two"}
			have := slice.Hoist(give, "initial")
			want := []string{"initial", "one", "two"}
			assert.Equal(t, want, have)
		})
		t.Run("contains the element to hoist", func(t *testing.T) {
			t.Parallel()
			give := []string{"alpha", "initial", "omega"}
			have := slice.Hoist(give, "initial")
			want := []string{"initial", "alpha", "omega"}
			assert.Equal(t, want, have)
		})
		t.Run("empty list", func(t *testing.T) {
			t.Parallel()
			give := []string{}
			have := slice.Hoist(give, "initial")
			want := []string{}
			assert.Equal(t, want, have)
		})
		t.Run("aliased slice type", func(t *testing.T) {
			t.Parallel()
			give := domain.LocalBranchNames{domain.NewLocalBranchName("alpha"), domain.NewLocalBranchName("initial"), domain.NewLocalBranchName("omega")}
			have := slice.Hoist(give, domain.NewLocalBranchName("initial"))
			want := domain.LocalBranchNames{domain.NewLocalBranchName("initial"), domain.NewLocalBranchName("alpha"), domain.NewLocalBranchName("omega")}
			assert.Equal(t, want, have)
		})
	})

	t.Run("LowerAll", func(t *testing.T) {
		t.Parallel()
		t.Run("list contains element at the last position", func(t *testing.T) {
			t.Parallel()
			give := []string{"one", "two", "last"}
			have := slice.LowerAll(give, "last")
			want := []string{"one", "two", "last"}
			assert.Equal(t, want, have)
		})
		t.Run("list contains element in the middle", func(t *testing.T) {
			t.Parallel()
			give := []steps.Step{
				&steps.AbortMergeStep{},
				&steps.RestoreOpenChangesStep{EmptyStep: steps.EmptyStep{}},
				&steps.AbortRebaseStep{},
			}
			have := slice.LowerAll[steps.Step](give, &steps.RestoreOpenChangesStep{})
			want := []steps.Step{
				&steps.AbortMergeStep{},
				&steps.AbortRebaseStep{},
				&steps.RestoreOpenChangesStep{},
			}
			assert.Equal(t, want, have)
		})
		t.Run("list does not contain the element", func(t *testing.T) {
			t.Parallel()
			give := []string{"one", "two", "three"}
			have := slice.LowerAll(give, "last")
			want := []string{"one", "two", "three"}
			assert.Equal(t, want, have)
		})
		t.Run("complex example", func(t *testing.T) {
			t.Parallel()
			give := []int{1, 2, 1, 3, 1}
			have := slice.LowerAll(give, 1)
			want := []int{2, 3, 1}
			assert.Equal(t, want, have)
		})
	})

	t.Run("Remove", func(t *testing.T) {
		t.Parallel()
		t.Run("slice type", func(t *testing.T) {
			t.Parallel()
			give := []string{"one", "two", "three"}
			have := slice.Remove(give, "two")
			want := []string{"one", "three"}
			assert.Equal(t, have, want)
		})
		t.Run("slice alias type", func(t *testing.T) {
			t.Parallel()
			give := domain.SHAs{domain.NewSHA("111111"), domain.NewSHA("222222"), domain.NewSHA("333333")}
			have := slice.Remove(give, domain.NewSHA("222222"))
			want := domain.SHAs{domain.NewSHA("111111"), domain.NewSHA("333333")}
			assert.Equal(t, have, want)
		})
	})

	t.Run("RemoveAt", func(t *testing.T) {
		t.Parallel()
		t.Run("index is within the list", func(t *testing.T) {
			t.Parallel()
			give := []int{1, 2, 3}
			have := slice.RemoveAt(give, 1)
			want := []int{1, 3}
			assert.Equal(t, want, have)
		})
		t.Run("index is at end of list", func(t *testing.T) {
			t.Parallel()
			give := []int{1, 2, 3}
			have := slice.RemoveAt(give, 2)
			want := []int{1, 2}
			assert.Equal(t, want, have)
		})
		t.Run("index is at beginning of list", func(t *testing.T) {
			t.Parallel()
			give := []int{1, 2, 3}
			have := slice.RemoveAt(give, 0)
			want := []int{2, 3}
			assert.Equal(t, want, have)
		})
		t.Run("slice alias type", func(t *testing.T) {
			t.Parallel()
			give := domain.SHAs{domain.NewSHA("111111"), domain.NewSHA("222222"), domain.NewSHA("333333")}
			have := slice.RemoveAt(give, 0)
			want := domain.SHAs{domain.NewSHA("222222"), domain.NewSHA("333333")}
			assert.Equal(t, want, have)
		})
	})

	t.Run("TruncateLast", func(t *testing.T) {
		t.Parallel()
		t.Run("list contains no elements", func(t *testing.T) {
			t.Parallel()
			list := []int{}
			have := slice.TruncateLast(list)
			want := []int{}
			assert.Equal(t, want, have)
		})
		t.Run("list contains one element", func(t *testing.T) {
			t.Parallel()
			list := []int{1}
			have := slice.TruncateLast(list)
			want := []int{}
			assert.Equal(t, want, have)
		})
		t.Run("list contains multiple elements", func(t *testing.T) {
			t.Parallel()
			list := []int{1, 2, 3}
			have := slice.TruncateLast(list)
			want := []int{1, 2}
			assert.Equal(t, want, have)
		})
		t.Run("slice alias type", func(t *testing.T) {
			t.Parallel()
			list := domain.SHAs{domain.NewSHA("111111"), domain.NewSHA("222222"), domain.NewSHA("333333")}
			have := slice.TruncateLast(list)
			want := domain.SHAs{domain.NewSHA("111111"), domain.NewSHA("222222")}
			assert.Equal(t, want, have)
		})
	})
}
