package db_test

import (
	"context"
	"testing"

	"github.com/gimmickless/keat-kit-service/internal/domain"
	"github.com/gimmickless/keat-kit-service/pkg/custom"
	"github.com/stretchr/testify/assert"
)

var (
	initBurgerCatg = domain.Category{
		Name:    "Burger",
		Desc:    "Im lovin this shit",
		ImgPath: "https://png.pngtree.com/png-clipart/20190520/original/pngtree-burger-png-image_3622097.jpg",
	}
	updateBurgerCatg = domain.Category{
		Name:    "Burger",
		Desc:    "Im lovin these sparkles",
		ImgPath: "https://png.pngtree.com/png-clipart/20190920/original/pngtree-cartoon-delicious-burger-illustration-png-image_4602812.jpg",
	}
	arbitraryMongoID = "61a7e51098974915bebd048a"
)

func TestHappyCRUDCategory(t *testing.T) {
	ctx := context.Background()

	// Act
	burgerID, err := catgRepo.Insert(ctx, initBurgerCatg)
	assert.Nil(t, err)
	assert.NotEmpty(t, burgerID)

	// Get
	c, err := catgRepo.Get(ctx, burgerID)
	assert.Nil(t, err)
	assert.Equal(t, initBurgerCatg.Name, c.Name)

	// Update
	err = catgRepo.Update(ctx, burgerID, updateBurgerCatg)
	assert.Nil(t, err)
	uc, err := catgRepo.Get(ctx, burgerID)
	assert.Nil(t, err)
	assert.Equal(t, burgerID, uc.ID)
	assert.Equal(t, updateBurgerCatg.Name, uc.Name)
	assert.Equal(t, updateBurgerCatg.Desc, uc.Desc)
	assert.Equal(t, updateBurgerCatg.ImgPath, uc.ImgPath)

	// Delete
	err = catgRepo.Delete(ctx, burgerID)
	assert.Nil(t, err)
	_, err = catgRepo.Get(ctx, burgerID)
	assert.IsType(t, &custom.ElemNotFoundError{}, err)
}

func TestWrongCreateAndWrongGetCategory(t *testing.T) {
	ctx := context.Background()

	// Act
	burgerID, err := catgRepo.Insert(ctx, initBurgerCatg)
	assert.Nil(t, err)
	assert.NotEmpty(t, burgerID)

	// Get
	_, err = catgRepo.Get(ctx, burgerID+"_")
	assert.NotNil(t, err)
}

func TestGetNonexistingCategory(t *testing.T) {
	ctx := context.Background()

	// Get
	_, err := catgRepo.Get(ctx, arbitraryMongoID)
	assert.IsType(t, &custom.ElemNotFoundError{}, err)
}

func TestUpdateNonexistingCategory(t *testing.T) {
	ctx := context.Background()

	// Get
	err := catgRepo.Update(ctx, arbitraryMongoID, updateBurgerCatg)
	assert.IsType(t, &custom.ElemNotFoundError{}, err)
}

func TestDeleteNonexistingCategory(t *testing.T) {
	ctx := context.Background()

	// Get
	err := catgRepo.Delete(ctx, arbitraryMongoID)
	assert.IsType(t, &custom.ElemNotFoundError{}, err)
}
