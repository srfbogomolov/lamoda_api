package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/srfbogomolov/warehouse_api/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestItem_NewItem(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		size        uint
		expectedErr error
	}
	testCases := []testCase{
		{
			test:        "Create new item with not empty name",
			name:        "not empty name",
			size:        0,
			expectedErr: nil,
		},
		{
			test:        "Create new item with empty name",
			name:        "",
			size:        0,
			expectedErr: domain.ErrEmptyItemName,
		},
	}

	for _, tc := range testCases {
		_, err := domain.NewItem(tc.name, tc.size)
		assert.Equal(t, tc.expectedErr, err, tc.test)
	}
}

func TestItem_EqualTo(t *testing.T) {
	var id uint = 1
	var anotherId uint = 0
	code := uuid.New()
	anotherCode := uuid.New()
	item, _ := domain.NewItem("test", 0)
	item.Id = id
	item.Code = code

	type testCase struct {
		test     string
		compared domain.Item
		expected bool
	}
	testCases := []testCase{
		{
			test:     "Item with the same id and code",
			compared: domain.Item{Id: id, Code: code},
			expected: true,
		},
		{
			test:     "Item with another id",
			compared: domain.Item{Id: anotherId, Code: code},
			expected: false,
		},
		{
			test:     "Item with another code",
			compared: domain.Item{Id: id, Code: anotherCode},
			expected: false,
		},
		{
			test:     "Item with another id and code",
			compared: domain.Item{Id: anotherId, Code: anotherCode},
			expected: false,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, item.EqualTo(tc.compared), tc.test)
	}
}
