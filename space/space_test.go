package space_test

import (
	"github.com/egoholic/ra/relation"
	. "github.com/egoholic/ra/space"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("space", func() {
	Describe(".Add()", func() {
		It("successfully adds new relation", func() {
			space := New()
			r, err := space.Find("new-rel")
			Expect(r).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("relation `new-rel` not found"))
			space.Add("new-rel", &relation.Relation{})
			r, err = space.Find("new-rel")
			Expect(r).NotTo(BeNil())
			Expect(r).To(BeAssignableToTypeOf(&relation.Relation{}))
			Expect(err).NotTo(HaveOccurred())
		})

		It("fails to add a relation with existing relation variable (name in space)", func() {
			space := New()
			err := space.Add("new-rel", &relation.Relation{})
			Expect(err).NotTo(HaveOccurred())
			err = space.Add("new-rel", &relation.Relation{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("relation `new-rel` already exists"))
		})
	})

	Describe(".Find()", func() {
		It("successfully finds existing relation by relation variable", func() {
			space := New()
			space.Add("new-rel", &relation.Relation{})
			r, err := space.Find("new-rel")
			Expect(r).NotTo(BeNil())
			Expect(r).To(BeAssignableToTypeOf(&relation.Relation{}))
			Expect(err).NotTo(HaveOccurred())
		})

		It("fails to find relation that not exists in the space", func() {
			space := New()
			r, err := space.Find("new-rel")
			Expect(r).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("relation `new-rel` not found"))
		})
	})
})
